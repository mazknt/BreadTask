package dto

import "time"

type GetBreadInfoResponse struct {
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
