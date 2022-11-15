package waf_packages

import (
	"context"

	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWafPackagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWafPackagesGenerator{}

func (x *TableCloudflareWafPackagesGenerator) GetTableName() string {
	return "cloudflare_waf_packages"
}

func (x *TableCloudflareWafPackagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWafPackagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWafPackagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareWafPackagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId

			resp, err := svc.ClientApi.ListWAFPackages(ctx, zoneId)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp

			return nil
		},
	}
}

func (x *TableCloudflareWafPackagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ZoneMultiplex()
}

func (x *TableCloudflareWafPackagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detection_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DetectionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sensitivity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Sensitivity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("action_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ZoneID")).Build(),
	}
}

func (x *TableCloudflareWafPackagesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableCloudflareWafGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&TableCloudflareWafRulesGenerator{}),
	}
}
