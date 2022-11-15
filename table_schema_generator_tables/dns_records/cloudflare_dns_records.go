package dns_records

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareDnsRecordsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareDnsRecordsGenerator{}

func (x *TableCloudflareDnsRecordsGenerator) GetTableName() string {
	return "cloudflare_dns_records"
}

func (x *TableCloudflareDnsRecordsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareDnsRecordsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareDnsRecordsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareDnsRecordsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId

			records, err := svc.ClientApi.DNSRecords(ctx, zoneId, cloudflare.DNSRecord{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- records
			return nil
		},
	}
}

func (x *TableCloudflareDnsRecordsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ZoneMultiplex()
}

func (x *TableCloudflareDnsRecordsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("zone_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ZoneName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("TTL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meta").ColumnType(schema.ColumnTypeJSON).Description("`Extra Cloudflare-specific information about the record.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locked").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Locked")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ZoneID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxied").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Proxied")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxiable").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Proxiable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data").ColumnType(schema.ColumnTypeJSON).Description("`Metadata about the record.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ModifiedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Content")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
	}
}

func (x *TableCloudflareDnsRecordsGenerator) GetSubTables() []*schema.Table {
	return nil
}
