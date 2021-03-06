package usecases

import (
	"time"

	"github.com/sj14/web-demo/domain"
)

func NewPostUsecases(postRepository postRepositoryInterface) PostUsecases {
	return PostUsecases{repository: postRepository}
}

type PostUsecases struct {
	repository postRepositoryInterface
}

type postRepositoryInterface interface {
	StorePost(post domain.Post) (id int64, err error)
	FindPostByID(postID int64) (domain.Post, error)
	FindPostsByUserID(userID int64, limit, offset int64) ([]*domain.Post, error)
	FindNewestPosts(limit, offset int64) ([]*domain.Post, error)
}

func (interactor *PostUsecases) PublishPost(userID int64, text string, time time.Time) (id int64, err error) {
	p := domain.Post{UserID: userID, Text: text, CreatedAt: time, UpdatedAt: time}
	return interactor.repository.StorePost(p)
}

func (interactor *PostUsecases) FindPostByID(postID int64) (domain.Post, error) {
	return interactor.repository.FindPostByID(postID)
}

func (interactor *PostUsecases) FindPostsByUserID(userID int64, limit, offset int64) ([]*domain.Post, error) {
	if limit > 25 {
		limit = 25
	}
	return interactor.repository.FindPostsByUserID(userID, limit, offset)
}

func (interactor *PostUsecases) FindNewestPosts(limit, offset int64) ([]*domain.Post, error) {
	if limit > 25 {
		limit = 25
	}
	return interactor.repository.FindNewestPosts(limit, offset)
}
