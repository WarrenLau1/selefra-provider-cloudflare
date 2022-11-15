package cloudflare_client

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

const MaxItemsPerPage = 200

type AccountZones map[string]struct {
	AccountId string
	Zones     []string
}

type Clients map[string]Api

type Client struct {
	accountsZones AccountZones
	clients       Clients

	ClientApi Api
	//Cloudflare CloudflareServices
	AccountId string
	ZoneId    string
}

func (c *Client) ID() string {
	return c.AccountId
}

func (c *Client) withAccountID(accountId string) *Client {
	return &Client{
		accountsZones: c.accountsZones,
		clients:       c.clients,
		ClientApi:     c.clients[accountId],
		AccountId:     accountId,
	}
}

func (c *Client) withZoneID(accountId, zoneId string) *Client {
	return &Client{
		accountsZones: c.accountsZones,
		clients:       c.clients,
		ClientApi:     c.clients[accountId],
		AccountId:     accountId,
		ZoneId:        zoneId,
	}
}

func NewClients(configs CloudflareProviderConfigs) ([]*Client, error) {
	var clients []*Client
	for _, provider := range configs.Providers {
		client, err := newClient(provider)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func getCloudflareClient(config CloudflareProviderConfig) (*cloudflare.API, error) {
	//if config.Token == "" {
	//	return nil, errors.New("The configuration Cloudflare token name is missing")
	//}

	if config.ApiKey == "" {
		return nil, errors.New("The configuration Cloudflare ApiKey name is missing")
	}

	if config.ApiEmail == "" {
		return nil, errors.New("The configuration Cloudflare ApiEmail name is missing")
	}

	c, err := cloudflare.New(config.ApiKey, config.ApiEmail)

	if err != nil {
		return nil, errors.New("Cloudflare failed to create a connection")
	}

	return c, nil
}

func newClient(config CloudflareProviderConfig) (*Client, error) {
	apiClient, err := getCloudflareClient(config)
	if err != nil {
		return nil, err
	}

	var accountsZones = make(AccountZones)

	// Get available accounts
	accounts, _, err := apiClient.Accounts(context.Background(), cloudflare.AccountsListParams{})
	if err != nil {
		return nil, err
	}

	for _, account := range accounts {
		// Get available zones  for each account
		zones, err := apiClient.ListZonesContext(context.Background(), cloudflare.WithZoneFilters("", account.ID, ""))
		if err != nil {
			continue
		}
		var zoneIds []string
		for _, zone := range zones.Result {
			zoneIds = append(zoneIds, zone.ID)
		}

		accountsZones[account.ID] = struct {
			AccountId string
			Zones     []string
		}{
			AccountId: account.ID,
			Zones:     zoneIds,
		}
	}

	if len(accountsZones) == 0 {
		return nil, errors.New("no accounts found")
	}

	clients := make(Clients)
	for _, account := range accountsZones {
		c, err := getCloudflareClient(config)
		if err != nil {
			return nil, err
		}
		c.AccountID = account.AccountId
		clients[account.AccountId] = c
	}

	return &Client{
		ClientApi:     apiClient,
		accountsZones: accountsZones,
		clients:       clients,
	}, nil
}
