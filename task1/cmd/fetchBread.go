/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"task1/task1/cmd/interface/api"
	"task1/task1/cmd/interface/repository"

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
		contentfulAPI := api.ContentfulAPIImpl{}
		firestore := repository.FirestoreImpl{}

		// パン情報の取得
		breadInformations, getBreadInformationError := contentfulAPI.GetBreadInformation()
		if getBreadInformationError != nil {
			log.Println("getBreadInformationError:", getBreadInformationError)
		}

		// パン情報をDBに保存
		for _, bread := range breadInformations {
			setBreadError := firestore.SetBread(bread)
			if setBreadError != nil {
				log.Println("setBreadError:", setBreadError)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchBreadCmd)
	loadEnvError := godotenv.Load()
	if loadEnvError != nil {
		log.Println("loadEnvError:", loadEnvError)
		return
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchBreadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchBreadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
