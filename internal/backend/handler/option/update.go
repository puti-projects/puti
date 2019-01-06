package option

import (
	"fmt"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/errno"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// UpdateRequest update request include params which are dynamitly
type UpdateRequest struct {
	Parms map[string]interface{}
}

// Update update options by setting type
func Update(c *gin.Context) {
	log.Info("Option update function called.", lager.Data{"X-Request-Id": utils.GetReqID(c)})

	// Get setting type
	settingType := c.Query("settingType")

	if err := checkSettingType(settingType); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	var u UpdateRequest
	u.Parms = make(map[string]interface{})
	if err := utils.BindJSONIntoMap(c, u.Parms); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// filter those params
	defaultOptionNames := service.GetDefaultOptionsByType(settingType)
	u.paramFilter(defaultOptionNames)

	if err := service.UpdateOptions(u.Parms); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	// update options' cache
	for optionName, optionValue := range u.Parms {
		optionCache.Options.Put(optionName, fmt.Sprintf("%v", optionValue))
	}

	Response.SendResponse(c, nil, nil)
	return
}

func (u *UpdateRequest) paramFilter(filterArr []string) error {
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
