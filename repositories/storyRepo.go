package repositories

import "database/sql"

type StoryRepo struct {
	db *sql.DB
}

func NewStoryRepo(db *sql.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}
