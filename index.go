// Controller for the Restish Index resource
package main

import (
	"fmt"
	"github.com/llewekam/restish"
)

// Index controller. Implements the Controller interface
type Index struct {
	restish.ControllerAbstract
}

// Index Resource GET handler
func (index *Index) Read(resource *restish.Resource) (*restish.Resource, restish.StatusCode) {
	fmt.Println("In Index Read")
	resource.Properties = map[string]string{
		"title": "Hello",
	}

	return resource, restish.StatusOk
}
