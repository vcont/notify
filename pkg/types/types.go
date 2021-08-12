package types

import "github.com/projectdiscovery/goflags"

type Options struct {
	Verbose        bool                `yaml:"verbose,omitempty"`
	NoColor        bool                `yaml:"no_color,omitempty"`
	Silent         bool                `yaml:"silent,omitempty"`
	Version        bool                `yaml:"version,omitempty"`
	ProviderConfig string              `yaml:"provider_config,omitempty"`
	Providers      goflags.StringSlice `yaml:"providers,omitempty"`
	IDs            goflags.StringSlice `yaml:"ids,omitempty"`
	Proxy          string              `yaml:"proxy,omitempty"`

	Stdin     bool
	Bulk      bool   `yaml:"bulk,omitempty"`
	CharLimit int    `yaml:"char_limit,omitempty"`
	Data      string `yaml:"data,omitempty"`
}