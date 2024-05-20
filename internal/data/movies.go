package data

import "time"

//its important that each field is a capital letter so as to make it exported and visible externally
type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Version   int32     `json:"version"`
	// Runtime   int32     `json:"runtime,omitempty"`
}
