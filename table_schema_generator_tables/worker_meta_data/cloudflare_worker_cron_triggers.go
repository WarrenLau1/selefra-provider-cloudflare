package worker_meta_data

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWorkerCronTriggersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWorkerCronTriggersGenerator{}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetTableName() string {
	return "cloudflare_worker_cron_triggers"
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			accountId := svc.AccountId
			script := task.ParentRawResult.(cloudflare.WorkerMetaData)

			resp, err := svc.ClientApi.ListWorkerCronTriggers(ctx, accountId, script.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp

			return nil
		},
	}
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("worker_meta_data_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cron").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Cron")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("modified_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("ModifiedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloudflare_worker_meta_data_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to cloudflare_worker_meta_data.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableCloudflareWorkerCronTriggersGenerator) GetSubTables() []*schema.Table {
	return nil
}
