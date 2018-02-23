package usecases

import "github.com/sj14/web-demo/domain"

func NewPostUsecases(postRepository postRepositoryInterface) PostUsecases {
	return PostUsecases{repository: postRepository}
}

type PostUsecases struct {
	repository postRepositoryInterface
}

type postRepositoryInterface interface {
	StorePost(post domain.Post) (id int64, err error)
}

func (interactor *PostUsecases) PublishPost(post domain.Post) (id int64, err error) {
	return interactor.repository.StorePost(post)
}
