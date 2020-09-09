package config

import "path/filepath"

// router path config
const (
	PathRoot     = "/"
	PathBackend  = "/admin"
	PathAPI      = "/api"
	PathArticle  = "/article"
	PathCategory = "/category"
	PathTag      = "/tag"
	PathSubject  = "/subject"
	PathArchives = "/archives"

	PathRSS     = "/rss"
	PathSiteMap = "/sitemap.xml"
	PathRobots  = "/robots.txt"
	PathFavicon = "/favicon.ico"
)

// static path config
const (
	StaticPathRoot  = ""
	StaticPathTheme = "theme"
)

// upload path config
const (
	// UploadPath defines the media file save path uri
	UploadPath = "/uploads/"
	// savePath defines the use avatar saving path
	UploadUserAvatarPath = "/uploads/users/"
)

// StaticPath change path to static path base on the StaticPathRoot
func StaticPath(Path string) string {
	return filepath.ToSlash(filepath.Join(StaticPathRoot, Path))
}
