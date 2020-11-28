package theme

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"

	"go.uber.org/zap"
)

// Theme theme information for loading theme
type Theme struct {
	Name           string
	FaviconExist   bool
	ThumbnailExist bool
	RobotsExist    bool
}

// Themes saves all theme information
// TODO：改为并发安全。并发加载主题
var Themes map[string]*Theme

// LoadInstalled load all installed themes
func LoadInstalled() {
	Themes = make(map[string]*Theme, 0)

	f, err := os.Open(filepath.ToSlash(config.StaticPathTheme))
	defer func() {
		if err := f.Close(); err != nil {
			logger.Warnf("load theme error when closing dir: %s", err)
		}
	}()
	if err != nil {
		logger.Errorf("load theme error: %s", err)
	}

	dirName, err := f.Readdirnames(-1)
	if err != nil {
		logger.Errorf("load theme error: %s", err)
	}

	for _, theme := range dirName {
		if "common" == theme {
			continue
		}

		// if it is a dir (it may be a file); it should be a theme dir
		themePath := config.StaticPath(config.StaticPathTheme + "/" + theme)
		f, _ := os.Stat(themePath)
		if isDir := f.IsDir(); true == isDir {
			// it is a theme; init
			t := &Theme{
				Name: theme,
			}

			// check favicon
			if exist, _ := utils.PathExists(themePath + "/favicon.ico"); exist {
				t.FaviconExist = true
			}
			// check thumbnail
			if exist, _ := utils.PathExists(themePath + "/thumbnail.jpg"); exist {
				t.ThumbnailExist = true
			}
			// check robots.txt
			if exist, _ := utils.PathExists(themePath + "/robots.txt"); exist {
				t.RobotsExist = true
			}

			Themes[theme] = t
		}
	}

	logger.Info(fmt.Sprintf("loaded %d themes", len(Themes)), zap.Any("themes", Themes))
}
