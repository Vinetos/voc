package model

import (
	"openstack-tui/internal/openstack"
	"strconv"
	"strings"
)

type Image struct {
	OSClient *openstack.Client
}

func (s Image) ColumnHeaders() []string {
	return strings.Split("ID,NAME,SIZE,STATUS,VISIBILITY,PROTECTED,TAGS", ",")
}

func (s Image) RowData() [][]string {
	// Actually fetch the data from OpenStack
	osImages := s.OSClient.GetGlanceClient().GetAllImages()

	// Map the data to a printable format
	var data [][]string
	for _, Image := range osImages {
		row := []string{
			Image.ID,
			Image.Name,
			strconv.FormatInt(Image.SizeBytes, 10),
			string(Image.Status),
			string(Image.Visibility),
			strconv.FormatBool(Image.Protected),
			strings.Join(Image.Tags, ","),
		}
		data = append(data, row)
	}
	return data
}

func (s Image) ShortName() string {
	return "Images"
}
