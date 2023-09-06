package main

import "errors"

// ErrInvalidFileExtension represents the error state when the given file has invalid extension.
var ErrInvalidFileExtension = errors.New("invalid file extension")

// ErrMissingRequiredFlag represents the error state when a required flag is not given.
var ErrMissingRequiredFlag = errors.New("missing required flag")

// ErrUnknownFlag represents the error state when an unknown flag is given.
var ErrUnknownFlag = errors.New("unknown flag")
