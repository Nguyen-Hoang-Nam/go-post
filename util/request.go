package util

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

func Request(cmd *cobra.Command, method string, url string) error {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Url not valid: %v\n", err)
	}

	headers, err := cmd.Flags().GetStringSlice("header")
	if err != nil {
		return fmt.Errorf("Header not found: %v\n", err)
	}

	if len(headers) > 0 {
		for _, header := range headers {
			keyValue := strings.Split(header, ":")
			key := strings.Trim(keyValue[0], " ")
			value := strings.Trim(keyValue[1], " ")

			req.Header.Set(key, value)
		}
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Can not send request: %v\n", err)
	}

	err = Print(res)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}

	return nil
}
