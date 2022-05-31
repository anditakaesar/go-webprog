// try to parse form
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server running...")

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	files := http.FileServer(http.Dir("cmd/chapter4/421/public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	http.HandleFunc("/process", process)
	server.ListenAndServe()

}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}
