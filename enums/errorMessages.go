package enums

const (
	ErrorInvalidStoryId      = "400: Provided story id is invalid"
	ErrorInvalidStoryRequest = "400: The request provided is invalid"
	ErrorInvalidTag          = "400: Provided tag is invalid"
	ErrorStoryAlreadyExists  = "400: Provided story id already exists"
	ErrorStoryContentInvalid = "400: Provided story did not have associated content"
	ErrorAuthorNotFound      = "404: Author not found in Muse DB"
	ErrorStoryNotFound       = "404: Story not found in Muse DB"
	ErrorTagNotFound         = "404: Tag not found in Muse DB"
	ErrorURLNotFound         = "404: Provided URL yielded no results"
	ErrorDBError             = "503: Error accessing database"
)
