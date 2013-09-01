// Error Controller
package main

import "github.com/llewekam/restish"

type Error struct {
	restish.ControllerAbstract
}

func (_ Error) Read(resource *restish.Resource) (*restish.Resource, restish.StatusCode) {
	resource.Status = restish.StatusNotFound

	return resource, resource.Status
}
