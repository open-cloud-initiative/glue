package github

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/open-cloud-initiative/glue/auth/internal/adapters/auth"
	"github.com/open-cloud-initiative/glue/auth/internal/models"
	"github.com/open-cloud-initiative/glue/auth/internal/ports"

	"github.com/google/go-github/v56/github"
	"github.com/katallaxie/pkg/cast"
	"github.com/katallaxie/pkg/slices"
	"github.com/katallaxie/pkg/utilx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

var (
	ErrNoVerifiedPrimaryEmail = errors.New("goth: no verified primary email found")
	ErrFailedFetchUser        = errors.New("goth: no failed to fetch user")
	ErrNotAllowedOrg          = errors.New("goth: user not in allowed org")
	ErrNoName                 = errors.New("goth: user has no display name set")
	ErrAuthFailedParse        = errors.New("goth: failed to parse auth params, missing code or state")
)

const NoopEmail = ""

var _ auth.Provider = (*githubProvider)(nil)

// DefaultScopes holds the default scopes used for GitHub.
var DefaultScopes = []string{"user:email", "read:user"}

type githubProvider struct {
	id            string
	name          string
	clientKey     string
	secret        string
	callbackURL   string
	enterpriseURL string
	allowedOrgs   []string
	providerType  auth.ProviderType
	client        *http.Client
	config        *oauth2.Config
	scopes        []string
}

// Opt is a function that configures the GitHub provider.
type Opt func(*githubProvider)

// WithScopes sets the scopes for the GitHub provider.
func WithScopes(scopes ...string) Opt {
	return func(p *githubProvider) {
		p.config.Scopes = scopes
	}
}

// WithAllowedOrgs sets the allowed organizations for the GitHub provider.
func WithAllowedOrgs(orgs ...string) Opt {
	return func(p *githubProvider) {
		p.allowedOrgs = orgs
	}
}

// WithEnterpriseURL sets the enterprise URL for the GitHub provider.
func WithEnterpriseURL(url string) Opt {
	return func(p *githubProvider) {
		p.enterpriseURL = url
	}
}

// New creates a new GitHub provider.
func New(clientKey, secret, callbackURL string, opts ...Opt) auth.Provider {
	p := &githubProvider{
		id:            "github",
		name:          "GitHub",
		clientKey:     clientKey,
		secret:        secret,
		callbackURL:   callbackURL,
		enterpriseURL: "",
		providerType:  auth.ProviderTypeOAuth2,
		client:        auth.DefaultClient,
		allowedOrgs:   []string{},
		scopes:        DefaultScopes,
	}

	for _, opt := range opts {
		opt(p)
	}

	p.config = newConfig(p, p.scopes...)

	return p
}

// Debug sets the provider's debug mode.
func (g *githubProvider) Debug(debug bool) {
	// No debug implementation for GitHub provider.
}

// ID returns the provider's ID.
func (g *githubProvider) ID() string {
	return g.id
}

// Name returns the provider's name.
func (g *githubProvider) Name() string {
	return g.name
}

// Type returns the provider's type.
func (g *githubProvider) Type() auth.ProviderType {
	return g.providerType
}

type authIntent struct {
	authURL      string
	codeVerifier string
}

// GetAuthURL returns the URL for the authentication end-point.
func (a *authIntent) GetAuthURL() (string, error) {
	if a.authURL == "" {
		return "", auth.ErrNoAuthURL
	}

	return a.authURL, nil
}

// CodeVerifier returns the code verifier for PKCE.
func (a *authIntent) CodeVerifier() string {
	return a.codeVerifier
}

// BeginAuth starts the authentication process.
func (g *githubProvider) BeginAuth(_ context.Context, _ ports.Auth, state string, _ auth.AuthParams) (auth.AuthIntent, error) {
	verifier := oauth2.GenerateVerifier()

	uri := g.config.AuthCodeURL(
		state,
		oauth2.S256ChallengeOption(verifier),
	)

	return &authIntent{
		authURL:      uri,
		codeVerifier: verifier,
	}, nil
}

