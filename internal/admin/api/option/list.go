package option

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
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

	svc := service.New(c.Request.Context())
	rsp, err := svc.ListOption(settingType)
	if err != nil {
		api.SendResponse(c, err, nil)
	}

	api.SendResponse(c, nil, rsp)
}

func checkSettingType(settingType string) error {
	for _, v := range service.OptionSettingTypeMap {
		if v == settingType {
			return nil
		}
	}
	return errno.New(errno.ErrSettingType, nil)
}
