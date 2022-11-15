package cloudflare_client

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
)

type TestOptions struct{}

const (
	TestAccountID = "test_account"
	TestZoneID    = "test_zone"
)

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) Clients, _ TestOptions) {
	ctrl := gomock.NewController(t)
	testProvider := newTestProvider(t, ctrl, table, builder)
	config := "test : test"
	test_helper.RunProviderPullTables(testProvider, config, "./", "*")
}

func newTestProvider(t *testing.T, ctrl *gomock.Controller, table *schema.Table, builder func(*testing.T, *gomock.Controller) Clients) *provider.Provider {
	return &provider.Provider{
		Name:      "cloudflare",
		Version:   "v0.0.1",
		TableList: []*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				clients := builder(t, ctrl)

				c := &Client{
					accountsZones: AccountZones{
						TestAccountID: {
							AccountId: TestAccountID,
							Zones:     []string{TestZoneID},
						},
					},
					clients:   clients,
					ClientApi: clients[TestAccountID],
				}
				return []any{c.withZoneID(TestAccountID, TestZoneID)}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `##  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
#accounts:
#  - access_token: "xxx" # access_token
#    orgs: #  org list 
#      - "xxx"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
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
