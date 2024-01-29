package controller

import (
	"gowithprisma/data/request"
	"gowithprisma/data/response"
	"gowithprisma/helper"
	"gowithprisma/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController struct {
	PostService service.PostService
}

func NewPostController(postService service.PostService) *PostController {
	return &PostController{PostService: postService}
}

func (c *PostController) Create(writer http.ResponseWriter, req *http.Request, parms httprouter.Params) {
	postCreateRequest := request.PostCreateRequest{}
	helper.ReadRequestBody(req, &postCreateRequest)
	c.PostService.Create(req.Context(), postCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (c *PostController) Update(writer http.ResponseWriter, req *http.Request, parms httprouter.Params) {
	postUpdateRequest := request.PostUpdateRequest{}
	helper.ReadRequestBody(req, &postUpdateRequest)

	postId := parms.ByName("postId")
	postUpdateRequest.Id = postId

	c.PostService.Update(req.Context(), postUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (c *PostController) FindAll(writer http.ResponseWriter, req *http.Request, parms httprouter.Params) {
	result := c.PostService.FindAll(req.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (c *PostController) FindById(writer http.ResponseWriter, req *http.Request, parms httprouter.Params) {
	postId := parms.ByName("postId")
	result := c.PostService.FindById(req.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (c *PostController) Delete(writer http.ResponseWriter, req *http.Request, parms httprouter.Params) {
	postId := parms.ByName("postId")
	c.PostService.Delete(req.Context(), postId)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}
	helper.WriteResponseBody(writer, webResponse)
}
