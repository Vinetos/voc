package tui

import (
	"github.com/rivo/tview"
	"openstack-tui/internal/model"
)

// FillTable populates a tview.Table with the data from any OpenStackResource.
// It automatically sets the column headers (in the first row) and fills the table
// rows with the resource's data.
//
// Parameters:
//   - table: The tview.Table to fill.
//   - resource: An OpenStackResource implementation (e.g., Server) providing
//     the column headers and row data.
func FillTable(table *tview.Table, resource model.OpenStackResource) {
	// Fill the table with the headers
	headers := resource.ColumnHeaders()
	for col, header := range headers {
		table.SetCell(0, col, tview.NewTableCell(header).SetTextColor(tview.Styles.PrimaryTextColor).SetBackgroundColor(tview.Styles.ContrastBackgroundColor).SetSelectable(false))
	}

	// Populate the table with the data
	rows := resource.RowData()
	for rowIndex, row := range rows {
		for colIndex, value := range row {
			table.SetCell(rowIndex+1, colIndex, tview.NewTableCell(value).SetTextColor(tview.Styles.PrimaryTextColor))
		}
	}
}
