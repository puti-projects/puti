package config

import (
	"fmt"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Config instance struct of config
type Config struct {
	vp *viper.Viper
}

var (
	Server *ServerConfig
	Safety *SafetyConfig
	Log    *LogConfig
	Db     *DbConfig
)

// NewConfig set up viper config and return a Config struct instance
func InitConfig(cfg string) error {
	// config file struct
	conf := struct {
		configPath string
	}{
		cfg,
	}

	// 实例化 viper
	vp := viper.New()

	// viper 配置
	if conf.configPath != "" {
		// 如果指定了配置文件，解析指定配置文件
		vp.SetConfigFile(conf.configPath)
	} else {
		// 如果没有指定配置文件，则解析默认配置文件
		vp.SetConfigName("config")
		vp.AddConfigPath("configs/")
	}
	vp.SetConfigType("yaml") // 设置配置文件格式为 YAML
	vp.AutomaticEnv()        // 读取匹配的环境变量
	vp.SetEnvPrefix("PUTI")  // 读取环境变量的前缀为 PUTI
	replacer := strings.NewReplacer(".", "_")
	vp.SetEnvKeyReplacer(replacer)

	// viper 解析配置文件
	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	c := &Config{vp}
	c.readConfig()
	c.watchConfigChange()

	return nil
}

// ReadConfig read config
func (c *Config) readConfig() error {
	err := c.readConfigSections("server", &Server)
	if err != nil {
		return err
	}

	err = c.readConfigSections("safety", &Safety)
	if err != nil {
		return err
	}

	err = c.readConfigSections("log", &Log)
	if err != nil {
		return err
	}

	err = c.readConfigSections("db", &Db)
	if err != nil {
		return err
	}

	return nil
}

// sections use for reload config when watching config file change
var sections = make(map[string]interface{})

// ReadConfig read config sections
func (c *Config) readConfigSections(k string, v interface{}) error {
	err := c.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

// ReloadAllSection reload all config sections
func (c *Config) reloadAllSection() error {
	for k, v := range sections {
		err := c.readConfigSections(k, v)
		if err != nil {
			zap.L().Error(fmt.Sprintf("reload all config sections failed, err:%s", err))
			return err
		}
	}

	zap.L().Info("already reload all config sections")
	return nil
}

// watchConfigChange watch config file change and reload all config
func (c *Config) watchConfigChange() {
	go func() {
		c.vp.WatchConfig()
		c.vp.OnConfigChange(func(e fsnotify.Event) {
			zap.L().Info("config file changed", zap.String("config-path", e.Name))
			_ = c.reloadAllSection()
		})
	}()
}
