// chaining handlers
package main

import (
	"fmt"
	"net/http"
)

// for chaining handler function
// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello but log first!")
// }

// func log(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
// 		fmt.Println("Handle function called - " + name)
// 		h(w, r)
// 	}
// }

// for chaining handler
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello handler serveHTTP")
}

func log2(h http.Handler) http.Handler { // func of http.Handler return http.Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called (log) - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called (protect) - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	//http.HandleFunc("/hello", log(hello))
	hello := HelloHandler{}
	http.Handle("/hello", protect(log2(hello)))
	server.ListenAndServe()
}
