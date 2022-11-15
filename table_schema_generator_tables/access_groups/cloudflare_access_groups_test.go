package access_groups

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildAccessGroups(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var accessGroup cloudflare.AccessGroup
	if err := faker.FakeObject(&accessGroup); err != nil {
		t.Fatal(err)
	}
	accessGroup.Include = []interface{}{"a"}
	accessGroup.Exclude = []interface{}{"b"}
	accessGroup.Require = []interface{}{"c"}
	mock.EXPECT().ZoneLevelAccessGroups(
		gomock.Any(),
		cloudflare_client.TestZoneID,
		cloudflare.PaginationOptions{
			Page:    1,
			PerPage: 200,
		},
	).AnyTimes().Return(
		[]cloudflare.AccessGroup{accessGroup},
		cloudflare.ResultInfo{
			Page:    1,
			PerPage: 1,
			Count:   1,
			Total:   1,
		},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestAccessGroups(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareAccessGroupsGenerator{}), buildAccessGroups, cloudflare_client.TestOptions{})
}
