// try to parse form
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Server running...", os.Args[0])

	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	files := http.FileServer(http.Dir("cmd/chapter4/424/public"))
	http.Handle("/", http.StripPrefix("/", files))

	http.HandleFunc("/process", process)
	http.HandleFunc("/json", jsondecode)
	server.ListenAndServe()

}

// func process(w http.ResponseWriter, r *http.Request) {
// 	r.ParseMultipartForm(1024) // enctype="multipart/form-data" need to call this
// 	fileHeader := r.MultipartForm.File["uploaded"][0]
// 	file, err := fileHeader.Open() // if the file is text it will show the text content
// 	if err == nil {
// 		data, err := ioutil.ReadAll(file)
// 		if err == nil {
// 			fmt.Fprintln(w, string(data))
// 		}
// 	}
// 	fmt.Fprintln(w, r.Form)
// }

// faster if only upload 1 file
func process(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("uploaded") // return the first file
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

type body_struct struct {
	Test  string
	Hello string
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// json decode
func jsondecode(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t body_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	res, err := PrettyStruct(t)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

}
