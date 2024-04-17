package repository

import (
	"context"
	"log"
	"os"
	"task1/task1/cmd/model"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Firestore interface {
	SetBread(model.Bread) error
}

type FirestoreImpl struct{}

func (f *FirestoreImpl) SetBread(bread model.Bread) error {
	// クライアントの生成
	dbClient, createDBClientError := createFirestoreClient()
	if createDBClientError != nil {
		return createDBClientError
	}
	defer dbClient.Close()

	// データの設定
	ctx := context.Background()

	// すでに登録している場合は更新
	_, UpdateBreadError := dbClient.Collection("breadCollection").Doc(bread.ID).Update(ctx, []firestore.Update{
		{Path: "createdAt", Value: bread.BreadInfo.CreatedAt},
		{Path: "name", Value: bread.BreadInfo.Name},
	})
	if UpdateBreadError == nil {
		return nil
	}

	// 初めて登録する場合は作成
	_, setBreadError := dbClient.Collection("breadCollection").Doc(bread.ID).Set(ctx, bread.BreadInfo)
	if setBreadError != nil {
		return setBreadError
	}
	return nil
}

func createFirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, createClientError := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"), option.WithCredentialsFile(os.Getenv("CREDENTIAL_OPTION")))
	if createClientError != nil {
		log.Println("failed to create firebase client", createClientError)
		return nil, createClientError
	}
	return client, nil
}
