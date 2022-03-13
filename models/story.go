package models

type Story struct {
	StoryMetadata
	Content []byte
}

// Leave empty to leave story validation methods
