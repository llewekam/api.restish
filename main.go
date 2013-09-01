package main

import (
	"fmt"
	"github.com/llewekam/restish"
	"net/http"
)

// Dispatch a resource to the appropriate controller and controller action based on the HTTP Action string.
// POST = create
// GET = read
// PUT = update
// DELETE = delete
func Request(resource *restish.Resource, httpAction string) (responseResource *restish.Resource, status restish.StatusCode) {
	var action string

	responseResource = resource
	status = restish.StatusNotFound // The default response if the appropriate dispatcher cannot be found

	switch {
	case "POST" == httpAction:
		action = restish.ActionCreate

	case "GET" == httpAction:
		action = restish.ActionRead

	case "PUT" == httpAction:
		action = restish.ActionUpdate

	case "DELETE" == httpAction:
		action = restish.ActionDelete
	}

	fmt.Printf("Dispatching %s\n", resource.Self)
	dispatch, error := restish.GetDispatch(resource)
	if nil == error {
		responseResource, status = dispatch.Request(resource, action)
	} else {
		fmt.Printf("Dispatcher Not Found")
		responseResource.Status = status
	}

	return
}

func handler(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
			writer.WriteHeader(500)
		}
	}()

	fmt.Println("Request Received")
	resource := restish.NewResource(request.RequestURI)
	responseResource, status := Request(resource, request.Method)

	renderer := restish.NewRenderer()
	response := renderer.Render(responseResource)

	header := writer.Header()
	header.Add("Content-type", renderer.MimeType())
	writer.WriteHeader(status.Code)

	fmt.Fprintf(writer, "%s", response)
}

func main() {
	errorDispatch := restish.NewDispatch(&Error{restish.ControllerAbstract{}})
	restish.SetDefaultDispatch(errorDispatch)

	index := restish.NewDispatch(&Index{restish.ControllerAbstract{}})
	restish.AddRoute(index, "/")

	fmt.Println("Dispatchers initialised")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
