/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Test delete method",
	Long:  `Send DELETE request to server with headers, params`,
	Run: func(cmd *cobra.Command, args []string) {
		client := http.Client{}

		url := ""
		if len(args) > 0 && args[0] != "" {
			url = args[0]
		} else {
			fmt.Printf("Url not found")
			return
		}

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Printf("Url not valid: %v", err)
			return
		}

		headers, err := cmd.Flags().GetStringSlice("header")
		if err != nil {
			fmt.Printf("Header not found: %v", err)
			return
		}

		if len(headers) > 0 {
			for _, header := range headers {
				keyValue := strings.Split(header, ":")
				key := strings.Trim(keyValue[0], " ")
				value := strings.Trim(keyValue[1], " ")

				req.Header.Set(key, value)
			}
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("Can not send request: %v", err)
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("Can not read response: %v", err)
			return
		}

		result := string(body)
		fmt.Printf(result)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
