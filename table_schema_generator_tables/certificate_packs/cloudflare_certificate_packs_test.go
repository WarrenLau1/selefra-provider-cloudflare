package certificate_packs

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildCertificatePacks(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var certPack cloudflare.CertificatePack
	if err := faker.FakeObject(&certPack); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListCertificatePacks(
		gomock.Any(),
		cloudflare_client.TestZoneID,
	).AnyTimes().Return(
		[]cloudflare.CertificatePack{certPack},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestCertificatePacks(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareCertificatePacksGenerator{}), buildCertificatePacks, cloudflare_client.TestOptions{})
}
