package repository

import (
	blogging "blogging_app"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type ArticleRepository struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) *ArticleRepository {
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

func (r *ArticleRepository) GetByID(id uuid.UUID) (blogging.Article, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", articleTable)
	var article blogging.Article
	err := r.db.Get(&article, query, id)
	if err != nil {
		return blogging.Article{}, err
	}

	return article, nil
}

func (r *ArticleRepository) Update(id uuid.UUID, input blogging.ArticleToUpdate) error {
	//title = $1, description = $2, tags = $3, updated_at = $4
	argId := 1
	var args []string
	var argsValues []interface{}
	if input.Title != "" {
		args = append(args, fmt.Sprintf("title = $%d", argId))
		argsValues = append(argsValues, input.Title)
		argId++
	}

	if input.Description != "" {
		args = append(args, fmt.Sprintf("description = $%d", argId))
		argsValues = append(argsValues, input.Description)
		argId++
	}

	if input.Tags != nil {
		args = append(args, fmt.Sprintf("tags = $%d", argId))
		argsValues = append(argsValues, input.Tags)
		argId++
	}

	updated_at := "NOW()"
	args = append(args, fmt.Sprintf("updated_at = $%d", argId))
	argsValues = append(argsValues, updated_at)
	argId++

	argStr := strings.Join(args, ",")
	argsValues = append(argsValues, id)
	logrus.Print(fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", articleTable, argStr, argId))
	logrus.Print(argsValues...)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", articleTable, argStr, argId)
	_, err := r.db.Exec(query, argsValues...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepository) Delete(id uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", articleTable)
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows deleted, id not found")
	}

	return nil
}
