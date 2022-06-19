package main

import (
	"encoding/json"
	"fmt"
	"os"
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
	filename := "post_marshall3.json"
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

	jsonFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding data to JSON:", err)
		return
	}

}
