package repository

import (
	blogging "blogging_app"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Article interface {
	Create(input blogging.Article) (uuid.UUID, error)
	GetAll() ([]blogging.Article, error)
}

type Repository struct {
	Article
}

func NewRepository(db *sqlx.DB) *Repository {
	if db == nil {
		logrus.Errorf("db is nil in NewRepository func")
	}
	return &Repository{
		Article: NewArticleRepository(db),
	}
}
