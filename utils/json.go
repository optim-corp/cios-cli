package utils

import (
	"encoding/json"
)

func StructToJsonStr(object interface{}) (string, error) {
	result, err := json.Marshal(object)
	return string(result), err
}
