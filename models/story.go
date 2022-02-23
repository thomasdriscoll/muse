package models

type Story struct {
	Metadata StoryMetadata
	Content  []byte
}

// Leave empty to leave story validation methods
