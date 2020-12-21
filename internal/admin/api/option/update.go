package option

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
)

// Update update options by setting type
func Update(c *gin.Context) {
	// Get setting type
	settingType := c.Query("settingType")

	if err := checkSettingType(settingType); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	var u service.OptionUpdateRequest
	u.Params = make(map[string]interface{})
	if err := utils.BindJSONIntoMap(c, u.Params); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	svc := service.New(c.Request.Context())
	// filter those params
	defaultOptionNames := svc.GetDefaultOptionNamesByType(settingType)
	if err := paramFilter(&u, defaultOptionNames); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	if err := svc.UpdateOptions(u.Params); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
	return
}

func paramFilter(u *service.OptionUpdateRequest, filterArr []string) error {
	for optionName := range u.Params {
		optionNameExist := false
		for _, value := range filterArr {
			if optionName == value {
				optionNameExist = true
				break
			}
			continue
		}

		if !optionNameExist {
			delete(u.Params, optionName)
		}
	}

	return nil
}
