package testhelper

import (
	"errors"
	"os"
	"path"
	"runtime"
)

func GetTextFile() (*os.File, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return nil, errors.New("something went wrong with the test helper")
	}

	path := path.Join(path.Dir(filename), "..", "testdata", "theSunAlsoRises.txt")

	return os.Open(path)
}
