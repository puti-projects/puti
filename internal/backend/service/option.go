package service

import (
	"fmt"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
)

// GetOptionsByType get default options by setting type
func GetOptionsByType(settingType string) (map[string]string, error) {
	optionNames := GetDefaultOptionsByType(settingType)
	options, err := model.GetOptionsByNames(optionNames)
	if err != nil {
		return nil, err
	}

	optionsMap := make(map[string]string)

	for _, option := range options {
		optionsMap[option.OptionName] = option.OptionValue
	}

	if "theme" == settingType {

	}

	return optionsMap, nil
}

// UpdateOptions update options
func UpdateOptions(options map[string]interface{}) error {
	tx := db.DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for optionName, optionValue := range options {
		oldOption, err := model.GetOption(optionName)
		if err != nil {
			tx.Rollback()
			return err
		}

		// if need update
		if oldOption.OptionValue != optionValue {
			oldOption.OptionValue = fmt.Sprintf("%v", optionValue)

			// update it
			if err = tx.Model(&model.OptionModel{}).Save(oldOption).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// GetDefaultOptionsByType get default setting's option name
func GetDefaultOptionsByType(settingType string) []string {
	var optionNames []string
	switch settingType {
	case "general":
		optionNames = []string{"blog_name", "blog_description", "site_url", "admin_email", "users_can_register", "timezone_string", "site_language"}
	case "property":
		optionNames = []string{"site_description", "site_keywords", "footer_copyright"}
	case "menu":
		optionNames = []string{""}
	case "theme":
		optionNames = []string{"current_theme"}
	case "pubic-account":
		optionNames = []string{""}
	case "github":
		optionNames = []string{""}
	case "discuss":
		optionNames = []string{"article_comment_status", "page_comment_status", "comment_need_register", "show_comment_page", "comment_per_page", "comment_page_first", "comment_page_top", "comment_before_show", "show_avatar"}
	case "media":
		optionNames = []string{"image_thumbnail_width", "image_thumbnail_height", "image_medium_width", "image_medium_height", "image_large_width", "image_large_height"}
	case "reading":
		optionNames = []string{"show_on_front", "show_on_front_page", "posts_per_page", "open_XML"}
	case "writing":
		optionNames = []string{"default_category", "default_link_category"}
	default:
		optionNames = []string{}
	}

	return optionNames
}
