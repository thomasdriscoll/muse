package models

type Story struct {
	Metadata StoryMetadata
	Content  string
}

// Leave empty to leave story validation methods
