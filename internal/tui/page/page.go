package page

import (
	"github.com/rivo/tview"
	"openstack-tui/internal/openstack"
)

type Description struct {
	Name    string
	Resize  bool
	Visible bool
}

type Page interface {
	Content(app *tview.Application, pages *tview.Pages, client *openstack.Client) tview.Primitive
	Description() Description
}
