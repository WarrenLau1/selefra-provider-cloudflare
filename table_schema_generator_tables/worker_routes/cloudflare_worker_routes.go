package worker_routes

import (
	"context"

	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWorkerRoutesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWorkerRoutesGenerator{}

func (x *TableCloudflareWorkerRoutesGenerator) GetTableName() string {
	return "cloudflare_worker_routes"
}

func (x *TableCloudflareWorkerRoutesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWorkerRoutesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWorkerRoutesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareWorkerRoutesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId

			resp, err := svc.ClientApi.ListWorkerRoutes(ctx, zoneId)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp.Routes

			return nil
		},
	}
}

func (x *TableCloudflareWorkerRoutesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ZoneMultiplex()
}

func (x *TableCloudflareWorkerRoutesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("`Zone identifier tag.`").
			Extractor(cloudflare_client.ExtractorResolveZoneID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pattern").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Pattern")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("script").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Script")).Build(),
	}
}

func (x *TableCloudflareWorkerRoutesGenerator) GetSubTables() []*schema.Table {
	return nil
}
