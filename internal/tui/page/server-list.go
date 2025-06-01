package page

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"openstack-tui/internal/model"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui"
	"openstack-tui/internal/tui/component"
)

type ServerList struct {
}

const ServerListPage = "server"

func (s ServerList) Description() Description {
	return Description{
		Name:    ServerListPage,
		Resize:  true,
		Visible: false,
	}
}

func (s ServerList) Content(app *tview.Application, pages *tview.Pages, client *openstack.Client) tview.Primitive {
	// Configure the main table
	table := tview.NewTable().SetSelectable(true, false)
	// Fix first raw as it contains headers
	table.SetFixed(1, 0)
	table.SetBorder(true).
		SetBorderPadding(0, 0, 0, 0)

	// Fill with the data
	serverListModel := model.Server{
		OSClient: client,
	}
	table.SetTitle(fmt.Sprintf("[blue]Servers[[red]%d[blue]]", len(serverListModel.RowData())))

	tui.FillTable(table, serverListModel)

	// Configure Header component
	header := component.Header{
		App:   app,
		Pages: pages,
	}.Build(table)

	// Build the main page
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(header, 0, 1, false).
		AddItem(table, 0, 3, true)

	// Open prompt to view data
	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			// Focus on header to trigger input prompt
			if !header.HasFocus() {
				app.SetFocus(header)
			}
		}
		return event
	})
	return flex
}
