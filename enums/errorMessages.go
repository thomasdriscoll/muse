package enums

type ErrorMessage string

const (
	ErrorInvalidURL    ErrorMessage = "The URL provided is invalid"
	ErrorStoryNotFound              = "Story not found in Muse DB"
	ErrorDBError                    = "Error accessing database"
)
