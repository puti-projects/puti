package option

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

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
