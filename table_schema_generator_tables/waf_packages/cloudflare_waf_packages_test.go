package waf_packages

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildWAFPackages(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var wafPackage cloudflare.WAFPackage
	if err := faker.FakeObject(&wafPackage); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWAFPackages(
		gomock.Any(),
		cloudflare_client.TestZoneID,
	).AnyTimes().Return(
		[]cloudflare.WAFPackage{wafPackage},
		nil,
	)

	var wafGroup cloudflare.WAFGroup
	if err := faker.FakeObject(&wafGroup); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWAFGroups(
		gomock.Any(),
		cloudflare_client.TestZoneID,
		wafPackage.ID,
	).AnyTimes().Return(
		[]cloudflare.WAFGroup{wafGroup},
		nil,
	)

	var wafRule cloudflare.WAFRule
	if err := faker.FakeObject(&wafRule); err != nil {
		t.Fatal(err)
	}

	wafRule.Group = struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		"fake-group-id",
		"fake-group-name",
	}

	mock.EXPECT().ListWAFRules(
		gomock.Any(),
		cloudflare_client.TestZoneID,
		wafPackage.ID,
	).AnyTimes().Return(
		[]cloudflare.WAFRule{wafRule},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestWAFPackages(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareWafPackagesGenerator{}), buildWAFPackages, cloudflare_client.TestOptions{})
}
