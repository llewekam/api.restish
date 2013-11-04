// Controller for the Restish Index resource
package main

import (
	"github.com/llewekam/restish"
)

// Index controller. Implements the Controller interface
type Index struct {
	restish.ControllerAbstract
}

// Index Resource GET handler
func (index *Index) Read(resource *restish.Resource) (*restish.Resource, restish.StatusCode) {
	resource.Type = "urn:com.restish.page"
	resource.Properties = map[string]string{
		"title": "Hello",
	}

	return resource, restish.StatusOk
}

// Resource Options for the requested action
func (_ *Index) Options(resource *restish.Resource, action string) (*restish.Resource, restish.StatusCode) {
	switch {
	case restish.ActionRead == action:
		return resource, restish.StatusOk

	}

	return resource, restish.StatusUnauthorized
}
