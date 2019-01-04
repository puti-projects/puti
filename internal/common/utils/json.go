package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

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
