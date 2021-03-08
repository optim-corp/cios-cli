package utils

import (
	"encoding/json"

	"github.com/optim-kazuhiro-seida/go-advance-type/convert"

	ftil "github.com/optim-kazuhiro-seida/go-advance-type/file"

	"github.com/tidwall/gjson"
)

func StrToMap(in string) map[string]interface{} {
	return ByteToMap([]byte(in))
}
func ByteToMap(in []byte) map[string]interface{} {
	return gjson.ParseBytes(in).Value().(map[string]interface{})
}

func StructToJsonStr(object interface{}) (string, error) {
	result, err := json.Marshal(object)
	return string(result), err
}

func WriteJson(path string, data interface{}) error {
	str, err := StructToJsonStr(data)
	j, err := convert.IndentJson(str)
	if err != nil {
		return err
	}
	return ftil.Path(path).WriteFileAsString(j)
}
