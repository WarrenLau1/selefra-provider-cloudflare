package accounts

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareAccountsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareAccountsGenerator{}

func (x *TableCloudflareAccountsGenerator) GetTableName() string {
	return "cloudflare_accounts"
}

func (x *TableCloudflareAccountsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareAccountsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareAccountsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareAccountsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			opt := cloudflare.AccountsListParams{
				PaginationOptions: cloudflare.PaginationOptions{
					Page:    1,
					PerPage: cloudflare_client.MaxItemsPerPage,
				},
			}

			for {
				accounts, resp, err := svc.ClientApi.Accounts(ctx, opt)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- accounts
				if !resp.HasMorePages() {
					break
				}
				opt.Page = resp.Page + 1
			}
			return nil
		},
	}
}

func (x *TableCloudflareAccountsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareAccountsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_on").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedOn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Settings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableCloudflareAccountsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableCloudflareAccountMembersGenerator{}),
	}
}
