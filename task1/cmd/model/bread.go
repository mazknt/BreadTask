package model

import "time"

type Bread struct {
	ID        string    `firestore:"id" json:"id"`
	BreadInfo BreadInfo `firestore:"bread_info" json:"bread_info"`
}

type BreadInfo struct {
	Name      string    `firestore:"name" json:"name"`
	CreatedAt time.Time `firestore:"createdAt" json:"createdAt"`
}
