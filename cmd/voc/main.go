package main

import (
	"github.com/alecthomas/kong"
	"github.com/rivo/tview"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui/page"
	"os"
)

type Globals struct {
	OsClientConfig string `help:"Path to cloud.yaml config file." type:"path"`
	OsCloud        string `help:"Clouds config to use to use. Use 'all' for all cloud." type:"string" default:"all"`
	ReadOnly       bool   `help:"Enable read-only mode." default:"false"`
}

type CLI struct {
	Globals
}

var app = tview.NewApplication()
var pages = []page.Page{
	page.SelectionPage{},
	page.ResourceListPage{},
}

func main() {
	cli := CLI{}
	kong.Parse(&cli, kong.Name("voc"), kong.Description("Vinetos Openstack Client"), kong.UsageOnError(), kong.ConfigureHelp(kong.HelpOptions{Compact: true}))

	if len(cli.OsClientConfig) != 0 {
		// gophercloud will try standard locations of clouds.yaml otherwise
		os.Setenv("OS_CLIENT_CONFIG_FILE", cli.OsClientConfig)
	}

	if len(cli.OsCloud) == 0 {
		// TODO: Add an option to support multi-"cloud"
		panic("Please specify a cloud to use.")
	}
	// Specify which cloud to use
	os.Setenv("OS_CLOUD", cli.OsCloud)

	// Init OpenStack client
	osClient, err := openstack.NewClientFromCloudConfig()
	if err != nil {
		panic(err)
	}

	// Init tui
	tpages := tview.NewPages()
	for _, page := range pages {
		tpages.AddPage(page.Description().Name, page.Content(app, tpages, osClient), page.Description().Resize, page.Description().Visible)
	}

	if err := app.EnableMouse(true).SetRoot(tpages, true).SetFocus(tpages).Run(); err != nil {
		panic(err)
	}
}
