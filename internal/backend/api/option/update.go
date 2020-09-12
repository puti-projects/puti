package option

import (
	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
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
	u.Parms = make(map[string]interface{})
	if err := utils.BindJSONIntoMap(c, u.Parms); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// filter those params
	defaultOptionNames := service.GetDefaultOptionNamesByType(settingType)
	updateParamFilter(&u, defaultOptionNames)

	if err := service.UpdateOptions(u.Parms); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
	return
}

func updateParamFilter(u *service.OptionUpdateRequest, filterArr []string) error {
	for optionName := range u.Parms {
		optionNameExist := false
		for _, value := range filterArr {
			if optionName == value {
				optionNameExist = true
				break
			}
			continue
		}

		if !optionNameExist {
			delete(u.Parms, optionName)
		}
	}

	return nil
}
