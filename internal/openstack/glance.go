package openstack

import (
	"context"
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
)

type GlanceClient struct {
	osGlanceClient *gophercloud.ServiceClient
}

func (c *GlanceClient) GetAllImages() []images.Image {
	ctx := context.Background()
	imagePages, err := images.List(c.osGlanceClient, images.ListOpts{}).AllPages(ctx)
	if err != nil {
		panic(err)
	}
	osImages, err := images.ExtractImages(imagePages)
	if err != nil {
		panic(err)
	}

	return osImages
}
