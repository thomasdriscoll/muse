package enums

type ErrorMessage string

const (
	ErrorInvalidURL     ErrorMessage = "400: The URL provided is invalid"
	ErrorAuthorNotFound              = "404: Author not found in Muse DB"
	ErrorStoryNotFound               = "404: Story not found in Muse DB"
	ErrorTagNotFound                 = "404: Tag not found in Muse DB"
	ErrorDBError                     = "503: Error accessing database"
)
