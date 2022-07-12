package test

import "errors"

var (
	ErrIdParseFailed = errors.New("failed to parse test id")
	ErrInsertFailed  = errors.New("failed to insert test to database")
	ErrNotFound      = errors.New("test(s) not found")
	ErrFindFailed    = errors.New("internal error while find test(s)")
	ErrUpdatedFailed = errors.New("error while updating test")
	ErrDeleteFailed  = errors.New("error while deleting test")
)
