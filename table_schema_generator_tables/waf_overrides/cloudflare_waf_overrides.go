package waf_overrides

import (
	"context"

	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWafOverridesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWafOverridesGenerator{}

func (x *TableCloudflareWafOverridesGenerator) GetTableName() string {
	return "cloudflare_waf_overrides"
}

func (x *TableCloudflareWafOverridesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWafOverridesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWafOverridesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareWafOverridesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId

			resp, err := svc.ClientApi.ListWAFOverrides(ctx, zoneId)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp

			return nil
		},
	}
}

func (x *TableCloudflareWafOverridesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ZoneMultiplex()
}

func (x *TableCloudflareWafOverridesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("urls").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("URLs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rewrite_action").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RewriteAction")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paused").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Paused")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("`Zone identifier tag.`").
			Extractor(cloudflare_client.ExtractorResolveZoneID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Groups")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Rules")).Build(),
	}
}

func (x *TableCloudflareWafOverridesGenerator) GetSubTables() []*schema.Table {
	return nil
}
