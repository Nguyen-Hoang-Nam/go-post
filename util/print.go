package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Print(res *http.Response) {
	result := fmt.Sprintf("%s %s\n", res.Proto, res.Status)

	for resHeaderKey, resHeaderValue := range res.Header {
		result += fmt.Sprintf("%s: %s\n", resHeaderKey, resHeaderValue[0])
	}

	result += "\n"

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Can not read response: %v", err)
		return
	}

	var obj map[string]interface{}
	err = json.Unmarshal([]byte(resBody), &obj)
	if err != nil {
		result += string(resBody)
	} else {
		resBodyBeaty, err := json.MarshalIndent(obj, "", "    ")
		if err != nil {
			fmt.Printf("Response is not json: %v", err)
			return
		}

		result += string(resBodyBeaty)
	}

	fmt.Printf("%s\n", result)
}
