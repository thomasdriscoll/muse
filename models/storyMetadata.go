package models

import "time"

type StoryMetadata struct {
	Author      string
	AuthorId    string
	DateCreated time.Time
	DateUpdated time.Time
	Description string
	Likes       uint64
	Source      string
	SourceType  string
	StoryId     string
	StoryTitle  string
	Tags        []string
	Views       uint64
	WordCount   uint64
}

func (smd *StoryMetadata) UpdateTime() time.Time {
	smd.DateUpdated = time.Now()
	return smd.DateUpdated
}

/* func (smd *StoryMetadata) CreatedTimeByDayMonthYear(day, month, year int) time.Time {
	smd.dateCreated = time.Date(year, month, day)
	return smd.dateCreated
} */
