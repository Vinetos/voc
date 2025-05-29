package model

import "github.com/rivo/tview"

// OpenStackResource defines an interface for representing different OpenStack resources
// that can be displayed in a TUI table. Any OpenStack resource type (e.g., servers,
// networks, volumes) should implement this interface to provide standardized methods
// for retrieving column headers and row data.
//
// Methods:
// - ColumnHeaders() []string: Returns the list of column names to display in the table header.
// - RowData() [][]string: Returns the list of data rows, where each row is a slice of string values.
type OpenStackResource interface {
	ColumnHeaders() []string
	RowData() [][]string
}

// FillTable populates a tview.Table with the data from any OpenStackResource.
// It automatically sets the column headers (in the first row) and fills the table
// rows with the resource's data.
//
// Parameters:
//   - table: The tview.Table to fill.
//   - resource: An OpenStackResource implementation (e.g., ServerList) providing
//     the column headers and row data.
func FillTable(table *tview.Table, resource OpenStackResource) {
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
