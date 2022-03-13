package models

type StoryFromURLRequest struct {
	Tags    []string
	Title   string
	UrlType string
	Url     string
}
