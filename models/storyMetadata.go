package models

import "time"

type StoryMetadata struct {
	author      string
	authorId    string
	dateCreated time.Time
	dateUpdated time.Time
	description string
	likes       uint64
	storyId     string
	storyTitle  string
	tags        []string
	views       uint64
}

func (smd *StoryMetadata) UpdateTime() time.Time {
	smd.dateUpdated = time.Now()
	return smd.dateUpdated
}

/* func (smd *StoryMetadata) CreatedTimeByDayMonthYear(day, month, year int) time.Time {
	smd.dateCreated = time.Date(year, month, day)
	return smd.dateCreated
} */
