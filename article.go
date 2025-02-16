package blogging

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Article struct {
	ID          uuid.UUID      `json:"id" db:"id"`
	Title       string         `json:"title" db:"title"`
	Description string         `json:"description" db:"description"`
	Tags        pq.StringArray `json:"tags" db:"tags"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

type ArticleToUpdate struct {
	Title       string         `json:"title" db:"title"`
	Description string         `json:"description" db:"description"`
	Tags        pq.StringArray `json:"tags" db:"tags"`
}

func (u *ArticleToUpdate) Validate() bool {
	if u.Title == "" && u.Description == "" && u.Tags == nil {
		return false
	}

	return true
}
