package cloudflare_client

// CloudflareProviderConfigs defines Provider Configuration
type CloudflareProviderConfigs struct {
	Providers []CloudflareProviderConfig `yaml:"providers"  mapstructure:"providers"`
}

// CloudflareProviderConfig cloudflare config
type CloudflareProviderConfig struct {
	Token    string   `yaml:"api_token,omitempty"  mapstructure:"api_token"`
	ApiKey   string   `yaml:"api_key,omitempty"  mapstructure:"api_key"`
	ApiEmail string   `yaml:"api_email,omitempty"  mapstructure:"api_email"`
	Accounts []string `yaml:"accounts,omitempty"  mapstructure:"accounts"`
	Zones    []string `yaml:"zones,omitempty"  mapstructure:"zones"`
}
