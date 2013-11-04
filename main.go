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
func Request(resource *restish.Resource, action string) (responseResource *restish.Resource, status restish.StatusCode) {
	responseResource = resource
	status = restish.StatusNotFound // The default response if the appropriate dispatcher cannot be found

	fmt.Printf("Dispatching %s %s\n", action, resource.Self.Href)
	dispatch, error := restish.GetDispatch(resource)
	if nil == error {
		responseResource, status = dispatch.Request(resource, action)
	} else {
		fmt.Printf("Dispatcher Not Found")
		responseResource.Status = status
	}

	return
}

// HTTP handler function. Looks after all requests.
func handler(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); nil != err {
			fmt.Println(err)
			writer.WriteHeader(restish.StatusServerError.Code)
		}
	}()

	var status restish.StatusCode

	fmt.Println("Request Received")

	resource := restish.RequestResource(request)

	resource, status = Request(resource, request.Method)
	resource.AddHeader("Access-Control-Allow-Origin", "http://localhost:3000")
	resource.AddHeader("Access-Control-Allow-Headers", "X-Requested-With")

	renderer := restish.NewRenderer()
	header := writer.Header()
	response := restish.ResourceResponse(resource, renderer, header)

	writer.WriteHeader(status.Code)
	fmt.Fprintf(writer, "%s", response)
}

//
func main() {
	restish.AddDefaultController(&Error{restish.ControllerAbstract{}})
	restish.AddController(&Index{restish.ControllerAbstract{}}, "/")

	fmt.Println("Controllers initialised")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
