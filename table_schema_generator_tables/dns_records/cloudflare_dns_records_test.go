package dns_records

import (
	"testing"

	"github.com/cloudflare/cloudflare-go"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client/mocks"
	"github.com/selefra/selefra-provider-cloudflare/faker"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
)

func buildDNSRecords(t *testing.T, ctrl *gomock.Controller) cloudflare_client.Clients {
	mock := mocks.NewMockApi(ctrl)

	var record cloudflare.DNSRecord
	if err := faker.FakeObject(&record); err != nil {
		t.Fatal(err)
	}

	record.Meta = map[string]interface{}{
		"foo": "bar",
	}

	record.Data = map[string]interface{}{
		"foo": "bar",
	}

	mock.EXPECT().DNSRecords(
		gomock.Any(),
		cloudflare_client.TestZoneID,
		gomock.Any(),
	).AnyTimes().Return(
		[]cloudflare.DNSRecord{record},
		nil,
	)

	return cloudflare_client.Clients{
		cloudflare_client.TestAccountID: mock,
	}
}

func TestDNSRecords(t *testing.T) {
	cloudflare_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableCloudflareDnsRecordsGenerator{}), buildDNSRecords, cloudflare_client.TestOptions{})
}
