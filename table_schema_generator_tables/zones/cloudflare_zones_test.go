package zones

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildZones(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var zonesResp cloudflare.ZonesResponse
	if err := faker.FakeObject(&zonesResp); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListZonesContext(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(
		zonesResp,
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestZones(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareZonesGenerator{}), buildZones, cloudflare_client.TestOptions{})
}
