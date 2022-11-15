package images

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareImagesGenerator{}

func (x *TableCloudflareImagesGenerator) GetTableName() string {
	return "cloudflare_images"
}

func (x *TableCloudflareImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareImagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			accountId := svc.AccountId

			records, err := svc.ClientApi.ListImages(ctx, accountId, cloudflare.PaginationOptions{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- records
			return nil
		},
	}
}

func (x *TableCloudflareImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ExpandByPartition()
}

func (x *TableCloudflareImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("variants").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Variants")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uploaded").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Uploaded")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filename").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Filename")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_signed_urls").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RequireSignedURLs")).Build(),
	}
}

func (x *TableCloudflareImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
