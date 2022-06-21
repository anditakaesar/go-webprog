package main

import "testing"

func TestDecode(t *testing.T) {
	post, err := decode("post_marshall1.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello world" {
		t.Error("Wrong content, was expecting 'Hello world' but got", post.Content)
	}
}
