package certificate_packs

import (
	"context"

	"github.com/selefra/selefra-provider-cloudflare/cloudflare_client"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareCertificatePacksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareCertificatePacksGenerator{}

func (x *TableCloudflareCertificatePacksGenerator) GetTableName() string {
	return "cloudflare_certificate_packs"
}

func (x *TableCloudflareCertificatePacksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareCertificatePacksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareCertificatePacksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableCloudflareCertificatePacksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*cloudflare_client.Client)
			zoneId := svc.ZoneId

			packs, err := svc.ClientApi.ListCertificatePacks(ctx, zoneId)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- packs
			return nil
		},
	}
}

func (x *TableCloudflareCertificatePacksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return cloudflare_client.ZoneMultiplex()
}

func (x *TableCloudflareCertificatePacksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosts").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Hosts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificates").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Certificates")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary_certificate").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrimaryCertificate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("validation_errors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ValidationErrors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("validation_records").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ValidationRecords")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The Account ID of the resource.`").
			Extractor(cloudflare_client.ExtractorResolveAccountID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("`Zone identifier tag.`").
			Extractor(cloudflare_client.ExtractorResolveZoneID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
	}
}

func (x *TableCloudflareCertificatePacksGenerator) GetSubTables() []*schema.Table {
	return nil
}
