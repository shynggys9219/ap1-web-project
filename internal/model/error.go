package model

import "errors"

var (
	ErrNoRecord = errors.New("models: no matching record found")

	ErrInvalidSnippetTitleLength   = errors.New("title length must be at least 1 symbol or less than 50")
	ErrInvalidSnippetContentLength = errors.New("content length must be at least 1 symbol or less than 1000")
	ErrInvalidSnippetExpiryDate    = errors.New("expiry date must be within a year")
)
