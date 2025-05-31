package model

import (
	"openstack-tui/internal/openstack"
	"strings"
)

type Server struct {
	OSClient *openstack.Client
}

func (s Server) ColumnHeaders() []string {
	return strings.Split("PROJECT,ID,NAME,STATUS,IP", ",")
}

func (s Server) RowData() [][]string {
	// Actually fetch the data from OpenStack
	osServers := s.OSClient.GetComputeClient().GetAllServers()

	// Map the data to a printable format
	var data [][]string
	for _, server := range osServers {
		row := []string{
			server.TenantID,
			server.ID,
			server.Name,
			server.Status,
			server.AccessIPv4,
		}
		data = append(data, row)
	}
	return data
}

func (s Server) ShortName() string {
	return "servers"
}
