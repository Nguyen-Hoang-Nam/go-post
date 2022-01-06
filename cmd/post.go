package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Nguyen-Hoang-Nam/go-post/util"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Test post method",
	Long:  `Send POST request to server with headers, params`,
	Run: func(cmd *cobra.Command, args []string) {
		client := http.Client{}

		url := ""
		if len(args) > 0 && args[0] != "" {
			url = args[0]
		} else {
			fmt.Printf("Url not found")
			return
		}

		req, err := http.NewRequest("POST", url, nil)
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

		util.Print(res)
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
}
