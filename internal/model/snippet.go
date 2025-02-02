package model

import (
	"time"
	"unicode/utf8"
)

const (
	titleMaxLength    = 50
	contentMaxLength  = 1000
	expirationMaxTime = 1
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

func (s Snippet) Validate(errMap map[string]string) {
	if utf8.RuneCountInString(s.Title) < 1 || utf8.RuneCountInString(s.Title) > titleMaxLength {
		errMap["title"] = ErrInvalidSnippetTitleLength.Error()
	}

	if utf8.RuneCountInString(s.Content) < 1 || utf8.RuneCountInString(s.Content) > contentMaxLength {
		errMap["content"] = ErrInvalidSnippetContentLength.Error()
	}

	timeNow := time.Now().UTC()
	timeDifference := s.Expires.Sub(timeNow).Hours() / (24 * 365) // 1.01
	if !s.Expires.After(timeNow) || timeDifference > 1 {
		errMap["expiry"] = ErrInvalidSnippetExpiryDate.Error()
	}
}
