// template actions
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("cmd/chapter5/533/tmpl.html")
	t.Execute(w, "H3llo")
}

func main() {
	fmt.Println("Server running...", os.Args[0:])
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
