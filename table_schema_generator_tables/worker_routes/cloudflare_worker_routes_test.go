package worker_routes

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildWorkerRoutes(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var workerRoute cloudflare.WorkerRoute
	if err := faker.FakeObject(&workerRoute); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkerRoutes(
		gomock.Any(),
		cloudflare_client.TestZoneID,
	).AnyTimes().Return(
		cloudflare.WorkerRoutesResponse{
			Routes: []cloudflare.WorkerRoute{workerRoute},
		},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestWorkerRoutes(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareWorkerRoutesGenerator{}), buildWorkerRoutes, cloudflare_client.TestOptions{})
}
