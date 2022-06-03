// triggering conditional template
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("cmd/chapter5/532/tmpl.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func main() {
	fmt.Println("Server running...", os.Args[0:])
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
