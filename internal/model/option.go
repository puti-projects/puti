package model

import (
	"gorm.io/gorm"
)

// Option site options
type Option struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	OptionName  string `gorm:"column:option_name;not null"`
	OptionValue string `gorm:"column:option_value;not null"`
	Autoload    uint64 `gorm:"column:autoload;not null"`
}

// TableName is the resource table name in db
func (o *Option) TableName() string {
	return "pt_option"
}

// GetByName get option by name
func (o *Option) GetByName(db *gorm.DB) error {
	if o.OptionName != "" {
		db = db.Where("`option_name` = ?", o.OptionName).First(&o)
	}

	return db.Error
}

// Save save option
func (o *Option) Save(db *gorm.DB) error {
	return db.Save(o).Error
}

// GetAll get all options by options name
func (o *Option) GetAll(db *gorm.DB, optionNames []string) ([]*Option, error) {
	options := make([]*Option, len(optionNames))
	if err := db.Where("option_name in (?)", optionNames).Find(&options).Error; err != nil {
		return nil, err
	}

	return options, nil
}

// GetAllAutoLoad get options need autoload
func (o *Option) GetAllAutoLoad(db *gorm.DB) ([]*Option, error) {
	var options []*Option
	if err := db.Where("`autoload` = 1").Find(&options).Error; err != nil {
		return options, err
	}

	return options, nil
}
