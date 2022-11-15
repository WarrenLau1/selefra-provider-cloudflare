package worker_meta_data

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareWorkersSecretsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareWorkersSecretsGenerator{}

func (x *TableCloudflareWorkersSecretsGenerator) GetTableName() string {
	return "cloudflare_workers_secrets"
}

func (x *TableCloudflareWorkersSecretsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareWorkersSecretsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareWorkersSecretsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableCloudflareWorkersSecretsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			script := task.ParentRawResult.(cloudflare.WorkerMetaData)

			resp, err := svc.ClientApi.ListWorkersSecrets(ctx, script.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp.Result

			return nil
		},
	}
}

func (x *TableCloudflareWorkersSecretsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareWorkersSecretsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("worker_meta_data_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_text").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloudflare_worker_meta_data_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to cloudflare_worker_meta_data.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableCloudflareWorkersSecretsGenerator) GetSubTables() []*schema.Table {
	return nil
}
