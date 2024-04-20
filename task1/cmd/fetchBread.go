/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"task1/task1/cmd/config"
	"task1/task1/cmd/interface/api"
	"task1/task1/cmd/interface/repository"
	"task1/task1/cmd/usecase"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// fetchBreadCmd represents the fetchBread command
var fetchBreadCmd = &cobra.Command{
	Use:   "fetchBread",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("fetchBread called")
		fetchBreadService := usecase.FetchBreadServiceImpl{
			Firestore:     repository.FirestoreImpl{},
			ContentfulAPI: api.ContentfulAPIImpl{},
		}

		getBreadInfoFromContentfulAndSaveError := fetchBreadService.GetBreadInfoFromContentfulAndSave()
		if getBreadInfoFromContentfulAndSaveError != nil {
			log.Println("getBreadInfoFromContentfulAndSaveError:", getBreadInfoFromContentfulAndSaveError)
			return
		}
		log.Println("fetchBread successed")
	},
}
var AppConfig config.Config

func init() {
	rootCmd.AddCommand(fetchBreadCmd)
	loadEnvError := godotenv.Load()
	if loadEnvError != nil {
		log.Println("loadEnvError:", loadEnvError)
		return
	}
	config.LoadConfig()

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchBreadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchBreadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
