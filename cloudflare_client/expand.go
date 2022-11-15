package cloudflare_client

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func ExpandByPartition() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)
		cli := client.(*Client)
		for _, accountZones := range cli.accountsZones {
			clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
				Client: cli.withAccountID(accountZones.AccountId),
				Task:   task.Clone(),
			})
		}
		return clientTaskContextSlice
	}
}

func ZoneMultiplex() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		clientTaskContextSlice := make([]*schema.ClientTaskContext, 0)
		cli := client.(*Client)
		for _, accountZones := range cli.accountsZones {
			for _, zone := range accountZones.Zones {
				clientTaskContextSlice = append(clientTaskContextSlice, &schema.ClientTaskContext{
					Client: cli.withZoneID(accountZones.AccountId, zone),
					Task:   task.Clone(),
				})
			}

		}
		return clientTaskContextSlice
	}
}
