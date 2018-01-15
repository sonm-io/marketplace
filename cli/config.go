package cli

import (
	"fmt"

	"github.com/jinzhu/configor"
	"github.com/sonm-io/marketplace/infra/accounts"
)

// Config application configuration object.
type Config struct {
	CfgPath string

	ListenAddr string `yaml:"address"`
	DataDir    string `yaml:"data_dir" required:"true" default:"./data"`

	EthCfg accounts.EthConfig `yaml:"ethereum" required:"true"`

	OrdersCleanUpPeriod string `yaml:"orders_cleanup_period" required:"true" default:"1m"`
	OrdersTTL           string `yaml:"orders_ttl" required:"true" default:"5m"`
}

// Option is a configuration parameter.
type Option func(f *Config)

// WithConfigPath sets the path to config file to load options from.
func WithConfigPath(path string) Option {
	return func(c *Config) {
		c.CfgPath = path
	}
}

// WithListenAddr sets listen address.
func WithListenAddr(addr string) Option {
	return func(c *Config) {
		c.ListenAddr = addr
	}
}

// WithDataDir sets the database path.
func WithDataDir(dirPath string) Option {
	return func(c *Config) {
		c.DataDir = dirPath
	}
}

// NewConfig instantiates Config.
func NewConfig(opts ...Option) *Config {
	conf := &Config{}
	conf.WithOptions(opts...)
	return conf
}

// FromFile loads options from file.
func (c *Config) FromFile(filePath string) error {
	if err := configor.Load(c, filePath); err != nil {
		return fmt.Errorf("cannot load config from file: %v", err)
	}
	return nil
}

// WithOptions sets the given options.
func (c *Config) WithOptions(opts ...Option) {
	for _, option := range opts {
		option(c)
	}
}
