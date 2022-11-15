package images

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildImages(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var image cloudflare.Image
	if err := faker.FakeObject(&image); err != nil {
		t.Fatal(err)
	}
	image.Metadata = map[string]interface{}{"a": "b"}

	mock.EXPECT().ListImages(
		gomock.Any(),
		cloudflare_client.TestAccountID,
		cloudflare.PaginationOptions{},
	).AnyTimes().Return(
		[]cloudflare.Image{image},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestImages(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareImagesGenerator{}), buildImages, cloudflare_client.TestOptions{})
}
