package testhelper

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"runtime"

	"github.com/thomasdriscoll/muse/models"
)

func GetStoryContent() []byte {
	path, pathErr := GetTextFilePath()
	if pathErr != nil {
		panic(pathErr.Error())
	}
	content, err := os.ReadFile(path)
	if err != nil {
		jsonContent, _ := json.Marshal(content)
		return jsonContent
	} else {
		panic(errors.New("whoopsie goof, you messed up good on the testdata"))
	}
}

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
