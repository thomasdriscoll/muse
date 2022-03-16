package testhelper

import (
	"errors"
	"path"
	"runtime"

	"github.com/thomasdriscoll/muse/models"
)

func GetTextFilePath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("something went wrong with the test helper")
	}

	path := path.Join(path.Dir(filename), "..", "testdata", "simpleTextFile.txt")

	return path, nil
}

func GetStoryMetadata() models.StoryMetadata {
	return models.StoryMetadata{}
}
