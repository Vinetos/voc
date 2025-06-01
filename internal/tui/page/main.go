package page

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"openstack-tui/internal/openstack"
	"openstack-tui/internal/tui/component"
)

// MainPage is a special page that will display all the content
type MainPage struct {
}

const MainPagePage = "main"

func (s MainPage) Description() Description {
	return Description{
		Name:    MainPagePage,
		Resize:  true,
		Visible: true,
	}
}

func (s MainPage) Content(app *tview.Application, pages *tview.Pages, client *openstack.Client) tview.Primitive {
	// Configure the main table
	// TODO: May be I should select cloud here ?
	table := tview.NewTable().SetSelectable(true, false)
	// Fix first raw as it contains headers
	table.SetFixed(1, 0)
	table.SetBorder(true).
		SetBorderPadding(0, 0, 0, 0).
		SetTitle("Hello")

	// Configure Header component
	header := component.Header{
		App: app,
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
