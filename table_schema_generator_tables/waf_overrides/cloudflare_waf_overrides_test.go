package waf_overrides

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildWAFOverrides(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var wafOverride cloudflare.WAFOverride
	if err := faker.FakeObject(&wafOverride); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWAFOverrides(
		gomock.Any(),
		cloudflare_client.TestZoneID,
	).AnyTimes().Return(
		[]cloudflare.WAFOverride{wafOverride},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestWafOverrides(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareWafOverridesGenerator{}), buildWAFOverrides, cloudflare_client.TestOptions{})
}
