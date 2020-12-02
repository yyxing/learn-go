package util

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func ModelToMap(model interface{}) map[string]interface{} {
	var resultMap map[string]interface{}
	data, err := Marshal(model)
	if err != nil {
		logrus.Errorf("json Marshal failed %s", err)
	}
	if err = Unmarshal(data, &resultMap); err != nil {
		logrus.Errorf("json Unmarshal failed %s", err)
	}
	return resultMap
}

func Marshal(model interface{}) ([]byte, error) {
	json, err := json.Marshal(model)
	return json, err
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}
