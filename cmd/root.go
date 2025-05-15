package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lunch-cli",
	Short: "lunch-cli is a cli tool that fetches the menu",
	Long:  "lunch-cli is a cli tool that fetches the menu for the location in Skanderborgvej 190",
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://shop.foodandco.dk/api/WeeklyMenu?restaurantId=1234&languageCode=da-DK")

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing lunch-cli '%s'\n", err)
		os.Exit(1)
	}
}
