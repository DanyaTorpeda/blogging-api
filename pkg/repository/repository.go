package repository

import (
	blogging "blogging_app"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Article interface {
	Create(input blogging.Article) (uuid.UUID, error)
	GetAll() ([]blogging.Article, error)
	GetByID(id uuid.UUID) (blogging.Article, error)
	Update(id uuid.UUID, input blogging.ArticleToUpdate) error
	Delete(id uuid.UUID) error
}

type Repository struct {
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Article: NewArticleRepository(db),
	}
}
