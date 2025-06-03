package main

import (
	"github.com/alecthomas/kong"
	"github.com/rivo/tview"
	"openstack-tui/internal/tui/page"
	"os"
)

type Globals struct {
	OsClientConfig string `help:"Path to cloud.yaml config file." type:"path"`
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
		// gophercloud will try standard locations of clouds.yaml otherwise
		os.Setenv("OS_CLIENT_CONFIG_FILE", cli.OsClientConfig)
	}

	// Init tui and Add default page
	tpages := tview.NewPages()
	tpages.AddAndSwitchToPage("cloud", page.MainPage{}.Content(app, tpages), true)

	if err := app.EnableMouse(true).SetRoot(tpages, true).SetFocus(tpages).Run(); err != nil {
		panic(err)
	}
}
