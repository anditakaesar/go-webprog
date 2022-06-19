package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	filename := "post_marshall2.json"
	post := Post{
		Id:      1,
		Content: " This is an xml string value",
		Author: Author{
			Id:   2,
			Name: "Andita Fahmi",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Wow interesting",
				Author:  "Fahmi",
			},
			{
				Id:      4,
				Content: "Also nice info",
				Author:  "Kaesar",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	err = ioutil.WriteFile(filename, output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
