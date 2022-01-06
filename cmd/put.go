package cmd

import (
	"fmt"

	"github.com/Nguyen-Hoang-Nam/go-post/util"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Test put method",
	Long:  `Send PUT request to server with headers, params`,
	Run: func(cmd *cobra.Command, args []string) {
		url := ""
		if len(args) > 0 && args[0] != "" {
			url = args[0]
		} else {
			fmt.Printf("Url not found")
			return
		}

		err := util.Request(cmd, "DELETE", url)
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
