package service

import (
	"context"
	"gowithprisma/data/request"
	"gowithprisma/data/response"
)

type PostService interface {
	Create(ctx context.Context, req request.PostCreateRequest)
	Update(ctx context.Context, req request.PostUpdateRequest)
	Delete(ctx context.Context, postId string)
	FindById(ctx context.Context, postId string) response.PostResponse
	FindAll(ctx context.Context) []response.PostResponse
}
