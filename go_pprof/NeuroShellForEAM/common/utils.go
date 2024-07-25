package common

import (
	"encoding/json"
)

func StringToMap(str string) (map[string]string, error) {
	if len(str) == 0 {
		return map[string]string{}, nil
	}
	var m map[string]string
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
