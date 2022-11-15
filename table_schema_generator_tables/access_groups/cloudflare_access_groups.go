package access_groups

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareAccessGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareAccessGroupsGenerator{}

func (x *TableCloudflareAccessGroupsGenerator) GetTableName() string {
	return "cloudflare_access_groups"
}

func (x *TableCloudflareAccessGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareAccessGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareAccessGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareAccessGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneID := svc.ZoneId

			pagination := cloudflare.PaginationOptions{
				Page:    1,
				PerPage: cloudflare_client.MaxItemsPerPage,
			}

			for {
				resp, info, err := svc.ClientApi.ZoneLevelAccessGroups(ctx, zoneID, pagination)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- resp

				if !info.HasMorePages() {
					break
				}
				pagination.Page++
			}
			return nil
		},
	}
}

func (x *TableCloudflareAccessGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ZoneMultiplex()
}

func (x *TableCloudflareAccessGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("`Zone identifier tag.`").
			Extractor(cloudflare_client.ExtractorResolveZoneID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exclude").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Exclude")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Require")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("include").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Include")).Build(),
	}
}

func (x *TableCloudflareAccessGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
