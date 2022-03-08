package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
)

type StoryRepository interface {
	FindById(ID uint64) (*models.Story, error)
	Save(story *models.Story) error
	DeleteById(ID uint64) error
	GetStoriesByField(field, fieldId string) ([]*models.Story, error)
}

type StoryRepo struct {
	db *pgx.Conn
}

func NewStoryRepo(db *pgx.Conn) *StoryRepo {
	return &StoryRepo{
		db: db,
	}
}

func (r *StoryRepo) FindById(ID uint64) (*models.Story, error) {
	var storyFromID models.Story
	err := r.db.QueryRow(context.Background(), "select * from story where id=$d", ID).Scan(&storyFromID)
	if err != nil {
		// Add logger statement here
		return nil, errors.New(enums.ErrorStoryNotFound)
	}
	return &storyFromID, nil
}

func (r *StoryRepo) Save(story *models.Story) error {
	return nil
}

func (r *StoryRepo) DeleteById(ID uint64) error {
	return nil
}

func (r *StoryRepo) GetStoriesByField(field, fieldId string) ([]*models.Story, error) {
	return nil, nil
}
