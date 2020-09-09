package utils

import (
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

// pictureMimeType file type map
var fileMimeType = map[string]string{
	"image/jpeg":      "picture",
	"image/png":       "picture",
	"image/gif":       "picture",
	"text/plain":      "text",
	"video/mp4":       "video",
	"video/3gpp":      "video",
	"video/x-msvideo": "video",
	"video/x-ms-wmv":  "video",
}

// GetFileExt returns the file name extension
// The extension is the suffix beginning at the final dot
func GetFileExt(file *multipart.FileHeader) string {
	var ext string
	// get by Ext func first
	ext = filepath.Ext(file.Filename)
	if ext == "" {
		// get by content-type
		typ := file.Header.Get("Content-Type")
		exts, _ := mime.ExtensionsByType(typ)
		if 0 < len(exts) {
			ext = exts[0]
		} else {
			ext = "." + strings.Split(typ, "/")[1]
		}
	}

	return ext
}

// GetFileMimeTypeByExt returns the file mine-type by ext
func GetFileMimeTypeByExt(ext string) string {
	return mime.TypeByExtension(ext)
}

// GetFileType returns file type in word that defined
func GetFileType(mimeType string) string {
	fileType, exists := fileMimeType[mimeType]
	if exists {
		return fileType
	}

	return "others"
}

// PathExists check the path if exist
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CheckPathAndCreate check path if exist, fi not exist, make the path dir
func CheckPathAndCreate(path string) error {
	coverPathExist, err := PathExists(path)
	if err != nil {
		return err
	}

	if !coverPathExist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
