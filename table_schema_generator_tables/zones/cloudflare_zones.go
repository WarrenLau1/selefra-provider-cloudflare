package zones

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareZonesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareZonesGenerator{}

func (x *TableCloudflareZonesGenerator) GetTableName() string {
	return "cloudflare_zones"
}

func (x *TableCloudflareZonesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareZonesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareZonesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareZonesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)

			opts := cloudflare.WithZoneFilters("", svc.AccountId, "")

			resp, err := svc.ClientApi.ListZonesContext(ctx, opts)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp.Result

			return nil
		},
	}
}

func (x *TableCloudflareZonesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ExpandByPartition()
}

func (x *TableCloudflareZonesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("verification_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VerificationKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_name_servers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("OriginalNS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Plan")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan_pending").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PlanPending")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("betas").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Betas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ModifiedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Host")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Account")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deactivation_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeactReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("development_mode").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("DevMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_dnshost").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OriginalDNSHost")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Permissions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vanity_name_servers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("VanityNS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paused").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Paused")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("meta").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Meta")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("original_registrar").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OriginalRegistrar")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name_servers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("NameServers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Owner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
	}
}

func (x *TableCloudflareZonesGenerator) GetSubTables() []*schema.Table {
	return nil
}
