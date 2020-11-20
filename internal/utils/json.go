package utils

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// BindJSONIntoMap decodes json body to map, skips fields which are not in allowedFields
// Example usage:
// u.Parms = make(map[string]interface{})
// err := utils.BindJSONIntoMap(c, u.Parms)
func BindJSONIntoMap(context *gin.Context, obj map[string]interface{}) error {
	if context.Request.Method == "GET" {
		return nil
	}

	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}

// JSON2Map convert JSON to map
func JSON2Map(j []byte, m *map[string]interface{}) error {
	return json.Unmarshal(j, &m)
}

// Map2JSON convert map to JSON
func Map2JSON(m interface{}) ([]byte, error) {
	return json.Marshal(&m)
}
