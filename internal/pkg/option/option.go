package option

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/pkg/logger"
)

// DefaultExpiration default expiration time of 2 hours for puti
var DefaultExpiration = 2 * time.Hour

// DefaultPurgesExpiration default purges expired items of 3 hours for puti
var DefaultPurgesExpiration = 3 * time.Hour

// Options cache
var Options = &optionCache{
	gocacheBody: gocache.New(DefaultExpiration, DefaultPurgesExpiration),
}

type optionCache struct {
	gocacheBody *gocache.Cache
}

// LoadOptions load default options
func LoadOptions() error {
	if err := getAutoLoadOptions(); err != nil {
		return err
	}

	logger.Info("options has been deployed successfully")

	return nil
}

// getAutoLoadOptions get options need to load
func getAutoLoadOptions() error {
	options, err := model.GetAutoLoadOptions()
	if err != nil {
		return err
	}

	for _, option := range options {
		Options.Put(option.OptionName, option.OptionValue)
	}

	return nil
}

// SetCache set option into cache
func (cache *optionCache) Put(optionName, optionValue string) {
	cache.gocacheBody.Set(optionName, optionValue, gocache.DefaultExpiration)
}

// Get one option value by optionName
func (cache *optionCache) Get(optionName string) string {
	optionValue, found := cache.gocacheBody.Get(optionName)
	if found {
		return optionValue.(string)
	}

	// If can not find the option by name in cache, get from db
	option, err := model.GetOption(optionName)
	if err != nil {
		logger.Errorf("getting option failed. Option name: %s. %s", optionName, err)
		return ""
	}

	// if the cache invalid is a autoload type option, reload all autoloaded options
	if option.Autoload == 1 {
		// option reload
		if err := getAutoLoadOptions(); err != nil {
			logger.Errorf("reload default options failed. %s", err)
			return ""
		}

		return option.OptionValue
	}

	// set cache
	Options.Put(option.OptionName, option.OptionValue)

	return option.OptionValue
}

// Delete an option from the cache. Does nothing if the option name is not in the cache as a key.
func (cache *optionCache) Delete(optionName string) {
	cache.gocacheBody.Delete(optionName)
}

// Flush delete all options from the cache
func (cache *optionCache) Flush() {
	cache.gocacheBody.Flush()
}

func (cache *optionCache) All() map[string](gocache.Item) {
	return cache.gocacheBody.Items()
}
