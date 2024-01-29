package router

import (
	"fmt"
	"gowithprisma/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(postController *controller.PostController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Print(w, "Welcome Home")
	})
	router.GET("/api/post", postController.FindAll)
	router.GET("/api/post/:postId", postController.FindById)
	router.POST("/api/post", postController.Create)
	router.PATCH("/api/post/:postId", postController.Update)
	router.DELETE("/api/post/:postId", postController.Delete)

	return router
}
