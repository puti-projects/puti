package option

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List option list handler
func List(c *gin.Context) {
	settingType := c.Query("settingType")

	if err := checkSettingType(settingType); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	rsp, err := service.ListOption(settingType)
	if err != nil {
		api.SendResponse(c, err, nil)
	}

	api.SendResponse(c, nil, rsp)
}

func checkSettingType(settingType string) error {
	if settingType != "general" &&
		settingType != "property" &&
		settingType != "theme" &&
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
