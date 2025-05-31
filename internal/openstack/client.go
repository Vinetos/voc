package openstack

import (
	"context"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/config"
	"github.com/gophercloud/gophercloud/v2/openstack/config/clouds"
)

type Client struct {
	Provider        *gophercloud.ProviderClient
	EndpointOptions gophercloud.EndpointOpts
	computeClient   *ComputeClient
	glanceClient    *GlanceClient
}

func NewClientFromCloudConfig() (*Client, error) {
	ctx := context.Background()
	// Fetch coordinates from a `cloud.yaml` in the current directory, or
	// in the well-known config directories (different for each operating
	// system).
	authOptions, endpointOptions, tlsConfig, err := clouds.Parse()
	if err != nil {
		panic(err)
	}

	// Call Keystone to get an authentication token, and use it to
	// construct a ProviderClient. All functions hitting the OpenStack API
	// accept a `context.Context` to enable tracing and cancellation.
	providerClient, err := config.NewProviderClient(ctx, authOptions, config.WithTLSConfig(tlsConfig))
	if err != nil {
		panic(err)
	}

	return &Client{
		Provider:        providerClient,
		EndpointOptions: endpointOptions,
	}, nil
}

func (c *Client) GetComputeClient() *ComputeClient {
	if c.computeClient == nil {
		computeClient, err := openstack.NewComputeV2(c.Provider, c.EndpointOptions)
		if err != nil {
			panic(err)
		}

		c.computeClient = &ComputeClient{osComputeClient: computeClient}
	}
	return c.computeClient
}

func (c *Client) GetGlanceClient() *GlanceClient {
	if c.glanceClient == nil {
		client, err := openstack.NewComputeV2(c.Provider, c.EndpointOptions)
		if err != nil {
			panic(err)
		}

		c.glanceClient = &GlanceClient{osGlanceClient: client}
	}
	return c.glanceClient
}
