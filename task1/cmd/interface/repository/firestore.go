package repository

import (
	"context"
	"log"
	"task1/task1/cmd/config"
	"task1/task1/cmd/entity/dto"
	"task1/task1/cmd/entity/model"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Firestore interface {
	// 追加
	SetBread(model.Bread) error
}

type FirestoreImpl struct{}

func (f FirestoreImpl) SetBread(bread model.Bread) error {
	log.Println("SetBread Start")
	ctx := context.Background()
	breadDocument := dto.BreadDocument{
		BreadInfo: model.BreadInfo{
			Name:      bread.Name,
			CreatedAt: bread.CreatedAt,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// クライアントの生成
	dbClient, createDBClientError := createFirestoreClient()
	if createDBClientError != nil {
		return createDBClientError
	}
	defer dbClient.Close()

	// データの設定
	_, setBreadError := dbClient.Collection("breadCollection").Doc(bread.ID).Set(ctx, breadDocument)
	if setBreadError != nil {
		log.Println("setBreadError: ", setBreadError)
		return setBreadError
	}
	log.Println("SetBread Success")
	return nil
}

func createFirestoreClient() (*firestore.Client, error) {
	log.Println("config.AppConfig")
	log.Println(config.AppConfig)
	ctx := context.Background()
	client, createClientError := firestore.NewClient(ctx, config.AppConfig.ProjectID, option.WithCredentialsFile(config.AppConfig.CredentialOption))
	if createClientError != nil {
		log.Println("failed to create firebase client", createClientError)
		return nil, createClientError
	}
	return client, nil
}
