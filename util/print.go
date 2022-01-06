package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

func Print(res *http.Response, isSortHeader bool) error {
	result := fmt.Sprintf("%s %s\n", res.Proto, res.Status)

	resHeaders := res.Header
	if isSortHeader {
		resHeaderKeys := make([]string, 0, len(resHeaders))
		for resHeaderKey := range resHeaders {
			resHeaderKeys = append(resHeaderKeys, resHeaderKey)
		}

		sort.Strings(resHeaderKeys)

		for _, resHeaderKey := range resHeaderKeys {
			result += fmt.Sprintf("%s: %s\n", resHeaderKey, resHeaders[resHeaderKey][0])
		}
	} else {
		for resHeaderKey, resHeaderValue := range resHeaders {
			result += fmt.Sprintf("%s: %s\n", resHeaderKey, resHeaderValue[0])
		}
	}

	result += "\n"

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Can not read response: %v\n", err)
	}

	var obj map[string]interface{}
	err = json.Unmarshal([]byte(resBody), &obj)
	if err != nil {
		result += string(resBody)
	} else {
		resBodyBeaty, err := json.MarshalIndent(obj, "", "    ")
		if err != nil {
			return fmt.Errorf("Response is not json: %v\n", err)
		}

		result += string(resBodyBeaty)
	}

	fmt.Printf("%s\n", result)
	return nil
}
