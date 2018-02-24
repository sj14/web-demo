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
	FindPost(postID int64) (domain.Post, error)
}

func (interactor *PostUsecases) PublishPost(userID int64, text string, time time.Time) (id int64, err error) {
	p := domain.Post{UserID: userID, Text: text, CreatedAt: time, UpdatedAt: time}
	return interactor.repository.StorePost(p)
}

func (interactor *PostUsecases) FindPost(postID int64) (domain.Post, error) {
	return interactor.repository.FindPost(postID)
}
