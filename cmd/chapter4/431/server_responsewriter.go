// example of using responseWriter
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go web programming chapter 431</title></head>
	<body><h1>Hello, go web programming chapter 431</h1></body>
	</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501) // cannot modify any response after it's called
	fmt.Fprintln(w, "No such service!")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302) // cannot modify any response after it's called
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Andita",
		Threads: []string{"one", "two", "three"},
	}
	json, _ := json.Marshal(post) //ignore error
	w.Write(json)
}

func main() {
	fmt.Println("Server running", os.Args[0:])
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
