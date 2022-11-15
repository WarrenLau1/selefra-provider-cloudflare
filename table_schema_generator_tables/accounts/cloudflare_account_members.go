package accounts

import (
	"context"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareAccountMembersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareAccountMembersGenerator{}

func (x *TableCloudflareAccountMembersGenerator) GetTableName() string {
	return "cloudflare_account_members"
}

func (x *TableCloudflareAccountMembersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareAccountMembersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareAccountMembersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableCloudflareAccountMembersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			account := task.ParentRawResult.(cloudflare.Account)

			opt := cloudflare.PaginationOptions{
				Page:    1,
				PerPage: cloudflare_client.MaxItemsPerPage,
			}

			for {
				accountMembers, resp, err := svc.ClientApi.AccountMembers(ctx, account.ID, opt)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- accountMembers
				if !resp.HasMorePages() {
					break
				}
				opt.Page = resp.Page + 1
			}
			return nil
		},
	}
}

func (x *TableCloudflareAccountMembersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareAccountMembersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloudflare_accounts_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to cloudflare_accounts.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Code")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("User")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("roles").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Roles")).Build(),
	}
}

func (x *TableCloudflareAccountMembersGenerator) GetSubTables() []*schema.Table {
	return nil
}
