package service

import (
	blogging "blogging_app"
	"blogging_app/pkg/repository"

	"github.com/google/uuid"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(input blogging.Article) (uuid.UUID, error) {
	return s.repo.Create(input)
}

func (s *ArticleService) GetAll() ([]blogging.Article, error) {
	return s.repo.GetAll()
}
