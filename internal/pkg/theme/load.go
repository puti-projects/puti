package theme

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/puti-projects/puti/internal/common/config"
	"github.com/puti-projects/puti/internal/pkg/logger"

	"go.uber.org/zap"
)

// Themes saves all theme names.
var Themes []string

// LoadInstalled load all installed themes
func LoadInstalled() {
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
		// if it is a dir (it may be a file)
		f, _ := os.Stat(filepath.ToSlash(config.StaticPathTheme + "/" + theme))
		isDir := f.IsDir()

		if "common" != theme && true == isDir {
			Themes = append(Themes, theme)
		}
	}

	logger.Info(fmt.Sprintf("loaded %d themes", len(Themes)), zap.Any("themes", Themes))
}
