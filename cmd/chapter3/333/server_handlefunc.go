// web server using handlefunc
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello handlefunc!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World handlefunc!")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
