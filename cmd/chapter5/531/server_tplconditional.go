// triggering conditional template
package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("cmd/chapter5/531/tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	fmt.Println("Server running...", os.Args[0:])
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
