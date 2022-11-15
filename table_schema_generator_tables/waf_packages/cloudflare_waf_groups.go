package waf_packages

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWafGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWafGroupsGenerator{}

func (x *TableCloudflareWafGroupsGenerator) GetTableName() string {
	return "cloudflare_waf_groups"
}

func (x *TableCloudflareWafGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWafGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWafGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableCloudflareWafGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId
			pack := task.ParentRawResult.(cloudflare.WAFPackage)

			resp, err := svc.ClientApi.ListWAFGroups(ctx, zoneId, pack.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp

			return nil
		},
	}
}

func (x *TableCloudflareWafGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareWafGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cloudflare_waf_packages_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to cloudflare_waf_packages.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules_count").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("RulesCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_rules_count").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("ModifiedRulesCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Mode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_modes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AllowedModes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("waf_package_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
	}
}

func (x *TableCloudflareWafGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
