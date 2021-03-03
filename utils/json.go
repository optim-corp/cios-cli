package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/mitchellh/mapstructure"
	"github.com/tidwall/gjson"
)

func StrToMap(in string) map[string]interface{} {
	return ByteToMap([]byte(in))
}
func ByteToMap(in []byte) map[string]interface{} {
	return gjson.ParseBytes(in).Value().(map[string]interface{})
}
func DecodeJson(in string, stc interface{}) error {
	return DecodeJsonFromBytes([]byte(in), stc)
}
func DecodeJsonFromBytes(in []byte, stc interface{}) error {
	tmp := gjson.ParseBytes(in).Value().(interface{})
	config := &mapstructure.DecoderConfig{
		TagName:  "json",
		Result:   stc,
		Metadata: nil,
	}
	decoder, err := mapstructure.NewDecoder(config)
	err = decoder.Decode(tmp)
	return err
}
func StructToJsonStr(object interface{}) (string, error) {
	result, err := json.Marshal(object)
	return string(result), err
}
func IndentJson(object string) (string, error) {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(object), "", "  ")
	if err != nil {
		return "", err
	}
	return buf.String(), err
}
func WriteJson(path string, data interface{}) error {
	str, err := StructToJsonStr(data)
	j, err := IndentJson(str)
	if err != nil {
		Log.Error(err.Error())
		return err
	}
	return Path(path).WriteFileAsString(j)
}
func ReadJsonMap(path string) (interface{}, bool) {
	byts, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, false
	}
	result, ok := gjson.ParseBytes(byts).Value().(interface{})
	return result, ok
}
func LoadJsonStruct(path string, st interface{}) error {
	jMap, ok := ReadJsonMap(path)
	if !ok {
		return errors.New("Cant Map")
	}
	config := &mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &st,
	}
	decoder, _ := mapstructure.NewDecoder(config)
	return decoder.Decode(jMap)
}
func GetKeys(mymap map[string]interface{}) []string {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}

func GetKeys2(mymap map[string]map[string]interface{}) []string {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}
	return keys
}
