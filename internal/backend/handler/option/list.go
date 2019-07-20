package option

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/common/config"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/theme"

	"github.com/gin-gonic/gin"
)

// ListResponse option list and extra data if exist
type ListResponse struct {
	Options   map[string]string `json:"options"`
	ExtraData interface{}       `json:"extraData"`
}

// ListTheme the list of all theme
type ListTheme struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

// List get options data list by setting type
func List(c *gin.Context) {
	settingType := c.Query("settingType")
	if err := checkSettingType(settingType); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	// get options list
	options, err := service.GetOptionsByType(settingType)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	listResponse := &ListResponse{
		Options: options,
	}

	if "theme" == settingType {
		var themes []*ListTheme
		for _, theme := range theme.Themes {
			// thumbnail is exist
			var themeThumbnail string
			if exist, _ := utils.PathExists(config.StaticPath("theme/"+theme) + "/thumbnail.jpg"); exist == true {
				themeThumbnail = config.StaticPath("theme/"+theme) + "/thumbnail.jpg"
			} else {
				themeThumbnail = config.StaticPath("assets/images/") + "/image_default.png"
			}

			themeInfo := &ListTheme{
				Name:      theme,
				Thumbnail: themeThumbnail,
			}
			themes = append(themes, themeInfo)
		}
		listResponse.ExtraData = themes
	}

	Response.SendResponse(c, nil, listResponse)
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
