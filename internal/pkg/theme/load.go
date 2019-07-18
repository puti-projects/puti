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
	if err != nil {
		logger.Errorf("load theme error: %s", err)
	}

	dirName, err := f.Readdirnames(-1)
	if err != nil {
		logger.Errorf("load theme error: %s", err)
	}
	f.Close()

	for _, theme := range dirName {
		if "common" != theme {
			Themes = append(Themes, theme)
		}
	}

	logger.Info(fmt.Sprintf("loaded %d themes", len(Themes)), zap.Any("themes", Themes))
}
