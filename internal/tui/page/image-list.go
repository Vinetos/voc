package page

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"openstack-tui/internal/model"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui"
)

type ImageListPage struct {
}

const ImageListPageName = "image"

func (s ImageListPage) Description() Description {
	return Description{
		Name:    ImageListPageName,
		Resize:  true,
		Visible: true,
	}
}

func (s ImageListPage) Content(app *tview.Application, pages *tview.Pages, client *openstack.Client) tview.Primitive {
	table := tview.NewTable().SetSelectable(true, false)
	table.SetBorder(true).
		SetBorderPadding(0, 0, 0, 0).
		SetTitle("Images")

	listModel := model.Image{
		OSClient: client,
	}

	// Fill with the data
	tui.FillTable(table, listModel)

	topFlex := tview.NewFlex().AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false)
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topFlex, 0, 1, false).
		AddItem(table, 0, 3, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom (3 rows)"), 3, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == ':' {
			pages.SwitchToPage(SelectionListPage)
		}
		return event
	})

	return flex
}
