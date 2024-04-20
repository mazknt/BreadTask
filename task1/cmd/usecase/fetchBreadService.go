package usecase

import (
	"log"
	"task1/task1/cmd/config"
	"task1/task1/cmd/interface/api"
	"task1/task1/cmd/interface/repository"
)

type FetchBreadService interface {
	GetBreadInfoFromContentfulAndSave() error
}

type FetchBreadServiceImpl struct {
	ContentfulAPI api.ContentfulAPI
	Firestore     repository.Firestore
}

func (s FetchBreadServiceImpl) GetBreadInfoFromContentfulAndSave() error {
	log.Println("GetBreadInfoFromContentfulAndSave called")
	log.Println(config.AppConfig) // TODO
	// パン情報の取得
	spaceID := config.AppConfig.SpaceID
	entryIDList := config.AppConfig.EntryIDList
	accesstoken := config.AppConfig.Accesstoken
	allBbreadsInformation, getAllBreadsInformationError := s.ContentfulAPI.GetAllBreadsInformation(spaceID, entryIDList, accesstoken)
	if getAllBreadsInformationError != nil {
		log.Println("getAllBreadsInformationError:", getAllBreadsInformationError)
		return getAllBreadsInformationError
	}

	// パン情報をDBから取得
	for _, bread := range allBbreadsInformation {
		setBreadError := s.Firestore.SetBread(bread)
		if setBreadError != nil {
			log.Println("setBreadError: ", setBreadError)
			return setBreadError
		}
	}

	log.Println("GetBreadInfoFromContentfulAndSave successed")
	return nil
}
