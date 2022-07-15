package routes

import (
	"Learning-Go/entity"
	"Learning-Go/repository"
	"encoding/json"
	"math/rand"
	"net/http"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, error := repo.FindAll()
	if error != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error Getting the Posts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entity.Post
	error := json.NewDecoder(req.Body).Decode(&post)
	if error != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the posts array"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
