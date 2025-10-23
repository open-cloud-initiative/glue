package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
)

// Flags contains the command line flags.
type Flags struct {
	Addr        string `envconfig:"TAGS_ADDR" default:":4040"`
	DatabaseURI string `envconfig:"TAGS_DATABASE_URI" default:""`
	Environment string `envconfig:"TAGS_ENV" default:"production"`
}

// NewFlags returns a new flags.
func NewFlags() *Flags {
	return &Flags{}
}

// Config contains the configuration.
type Config struct {
	// Flags ...
	Flags *Flags
	// Stdin ...
	Stdin *os.File
	// Stdout ...
	Stdout *os.File
	// Stderr ...
	Stderr *os.File
}

// New returns a new config.
func New() *Config {
	return &Config{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Flags:  NewFlags(),
	}
}

// InitDefaultConfig initializes the default configuration.
func (c *Config) InitDefaultConfig() error {
	err := envconfig.Process("", c.Flags)
	if err != nil {
		return err
	}

	return nil
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
