package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func main() {
	fmt.Println("Chapter 7 server running...")
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodGet:
		err = handleGet(w, r)
	case http.MethodPost:
		err = handlePost(w, r)
	case http.MethodPut:
		err = handlePut(w, r)
	case http.MethodDelete:
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieveById(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieveById(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := retrieveById(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusAccepted)
	return
}
