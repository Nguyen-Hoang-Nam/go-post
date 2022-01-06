package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Test put method",
	Long:  `Send PUT request to server with headers, params`,
	Run: func(cmd *cobra.Command, args []string) {
		client := http.Client{}

		url := ""
		if len(args) > 0 && args[0] != "" {
			url = args[0]
		} else {
			fmt.Printf("Url not found")
			return
		}

		req, err := http.NewRequest("PUT", url, nil)
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
	rootCmd.AddCommand(putCmd)
}
