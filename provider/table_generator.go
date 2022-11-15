package provider

import (
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/access_groups"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/accounts"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/certificate_packs"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/dns_records"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/images"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/ips"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/waf_overrides"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/waf_packages"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/worker_meta_data"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/worker_routes"
	"github.com/selefra/selefra-provider-cloudflare/table_schema_generator_tables/zones"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&worker_routes.TableCloudflareWorkerRoutesGenerator{}),
		table_schema_generator.GenTableSchema(&zones.TableCloudflareZonesGenerator{}),
		table_schema_generator.GenTableSchema(&access_groups.TableCloudflareAccessGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&images.TableCloudflareImagesGenerator{}),
		table_schema_generator.GenTableSchema(&ips.TableCloudflareIpsGenerator{}),
		table_schema_generator.GenTableSchema(&waf_overrides.TableCloudflareWafOverridesGenerator{}),
		table_schema_generator.GenTableSchema(&worker_meta_data.TableCloudflareWorkerMetaDataGenerator{}),
		table_schema_generator.GenTableSchema(&accounts.TableCloudflareAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&certificate_packs.TableCloudflareCertificatePacksGenerator{}),
		table_schema_generator.GenTableSchema(&dns_records.TableCloudflareDnsRecordsGenerator{}),
		table_schema_generator.GenTableSchema(&waf_packages.TableCloudflareWafPackagesGenerator{}),
	}
}
