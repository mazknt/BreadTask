package dto

import (
	"task1/task1/cmd/entity/model"
	"time"
)

type BreadCollection struct {
	BreadDocuments map[string]BreadDocument
}

type BreadDocument struct {
	BreadInfo model.BreadInfo `firestore:"breadInfo" json:"breadInfo"`
	CreatedAt time.Time       `firestore:"createdAt" json:"createdAt"`
	UpdatedAt time.Time       `firestore:"updatedAt" json:"updatedAt"`
}
