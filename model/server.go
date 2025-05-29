package model

import (
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"strings"
)

type ServerList struct {
	Servers []servers.Server
}

func (s ServerList) ColumnHeaders() []string {
	return strings.Split("PROJECT,ID,NAME,STATUS,IP", ",")
}

func (s ServerList) RowData() [][]string {
	var data [][]string
	for _, server := range s.Servers {
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
