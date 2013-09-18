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

	fmt.Printf("Dispatching %s\n", resource.Self.Href)
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

	var status restish.StatusCode
	renderer := restish.NewRenderer()

	fmt.Println("Request Received")

	resource := restish.RequestResource(request)
	resource, status = Request(resource, request.Method)

	response := restish.ResourceResponse(resource)
	responseString := renderer.Render(response)

	header := writer.Header()
	header.Add("Content-type", renderer.MimeType())
	writer.WriteHeader(status.Code)

	fmt.Fprintf(writer, "%s", responseString)
}

func main() {
	restish.AddDefaultController(&Error{restish.ControllerAbstract{}})
	restish.AddController(&Index{restish.ControllerAbstract{}}, "/")

	fmt.Println("Controllers initialised")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
