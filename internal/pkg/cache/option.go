package cache

import (
	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
)

type optionCache struct {
	cacheBody *Cache
	dao       *dao.Dao
}

// Options option instance
var Options *optionCache

// LoadOptions load default options
func LoadOptions() error {
	Options = &optionCache{
		cacheBody: GetInstance(),
		dao:       dao.New(db.Engine),
	}

	if err := getAutoLoadOptions(); err != nil {
		return err
	}

	return nil
}

// getAutoLoadOptions get options need to load
func getAutoLoadOptions() error {
	options, err := Options.dao.GetAutoLoadOptions()
	if err != nil {
		return err
	}

	for _, option := range options {
		if err := Options.Put(setOptionKeyPrefix(option.OptionName), option.OptionValue); err != nil {
			return err
		}
	}
	return nil
}

// setOptionKeyPrefix before set cache, add key prefix for option
func setOptionKeyPrefix(optionName string) string {
	return config.CacheOptionPrefix + optionName
}

// SetCache set option into cache
func (oc *optionCache) Put(optionName, optionValue string) error {
	if err := oc.cacheBody.SetCache(setOptionKeyPrefix(optionName), optionValue); err != nil {
		return err
	}
	return nil
}

// Get one option value by optionName
func (oc *optionCache) Get(optionName string) string {
	optionValue, found := oc.cacheBody.GetCache(setOptionKeyPrefix(optionName))
	if found {
		return optionValue
	}

	// If can not find the option by name in cache, get from db
	option, err := oc.dao.GetOptionByName(optionName)
	if err != nil {
		logger.Errorf("getting option failed. Option name: %s. %s", optionName, err)
		return ""
	}

	// if the cache invalid is a autoed load type option, reload all autoed load options
	if option.Autoload == 1 {
		// option reload
		if err := getAutoLoadOptions(); err != nil {
			logger.Errorf("reload default options failed. %s", err)
			return ""
		}

		return option.OptionValue
	}

	// set cache
	if err := Options.Put(setOptionKeyPrefix(option.OptionName), option.OptionValue); err != nil {
		logger.Errorf("error when setting option to cache. setting key: %s. err: %s", option.OptionName, err)
	}

	return option.OptionValue
}

// Delete an option from the cache. Does nothing if the option name is not in the cache as a key.
func (oc *optionCache) Delete(optionName string) error {
	return oc.cacheBody.DeleteCache(setOptionKeyPrefix(optionName))
}

// Flush delete all options from the cache
func (oc *optionCache) Flush() error {
	return oc.cacheBody.Flush()
}
