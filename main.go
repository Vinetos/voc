package main

import (
	"context"
	"github.com/alecthomas/kong"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/v2/openstack/config"
	"github.com/gophercloud/gophercloud/v2/openstack/config/clouds"
	"github.com/rivo/tview"
	"os"
)

type Globals struct {
	OsClientConfig string `help:"Path to cloud.yaml config file." type:"path"`
	OsCloud        string `help:"OsCloud name to use." type:"string" default:"default"`
	ReadOnly       bool   `help:"Enable read-only mode." default:"false"`
}

type CLI struct {
	Globals
}

var app = tview.NewApplication()

func main() {
	cli := CLI{}
	kong.Parse(&cli, kong.Name("voc"), kong.Description("Vinetos Openstack Client"), kong.UsageOnError(), kong.ConfigureHelp(kong.HelpOptions{Compact: true}))

	if len(cli.OsClientConfig) != 0 {
		os.Setenv("OS_CLIENT_CONFIG_FILE", cli.OsClientConfig)
	}

	if len(cli.OsCloud) != 0 {
		// TODO: Add an option to support multi-"cloud"
		os.Setenv("OS_CLOUD", cli.OsCloud)
	}

	servers := getServers()
	flex := initGrid(servers)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}

func getServers() []servers.Server {
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

	// Use the ProviderClient and the endpoint options fetched from
	// `clouds.yaml` to build a service client: a compute client in this
	// case. Note that the contructor does not accept a `context.Context`:
	// no further call to the OpenStack API is needed at this stage.
	computeClient, err := openstack.NewComputeV2(providerClient, endpointOptions)
	if err != nil {
		panic(err)
	}

	// use the computeClient
	serversPages, err := servers.List(computeClient, servers.ListOpts{}).AllPages(ctx)
	if err != nil {
		panic(err)
	}
	servers, err := servers.ExtractServers(serversPages)
	if err != nil {
		panic(err)
	}

	return servers
}

func initGrid(servers []servers.Server) *tview.Flex {
	table := tview.NewTable().SetBorders(true)

	for index, server := range servers {
		table.SetCell(index+1, 1, tview.NewTableCell(server.Name).SetTextColor(tview.Styles.PrimaryTextColor))
	}

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(table.SetBorder(true).SetTitle("List"), 0, 3, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (3 rows)"), 3, 1, false), 0, 2, false)
	return flex
}
