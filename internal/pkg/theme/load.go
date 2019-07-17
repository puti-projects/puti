// package theme

// import (
// 	"os"
// 	"path/filepath"
// )

// // LoadInstalled laod all installed themes
// func LoadInstalled() {

// 	f, _ := os.Open(filepath.ToSlash(filepath.Join(model.Conf.StaticRoot, "theme/x")))
// 	names, _ := f.Readdirnames(-1)
// 	f.Close()

// 	for _, name := range names {
// 		if !util.IsNumOrLetter(rune(name[0])) {
// 			continue
// 		}

// 		Themes = append(Themes, name)
// 	}

// 	logger.Debugf("loaded [%d] themes", len(Themes))
// }
