package accounts

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildAccounts(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var acc cloudflare.Account
	if err := faker.FakeObject(&acc); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().Accounts(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(
		[]cloudflare.Account{acc},
		cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
		nil,
	)

	var accMem cloudflare.AccountMember
	if err := faker.FakeObject(&accMem); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().AccountMembers(
		gomock.Any(),
		acc.ID,
		gomock.Any(),
	).AnyTimes().Return(
		[]cloudflare.AccountMember{accMem},
		cloudflare.ResultInfo{
			Page:       1,
			TotalPages: 1,
		},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestAccounts(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareAccountsGenerator{}), buildAccounts, cloudflare_client.TestOptions{})
}
