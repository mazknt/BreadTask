package config

import (
	"os"
)

var AppConfig *Config

type Config struct {
	SpaceID          string
	EntryIDList      []string
	Accesstoken      string
	ProjectID        string
	CredentialOption string
}

func LoadConfig() {
	AppConfig = &Config{}

	var entryIDList []string
	entryIDList = append(entryIDList, os.Getenv("HONEY_SOY_CRAN"))
	entryIDList = append(entryIDList, os.Getenv("BLACK_SESAME_POTE"))
	entryIDList = append(entryIDList, os.Getenv("SHICHIMI_SALT_FOCACCIA"))

	AppConfig.SpaceID = os.Getenv("SPACE_ID")
	AppConfig.EntryIDList = entryIDList
	AppConfig.Accesstoken = os.Getenv("ACCESS_TOKEN")
	AppConfig.ProjectID = os.Getenv("PROJECT_ID")
	AppConfig.CredentialOption = os.Getenv("CREDENTIAL_OPTION")
}
