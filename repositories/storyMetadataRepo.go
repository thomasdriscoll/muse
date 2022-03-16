package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/thomasdriscoll/muse/enums"
	"github.com/thomasdriscoll/muse/models"
)

type StoryMetadataRepository interface {
	GetStoryById(ID string) (*models.StoryMetadata, error)
	Save(story *models.StoryMetadata) error
	DeleteById(ID string) error
	GetStoryByRandom() (*models.StoryMetadata, error)
	GetStoriesByAuthorId(authorId string) (*[]models.StoryMetadata, error)
	GetStoriesByTag(tag string) (*[]models.StoryMetadata, error)
}

type StoryMetadataRepo struct {
	db *pgx.Conn
}

func NewStoryMetadataRepo(db *pgx.Conn) *StoryMetadataRepo {
	return &StoryMetadataRepo{
		db: db,
	}
}

func (r *StoryMetadataRepo) GetStoryById(ID string) (*models.StoryMetadata, error) {
	var storyFromID models.StoryMetadata
	err := r.db.QueryRow(context.Background(), "select * from story where id=$d", ID).Scan(&storyFromID)
	if err != nil {
		// Add logger statement here
		return nil, errors.New(enums.ErrorStoryNotFound)
	}
	return &storyFromID, nil
}

func (r *StoryMetadataRepo) Save(story *models.StoryMetadata) error {
	return nil
}

func (r *StoryMetadataRepo) DeleteById(ID string) error {
	return nil
}
func (r *StoryMetadataRepo) GetStoryByRandom() (*models.StoryMetadata, error) {
	return nil, nil
}
func (r *StoryMetadataRepo) GetStoriesByAuthorId(authorId string) (*[]models.StoryMetadata, error) {
	return nil, nil
}

func (r *StoryMetadataRepo) GetStoriesByTag(tag string) (*[]models.StoryMetadata, error) {
	return nil, nil
}
