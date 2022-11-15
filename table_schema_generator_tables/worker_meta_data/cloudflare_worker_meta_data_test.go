package worker_meta_data

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildWorkerMetaData(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var workerScript cloudflare.WorkerMetaData
	if err := faker.FakeObject(&workerScript); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkerScripts(
		gomock.Any(),
	).AnyTimes().Return(
		cloudflare.WorkerListResponse{
			WorkerList: []cloudflare.WorkerMetaData{workerScript},
		},
		nil,
	)

	var workerCronTrigger cloudflare.WorkerCronTrigger
	if err := faker.FakeObject(&workerCronTrigger); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkerCronTriggers(
		gomock.Any(),
		cloudflare_client.TestAccountID,
		workerScript.ID,
	).AnyTimes().Return(
		[]cloudflare.WorkerCronTrigger{workerCronTrigger},
		nil,
	)

	var workerSecret cloudflare.WorkersSecret
	if err := faker.FakeObject(&workerSecret); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWorkersSecrets(
		gomock.Any(),
		workerScript.ID,
	).AnyTimes().Return(
		cloudflare.WorkersListSecretsResponse{
			Result: []cloudflare.WorkersSecret{workerSecret},
		},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestWorkerMetaData(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareWorkerMetaDataGenerator{}), buildWorkerMetaData, cloudflare_client.TestOptions{})
}
