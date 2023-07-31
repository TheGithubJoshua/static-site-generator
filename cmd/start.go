/*
Copyright © 2023 Joshua Wheeler <joshuawhe@proton.me>

*/
package cmd

import (
  "net/http"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the webserver.",
	Long: `Starts the built-in webserver. Running on port 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
