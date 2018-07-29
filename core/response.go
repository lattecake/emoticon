package core

import "encoding/json"

type resJson struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func Response(success bool, message string, data map[string]interface{}, other ...interface{}) (result []byte) {
	res := resJson{
		success,
		message,
		data,
	}

	result, _ = json.Marshal(res)

	return result
}
