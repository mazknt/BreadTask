package model

import "time"

type BreadInfo struct {
	Name      string    `firestore:"name" json:"name"`
	CreatedAt time.Time `firestore:"createdAt" json:"createdAt"`
}

type Bread struct {
	ID        string    `firestore:"id" json:"id"`
	Name      string    `firestore:"name" json:"name"`
	CreatedAt time.Time `firestore:"createdAt" json:"createdAt"`
}
