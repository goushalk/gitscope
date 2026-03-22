package logic

import (
	"encoding/json"
)

func JsonOutput(data any) (string, error) {
	JsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(JsonData), nil
}
