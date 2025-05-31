package main

import (
	"github.com/alecthomas/kong"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"openstack-tui/internal/model"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui"
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

func main() {
	cli := CLI{}
	kong.Parse(&cli, kong.Name("voc"), kong.Description("Vinetos Openstack Client"), kong.UsageOnError(), kong.ConfigureHelp(kong.HelpOptions{Compact: true}))

	if len(cli.OsClientConfig) != 0 {
		// gophercloud will try standard locations of clouds.yaml otherwise
		os.Setenv("OS_CLIENT_CONFIG_FILE", cli.OsClientConfig)
	}

	if len(cli.OsCloud) == 0 {
		// TODO: Add an option to support multi-"cloud"
		panic("Please specify a cloud to use.s")
	}
	// Specify which cloud to use
	os.Setenv("OS_CLOUD", cli.OsCloud)

	// Init OpenStack client
	osClient, err := openstack.NewClientFromCloudConfig()
	if err != nil {
		panic(err)
	}

	flex, table := initGrid(osClient)

	if err := app.SetRoot(flex, true).SetFocus(table).Run(); err != nil {
		panic(err)
	}
}

func initGrid(client *openstack.Client) (*tview.Flex, *tview.Table) {
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).
		SetBorderPadding(0, 0, 0, 0).
		SetTitle("Servers")

	serverListModel := model.Server{
		OSClient: client,
	}

	// Fill with the data
	tui.FillTable(table, serverListModel)

	topFlex := tview.NewFlex().AddItem(initPrompt(), 0, 1, false)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topFlex, 0, 1, false).
		AddItem(table, 0, 3, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (3 rows)"), 3, 1, false)

	return flex, table
}

func initPrompt() *tview.InputField {
	inputField := tview.NewInputField().
		SetLabel("> ").
		SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

	inputField.SetBorder(true)

	return inputField
}
