package repositories

import (
	"database/sql"

	"github.com/thomasdriscoll/muse/models"
)

type StoryRepository interface {
	FindById(ID int) (*models.Story, error)
	Save(story *models.Story) error
	DeleteById(ID int) error
	GetStoriesByField(field, fieldId string) ([]*models.Story, error)
}

type StoryRepo struct {
	db *sql.DB
}

func NewStoryRepo(db *sql.DB) *StoryRepo {
	return &StoryRepo{
		db: db,
	}
}

func (r *StoryRepo) FindById(ID int) (*models.Story, error) {
	return &models.Story{}, nil
}

func (r *StoryRepo) Save(story *models.Story) error {
	return nil
}

func (r *StoryRepo) DeleteById(ID int) error {

}

func (r *StoryRepo) GetStoriesByField(field, fieldId string) ([]*models.Story, error) {

}
