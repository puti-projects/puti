package option

import (
	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// ListResponse is the options list response struct, for expecialy one setting type
type ListResponse struct {
	UserList []*service.UserInfo `json:"userList"`
}

// List get options data list by setting type
func List(c *gin.Context) {
	settingType := c.Query("settingType")
	if err := checkSettingType(settingType); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	options, err := service.GetOptionsByType(settingType)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, options)
}

func checkSettingType(settingType string) error {
	if settingType != "general" &&
		settingType != "property" &&
		settingType != "pubic-account" &&
		settingType != "github" &&
		settingType != "discuss" &&
		settingType != "media" &&
		settingType != "reading" &&
		settingType != "writing" {
		return errno.New(errno.ErrSettingType, nil)
	}

	return nil
}
