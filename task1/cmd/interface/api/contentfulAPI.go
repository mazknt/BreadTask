package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"task1/task1/cmd/entity/dto"
	"task1/task1/cmd/entity/model"
)

type ContentfulAPI interface {
	GetAllBreadsInformation(spaceID string, entryIDList []string, accessToken string) ([]model.Bread, error)
}

type ContentfulAPIImpl struct{}

/**
 * Contentfulからパンの情報を取得し返却
 */
func (a ContentfulAPIImpl) GetAllBreadsInformation(spaceID string, entryIDList []string, accessToken string) ([]model.Bread, error) {
	log.Println("GetAllBreadsInformation Start")
	var breads []model.Bread
	// 各entry idに対してリクエストを送る
	for _, entryID := range entryIDList {
		// リクエスト
		url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries/%s?access_token=%s", spaceID, entryID, accessToken)
		resp, httpError := http.Get(url)
		if httpError != nil {
			log.Println("failed to Get contentful:", httpError)
			return nil, httpError
		}
		defer resp.Body.Close()
		body, readBodyError := ioutil.ReadAll(resp.Body)
		if readBodyError != nil {
			fmt.Println("Error:", readBodyError)
			return nil, readBodyError
		}

		// 結果を配列に格納
		var res dto.GetBreadInfoResponse
		unmarshalError := json.Unmarshal(body, &res)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling json:", unmarshalError)
			return nil, unmarshalError
		}
		bread := model.Bread{
			ID:        res.Sys.ID,
			Name:      res.Fields.Name,
			CreatedAt: res.Sys.CreatedAt,
		}
		breads = append(breads, bread)
	}
	log.Println("GetAllBreadsInformation Success")
	return breads, nil
}
