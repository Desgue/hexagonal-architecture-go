package apperr

import "errors"

var (
	NotFound   = errors.New("No result found in query")
	DeleteErr  = errors.New("Error deleting from database")
	EntryExist = errors.New("Entry already exist in database")
)
