package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"task1/task1/cmd/model"
)

type ContentfulAPI interface {
	GetBreadInformation() ([]model.Bread, error)
}

type ContentfulAPIImpl struct{}

/**
 * Contentfulからパンの情報を取得し返却
 */
func (a ContentfulAPIImpl) GetBreadInformation() ([]model.Bread, error) {

	// パラメータの用意
	space_id := os.Getenv("SPACE_ID")
	entry_id_list := []string{}
	entry_id_list = append(entry_id_list, os.Getenv("HONEY_SOY_CRAN"))
	entry_id_list = append(entry_id_list, os.Getenv("BLACK_SESAME_POTE"))
	entry_id_list = append(entry_id_list, os.Getenv("SHICHIMI_SALT_FOCACCIA"))
	access_token := os.Getenv("ACCESS_TOKEN")

	var breads []model.Bread
	// 各entry idに対してリクエストを送る
	for _, entry_id := range entry_id_list {
		// リクエスト
		url := fmt.Sprintf("https://cdn.contentful.com/spaces/%s/entries/%s?access_token=%s", space_id, entry_id, access_token)
		resp, httpError := http.Get(url)
		if httpError != nil {
			log.Println("failed to Get contentful:", httpError)
			return []model.Bread{}, httpError
		}
		defer resp.Body.Close()
		body, readBodyError := ioutil.ReadAll(resp.Body)
		if readBodyError != nil {
			fmt.Println("Error:", readBodyError)
			return []model.Bread{}, readBodyError
		}

		// 結果を配列に格納
		var res model.ContentfulResponse
		unmarshalError := json.Unmarshal(body, &res)
		if unmarshalError != nil {
			fmt.Println("Error unmarshalling json:", unmarshalError)
			return []model.Bread{}, unmarshalError
		}
		bread := model.Bread{
			ID: res.Sys.ID,
			BreadInfo: model.BreadInfo{
				Name:      res.Fields.Name,
				CreatedAt: res.Sys.CreatedAt,
			},
		}
		breads = append(breads, bread)
	}
	return breads, nil
}
