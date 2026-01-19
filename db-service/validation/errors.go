package validation

import "errors"

var ErrTaskAlreadyExists = errors.New("task already exists")

var ErrTaskDoesNotExist = errors.New("task does not exist")

var ErrTitleRequired = errors.New("title is required")

var ErrTitleTooLong = errors.New("title is too long")
