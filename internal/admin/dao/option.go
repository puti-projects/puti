package dao

import (
	"fmt"

	"github.com/puti-projects/puti/internal/model"
)

// UpdateOptions update options
func (d *Dao) UpdateOptions(options map[string]interface{}) error {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for optionName, optionValue := range options {
		option := &model.Option{OptionName: optionName}
		if err := option.GetByName(tx); err != nil {
			tx.Rollback()
			return err
		}

		// if need update
		if option.OptionValue != optionValue {
			option.OptionValue = fmt.Sprintf("%v", optionValue)
			if err := option.Save(tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

// GetOptionByName get option by name
func (d *Dao) GetOptionByName(optionName string) (*model.Option, error) {
	option := &model.Option{OptionName: optionName}
	if err := option.GetByName(d.db); err != nil {
		return nil, err
	}

	return option, nil
}

// GetAllOptions get all options by options name
func (d *Dao) GetAllOptions(optionNames []string) ([]*model.Option, error) {
	return model.GetAllOptions(d.db, optionNames)
}

// GetAutoLoadOptions get options need autoload
func (d *Dao) GetAutoLoadOptions() ([]*model.Option, error) {
	return model.GetAutoLoadOptions(d.db)
}
