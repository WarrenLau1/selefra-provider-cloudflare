package ips

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableCloudflareIpsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableCloudflareIpsGenerator{}

func (x *TableCloudflareIpsGenerator) GetTableName() string {
	return "cloudflare_ips"
}

func (x *TableCloudflareIpsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableCloudflareIpsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableCloudflareIpsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableCloudflareIpsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			resp, err := cloudflare.IPs()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for _, ip := range resp.IPv4CIDRs {
				resultChannel <- IpWrapper{Ip: ip, Type: "ipv4"}
			}

			for _, ip := range resp.IPv6CIDRs {
				resultChannel <- IpWrapper{Ip: ip, Type: "ipv6"}
			}

			for _, ip := range resp.ChinaIPv4CIDRs {
				resultChannel <- IpWrapper{Ip: ip, Type: "ipv4_china"}
			}

			for _, ip := range resp.ChinaIPv6CIDRs {
				resultChannel <- IpWrapper{Ip: ip, Type: "ipv6_china"}
			}

			return nil
		},
	}
}

type IpWrapper struct {
	Ip   string
	Type string
}

func (x *TableCloudflareIpsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableCloudflareIpsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("ip").ColumnType(schema.ColumnTypeString).Description("Cloudflare ip cidr address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Ip type, ipv4, ipv6, ipv4_china, ipv6_china.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableCloudflareIpsGenerator) GetSubTables() []*schema.Table {
	return nil
}
