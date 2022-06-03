// template actions
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	//fmt.Println(t.Format(layout))
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("tmpl.html").Funcs(funcMap)
	t, _ = t.ParseFiles("cmd/chapter5/550/tmpl.html")
	t.Execute(w, time.Now())
}

func main() {
	fmt.Println("Server running...", os.Args[0:])
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
