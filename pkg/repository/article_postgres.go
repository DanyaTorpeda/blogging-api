package repository

import (
	blogging "blogging_app"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type ArticleRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *ArticleRepository {
	if db == nil {
		logrus.Errorf("db is nil in NewArticleRepository func")
	}
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) Create(input blogging.Article) (uuid.UUID, error) {
	query := fmt.Sprintf(`INSERT INTO %s (id, title, description, tags, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`, articleTable)

	var id uuid.UUID
	if r.db == nil {
		logrus.Errorf("error occured db is nil")
		return uuid.Nil, errors.New("db is nil")
	}
	row := r.db.QueryRow(query, input.ID, input.Title, input.Description, pq.Array(input.Tags), input.CreatedAt, input.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *ArticleRepository) GetAll() ([]blogging.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s", articleTable)
	var articles []blogging.Article
	err := r.db.Select(&articles, query)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
