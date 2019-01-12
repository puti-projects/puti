package model

// OptionModel site options
type OptionModel struct {
	ID          uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	OptionName  string `gorm:"column:option_name;not null"`
	OptionValue string `gorm:"column:option_value;not null"`
	Autoload    uint64 `gorm:"column:autoload;not null"`
}

// TableName is the resource table name in db
func (c *OptionModel) TableName() string {
	return "pt_option"
}

// GetOption get one option by name
func GetOption(optionName string) (*OptionModel, error) {
	option := &OptionModel{}
	result := DB.Local.Where("option_name = ?", optionName).First(&option)
	return option, result.Error
}

// GetOptionsByNames select potions by options' name
func GetOptionsByNames(optionNames []string) ([]*OptionModel, error) {
	options := make([]*OptionModel, len(optionNames))

	if err := DB.Local.Where("option_name in (?)", optionNames).Find(&options).Error; err != nil {
		return options, err
	}

	return options, nil
}

// GetAutoLoadOptions get options need autoload
func GetAutoLoadOptions() ([]*OptionModel, error) {
	var options []*OptionModel

	if err := DB.Local.Where("autoload = 1").Find(&options).Error; err != nil {
		return options, err
	}

	return options, nil
}
