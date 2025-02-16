package service

import (
	blogging "blogging_app"
	"blogging_app/pkg/repository"

	"github.com/google/uuid"
)

type Article interface {
	Create(input blogging.Article) (uuid.UUID, error)
	GetAll() ([]blogging.Article, error)
	GetByID(id uuid.UUID) (blogging.Article, error)
	Update(id uuid.UUID, input blogging.ArticleToUpdate) error
	Delete(id uuid.UUID) error
}

type Service struct {
	Article
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Article: NewArticleService(repo.Article),
	}
}
