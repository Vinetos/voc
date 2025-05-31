package openstack

import (
	"context"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

type ComputeClient struct {
	osComputeClient *gophercloud.ServiceClient
}

func (c *ComputeClient) GetAllServers() []servers.Server {
	ctx := context.Background()
	// use the computeClient
	serversPages, err := servers.List(c.osComputeClient, servers.ListOpts{}).AllPages(ctx)
	if err != nil {
		panic(err)
	}
	osServers, err := servers.ExtractServers(serversPages)
	if err != nil {
		panic(err)
	}

	return osServers
}
