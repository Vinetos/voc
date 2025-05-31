package page

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"openstack-tui/internal/model"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui"
)

type ResourceListPage struct {
}

const ResourceListPageName = "resource-list"

func (s ResourceListPage) Description() Description {
	return Description{
		Name:    ResourceListPageName,
		Resize:  true,
		Visible: true,
	}
}

func (s ResourceListPage) Content(app *tview.Application, pages *tview.Pages, client *openstack.Client) tview.Primitive {
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).
		SetBorderPadding(0, 0, 0, 0).
		SetTitle("Servers")

	serverListModel := model.Server{
		OSClient: client,
	}

	// Fill with the data
	tui.FillTable(table, serverListModel)

	topFlex := tview.NewFlex().AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topFlex, 0, 1, false).
		AddItem(table, 0, 3, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (3 rows)"), 3, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			pages.SwitchToPage(ServerListPage)
		}
		return event
	})

	return flex
}
