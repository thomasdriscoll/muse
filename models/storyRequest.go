package models

type StoryFromURLRequest struct {
	author   string
	authorId string
	urlType  string
	url      string
	title    string
}

type StoryFromFileRequest struct {
	author   string
	authorId string
	fileType string
	file     File
	title    string
}