// CompleteAuth completes the authentication process.
//
//nolint:gocyclo
func (g *githubProvider) CompleteAuth(ctx context.Context, adapter ports.Auth, params auth.AuthParams) (models.User, error) {
	u := struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		Name     string `json:"name"`
		Login    string `json:"login"`
		Picture  string `json:"avatar_url"`
		Location string `json:"location"`
	}{}

	verifier := oauth2.GenerateVerifier()

	code := params.Get("code")
	if code == "" {
		return models.User{}, auth.ErrUnimplemented
	}

	token, err := g.config.Exchange(ctx, code, oauth2.SetAuthURLParam("code_verifier", verifier))
	if err != nil {
		return models.User{}, err
	}

	gc := github.NewClient(g.config.Client(ctx, token))

	if utilx.NotEmpty(g.enterpriseURL) {
		gc, err = gc.WithEnterpriseURLs(g.enterpriseURL, g.enterpriseURL)
		if err != nil {
			return models.User{}, err
		}
	}

	gu, _, err := gc.Users.Get(ctx, "")
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:  gu.GetName(),
		Email: gu.GetEmail(),
		Image: gu.GetAvatarURL(),
		Accounts: []models.Account{
			{
				Type:              models.AccountTypeOAuth2,
				Provider:          g.ID(),
				ProviderAccountID: cast.Ptr(strconv.Itoa(u.ID)),
				AccessToken:       cast.Ptr(token.AccessToken),
				RefreshToken:      cast.Ptr(token.RefreshToken),
				ExpiresAt:         cast.Ptr(token.Expiry),
				SessionState:      token.Extra("state").(string),
			},
		},
	}

	if utilx.Empty(user.Email) && slices.Any(checkScope, g.config.Scopes...) {
		opt := &github.ListOptions{}

		for {
			emails, resp, err := gc.Users.ListEmails(ctx, opt)
			if err != nil {
				return models.User{}, err
			}

			user.Email, err = checkEmail(emails...)
			if err != nil {
				return models.User{}, err
			}

			if resp.NextPage == 0 {
				break
			}

			opt.Page = resp.NextPage
		}
	}

	if utilx.Empty(user.Email) {
		return user, ErrNoVerifiedPrimaryEmail
	}

	if len(g.allowedOrgs) > 0 && !slices.Any(checkOrg(ctx, gc, gu.GetLogin()), g.allowedOrgs...) {
		return models.User{}, ErrNotAllowedOrg
	}

	user, err = adapter.CreateUser(ctx, user)
	if err != nil {
		return models.User{}, err
	}

	user, err = adapter.GetUser(ctx, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func newConfig(p *githubProvider, scopes ...string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     p.clientKey,
		ClientSecret: p.secret,
		RedirectURL:  p.callbackURL,
		Endpoint:     endpoints.GitHub,
		Scopes:       append(DefaultScopes, scopes...),
	}

	if utilx.NotEmpty(p.enterpriseURL) {
		c.Endpoint = githubEnterpriseConfig(p.enterpriseURL)
	}

	return c
}

func checkScope(scope string) bool {
	return strings.TrimSpace(scope) == "user" || strings.TrimSpace(scope) == "user:email"
}

func checkOrg(ctx context.Context, c *github.Client, user string) func(string) bool {
	return func(org string) bool {
		m, _, err := c.Organizations.IsMember(ctx, org, user)
		if err != nil {
			return false
		}

		return m
	}
}

func checkEmail(emails ...*github.UserEmail) (string, error) {
	for _, e := range emails {
		if e.GetPrimary() && e.GetVerified() {
			return cast.Value(e.Email), nil
		}
	}

	return NoopEmail, ErrNoVerifiedPrimaryEmail
}

func githubEnterpriseConfig(url string) oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:       fmt.Sprintf("%s/login/oauth/authorize", strings.TrimSuffix(url, "/")),
		TokenURL:      fmt.Sprintf("%s/login/oauth/access_token", strings.TrimSuffix(url, "/")),
		DeviceAuthURL: fmt.Sprintf("%s/login/device/code", strings.TrimSuffix(url, "/")),
	}
}
