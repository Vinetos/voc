package model

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
	ShortName() string
}
