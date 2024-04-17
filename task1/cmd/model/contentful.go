package model

import "time"

type ContentfulResponse struct {
	Sys    Sys    `json:"sys"`
	Fields Fields `json:"fields"`
}

type Sys struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

type Fields struct {
	Name string `json:"name"`
}
