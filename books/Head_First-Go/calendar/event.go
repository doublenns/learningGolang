package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date  // Embed a Date using an anonymous field
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	// Return error is event's title has more than 30 chars
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title; longer than 30 characters")
	}
	e.title = title
	return nil
}
