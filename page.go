package main

import "github.com/llewekam/restish"

type Page struct {
}

func (_ Page) Read(resource *restish.Resource) (*restish.Resource, restish.StatusCode) {
	resource.Properties["title"] = "Page 1"

	return resource, restish.StatusOk
}
