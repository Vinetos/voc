package page

import (
	"github.com/rivo/tview"
	"openstack-tui/internal/openstack"
)

type Page interface {
	Name() string
	Content(app *tview.Application, pages *tview.Pages, osc func() *openstack.Client) tview.Primitive
}
