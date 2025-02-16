package service

import (
	blogging "blogging_app"
	"blogging_app/pkg/repository"
	"errors"

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

func (s *ArticleService) GetByID(id uuid.UUID) (blogging.Article, error) {
	return s.repo.GetByID(id)
}

func (s *ArticleService) Update(id uuid.UUID, input blogging.ArticleToUpdate) error {
	if !input.Validate() {
		return errors.New("invalid data")
	}
	return s.repo.Update(id, input)
}

func (s *ArticleService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
