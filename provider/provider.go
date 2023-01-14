package provider

import (
	"context"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
	"os"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "cloudflare",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var cloudflareConfig cloudflare_client.CloudflareProviderConfigs
				err := config.Unmarshal(&cloudflareConfig.Providers)

				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(cloudflareConfig.Providers) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: no configuration")
				}

				if cloudflareConfig.Providers[0].ApiKey == "" {
					cloudflareConfig.Providers[0].ApiKey = os.Getenv("CLOUDFLARE_API_KEY")
				}

				if cloudflareConfig.Providers[0].ApiKey == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing ApiKey in configuration")
				}

				if cloudflareConfig.Providers[0].ApiEmail == "" {
					cloudflareConfig.Providers[0].ApiEmail = os.Getenv("CLOUDFLARE_API_EMAIL")
				}

				if cloudflareConfig.Providers[0].ApiEmail == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing ApiEmail in configuration")
				}

				clients, err := cloudflare_client.NewClients(cloudflareConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# api_key: "<YOUR_KEY>"
# api_email: "<YOUR_EMAIL>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var cloudflareConfig cloudflare_client.CloudflareProviderConfigs
				err := config.Unmarshal(&cloudflareConfig.Providers)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
