package waf_packages

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWafRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWafRulesGenerator{}

func (x *TableCloudflareWafRulesGenerator) GetTableName() string {
	return "cloudflare_waf_rules"
}

func (x *TableCloudflareWafRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWafRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWafRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableCloudflareWafRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId
			pack := task.ParentRawResult.(cloudflare.WAFPackage)

			resp, err := svc.ClientApi.ListWAFRules(ctx, zoneId, pack.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp

			return nil
		},
	}
}

func (x *TableCloudflareWafRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareWafRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloudflare_waf_packages_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to cloudflare_waf_packages.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_modes").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AllowedModes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("waf_package_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("package_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PackageID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Group")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Mode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultMode")).Build(),
	}
}

func (x *TableCloudflareWafRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
