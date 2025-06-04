// A fundamental concept in net/http servers is handlers. A handler is an object implementing the http.Handlerinterface.
// A common way to write a handler is by using the http.HandleFunc adapter on functions with the appropriate signature.
// Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments. The response writer is used to fill inthe HTTP response.
// This handler does something a little more sophitacated by reading all the HTTP request headers and echoing them into the response body.
// We register our handlers on server routes using the http.HandleFunc convenience function. It sets up the default routerin the net/http package and takes a function as an argument.
// Finally, we call the ListenAndServe with the port and a handler.nil tells it to use the default router we've just set up.

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	





	fmt.Fprintf(w, "hello\n")

}

func headers(w http.ResponseWriter, req *http.Request) {
	
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)



	http.ListenAndServe(":8090", nil)
}
