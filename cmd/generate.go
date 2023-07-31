/*
Copyright © 2023 Joshua Wheeler <joshuawhe@proton.me>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the website files",
	Long:  `Generates the HTML, CSS and JS file of the website in the current directory. Dosen´t start the webserver.`,
	Run: func(cmd *cobra.Command, args []string) {
		os.Create("index.html")
    os.Create("sytle.css")
    os.Create("script.js")
    os.WriteFile("index.html", []byte(`<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>HTML 5 Boilerplate</title>
    <link rel="stylesheet" href="style.css">
  </head>
  <body>
	<script src="index.js"></script>
  </body>
</html>;`), 0666)
os.WriteFile("style.css", []byte(`html {
height: 100%;
}
body {
min-height: 100%;
}`), 0666)

    
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
