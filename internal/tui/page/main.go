package page

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gopkg.in/yaml.v2"
	"openstack-tui/internal/model"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui"
	"openstack-tui/internal/tui/component"
	"os"
)

// MainPage is a special page that will display all the content
type MainPage struct {
}

var vocPages = map[string]Page{
	ServerListPage:    ServerList{},
	ImageListPageName: ImageListPage{},
}

func (s MainPage) Content(app *tview.Application, pages *tview.Pages) tview.Primitive {
	// Configure the main table
	table := tview.NewTable().SetSelectable(true, false)
	// Fix first raw as it contains headers
	table.SetFixed(1, 0)
	table.SetBorder(true).
		SetBorderPadding(0, 0, 0, 0).
		SetTitle("Hello")

	// Configure Header component
	header := component.Header{
		App:   app,
		Pages: pages,
	}.Build(table)

	// List cloud
	config, err := LoadCloudConfig(os.Getenv("OS_CLIENT_CONFIG_FILE"))
	if err != nil {
		panic(err)
	}

	tui.FillTable(table, config)

	table.SetSelectedFunc(func(row, column int) {
		text := table.GetCell(row, 0).Text

		for _, page := range vocPages {
			// This is very slow as we load all content to generate page instead of doing it lazily/on-demand
			pages.AddPage(page.Name(), page.Content(app, pages, func() *openstack.Client {
				// Generate OpenStack client from input
				os.Setenv("OS_CLOUD", text)
				osClient, err := openstack.NewClientFromCloudConfig()
				if err != nil {
					panic(err)
				}
				return osClient
			}), true, false)

			// Switch to page listing servers
			pages.SwitchToPage(ServerListPage)
		}
	})

	// Build the main page
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 1, false).
		AddItem(table, 0, 3, true)

	// Open prompt to view data
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == ':' {
			// Focus on header to trigger input prompt
			if !header.HasFocus() {
				app.SetFocus(header)
			}
		}
		return event
	})
	return flex
}

// LoadCloudConfig loads and parses a c:oimlouds.yaml file from the given path.
func LoadCloudConfig(path string) (*model.CloudConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read clouds.yaml: %w", err)
	}

	var config model.CloudConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse clouds.yaml: %w", err)
	}

	return &config, nil
}
