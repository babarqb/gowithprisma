package service

import (
	"context"
	"gowithprisma/data/request"
	"gowithprisma/data/response"
	"gowithprisma/helper"
	"gowithprisma/model"
	"gowithprisma/repository"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
}

func NewPostServiceImpl(postRepo repository.PostRepository) PostService {
	return &PostServiceImpl{
		PostRepository: postRepo,
	}
}

// Create implements PostService.
func (p *PostServiceImpl) Create(ctx context.Context, req request.PostCreateRequest) {
	postData := model.Post{
		Title:       req.Title,
		Published:   req.Published,
		Description: req.Description,
	}
	p.PostRepository.Save(ctx, postData)
}

// Delete implements PostService.
func (p *PostServiceImpl) Delete(ctx context.Context, postId string) {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)
	p.PostRepository.Delete(ctx, post.Id)
}

// FindAll implements PostService.
func (p *PostServiceImpl) FindAll(ctx context.Context) []response.PostResponse {
	posts := p.PostRepository.FindAll(ctx)

	var postResp []response.PostResponse

	for _, value := range posts {
		post := response.PostResponse{
			Id:          value.Id,
			Title:       value.Title,
			Published:   value.Published,
			Description: value.Description,
		}
		postResp = append(postResp, post)
	}
	return postResp
}

// FindById implements PostService.
func (p *PostServiceImpl) FindById(ctx context.Context, postId string) response.PostResponse {
	post, err := p.PostRepository.FindById(ctx, postId)
	helper.ErrorPanic(err)
	return response.PostResponse{
		Id:          post.Id,
		Title:       post.Title,
		Published:   post.Published,
		Description: post.Description,
	}
}

// Update implements PostService.
func (p *PostServiceImpl) Update(ctx context.Context, req request.PostUpdateRequest) {
	postData := model.Post{
		Id:          req.Id,
		Title:       req.Title,
		Published:   req.Published,
		Description: req.Description,
	}
	p.PostRepository.Update(ctx, postData)
}
