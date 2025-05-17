package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"io"
	"encoding/json"
)

var rootCmd = &cobra.Command{
	Use:   "lunch-cli",
	Short: "lunch-cli is a cli tool that fetches the menu",
	Long:  "lunch-cli is a cli tool that fetches the menu for the location in Skanderborgvej 190",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://shop.foodandco.dk/api/WeeklyMenu?restaurantId=1234&languageCode=da-DK")
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var jsonData interface{}
		if err := json.Unmarshal(body, &jsonData); err != nil {
			panic(err)
		}
		prettyJson, err := json.MarshalIndent(jsonData, ""," ")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(prettyJson))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
