package providers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/open-cloud-initiative/glue/auth/internal/models"
	"github.com/open-cloud-initiative/glue/auth/internal/ports"
)

const (
	defaultTimeout             = 10 * time.Second
	defaultMaxIdleConnsPerHost = 20
)

// Unknown is an unknown string.
const Unknown = "unknown"

// DefaultClient is the default HTTP client used.
var DefaultClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: defaultMaxIdleConnsPerHost,
	},
	Timeout: defaultTimeout,
}

// ErrUnimplemented is returned when a method is not implemented.
var ErrUnimplemented = errors.New("not implemented")

// ErrNoAuthURL is returned when an AuthURL has not been set.
var ErrNoAuthURL = errors.New("an AuthURL has not been set")

// Provider needs to be implemented for each 3rd party authentication provider.
type Provider interface {
	// ID returns the provider's ID.
	ID() string
	// Debug sets the provider's debug mode.
	Debug(bool)
	// Name returns the provider's name.
	Name() string
	// Type returns the provider's type.
	Type() ProviderType
	// BeginAuth starts the authentication process.
	BeginAuth(ctx context.Context, adapter ports.Auth, state string, params AuthParams) (AuthIntent, error)
	// CompleteAuth completes the authentication process.
	CompleteAuth(ctx context.Context, adapter ports.Auth, params AuthParams) (models.User, error)
}

// AuthParams is the type of authentication parameters.
type AuthParams interface {
	//  Get returns the value of a parameter by name.
	Get(string) string
	// CodeVerifier returns the code verifier for PKCE, if applicable.
	CodeVerifier() string
}

// AuthIntent is the type of authentication intent.
type AuthIntent interface {
	// GetAuthURL returns the URL for the authentication end-point.
	GetAuthURL() (string, error)
	// CodeVerifier returns the code verifier for PKCE, if applicable.
	CodeVerifier() string
}

// PrioviderType is the type of provider.
type ProviderType string

const (
	// ProviderTypeOAuth2 represents an OAuth2 account type.
	ProviderTypeOAuth2 ProviderType = "oauth2"
	// ProviderTypeOIDC represents an OIDC account type.
	ProviderTypeOIDC ProviderType = "oidc"
	// ProviderTypeSAML represents a SAML account type.
	ProviderTypeSAML ProviderType = "saml"
	// ProviderTypeEmail represents an email account type.
	ProviderTypeEmail ProviderType = "email"
	// ProviderTypeWebAuthn represents a WebAuthn account type.
	ProviderTypeWebAuthn ProviderType = "webauthn"
	// ProviderTypeUnknown represents an unknown account type.
	ProviderTypeUnknown ProviderType = "unknow"
)

// Providers is list of known/available providers.
type Providers map[string]Provider

var providers = Providers{}

// RegisterProvider adds a provider to the list of available providers for use with Goth.
func RegisterProvider(provider ...Provider) {
	for _, p := range provider {
		providers[p.ID()] = p
	}
}

// GetProviders returns a list of all the providers currently in use.
func GetProviders() Providers {
	return providers
}

// GetProvider returns a previously created provider. If Goth has not
// been told to use the named provider it will return an error.
func GetProvider(name string) (Provider, error) {
	provider := providers[name]
	if provider == nil {
		return nil, fmt.Errorf("no provider for %s exists", name)
	}

	return provider, nil
}
