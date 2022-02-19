package models

import "os"

type StoryFromURLRequest struct {
	Author   string
	AuthorId string
	UrlType  string
	Url      string
	Title    string
}

type StoryFromFileRequest struct {
	Author   string
	AuthorId string
	FileType string
	File     *os.File
	Title    string
}
