package enums

const (
	ErrorInvalidStoryRequest = "400: The request provided is invalid"
	ErrorInvalidStoryId      = "400: Provided story id is invalid"
	ErrorStoryAlreadyExists  = "400: Provided story id already exists"
	ErrorInvalidTag          = "400: Provided tag is invalid"
	ErrorAuthorNotFound      = "404: Author not found in Muse DB"
	ErrorStoryNotFound       = "404: Story not found in Muse DB"
	ErrorURLNotFound         = "404: Provided URL yielded no results"
	ErrorTagNotFound         = "404: Tag not found in Muse DB"
	ErrorDBError             = "503: Error accessing database"
)
