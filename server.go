package main

import (
	"Learning-Go/controller"
	router "Learning-Go/http"
	"Learning-Go/repository"
	"Learning-Go/service"
	"fmt"
	"net/http"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running....")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
