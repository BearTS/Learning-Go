package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{ID: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, error := json.Marshal(posts)
	if error != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post Post
	error := json.NewDecoder(req.Body).Decode(&post)
	if error != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the posts array"}`))
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, error := json.Marshal(posts)
	resp.Write(result)
}
