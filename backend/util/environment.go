package util

import (
	"errors"
	"os"
	"os/user"
	"path"
	"regexp"
)

var _storageDirInitialized = false
var StorageDir []string = []string{}
var SlashesRegex *regexp.Regexp = regexp.MustCompile(`[/\\]+`)

func GetStorageDir(pathStr string) string {
	if !_storageDirInitialized {
		StorageDir = SlashesRegex.Split(os.Getenv("CR_STORAGE_DIR"), -1)
		if StorageDir[0] == "~" {
			usr, err := user.Current()
			if err != nil {
				panic(errors.New("Error getting current user"))
			}
			StorageDir = append(SlashesRegex.Split(usr.HomeDir, -1), StorageDir[1:]...)
		}
	}
	return path.Join(append(StorageDir, pathStr)...)
}

func CreateStorageDirectoryIfNotExists() {
	if os.Getenv("CR_STORAGE_DIR") == "" {
		panic(errors.New("CR_STORAGE_DIR not set"))
	}
	if _, err := os.Stat(GetStorageDir("")); os.IsNotExist(err) {
		err := os.MkdirAll(GetStorageDir(""), 0755)
		if err != nil {
			panic(err)
		}
	}
}

func IsMainNode() bool {
	return os.Getenv("MAIN_NODE") == "true"
}

func GetRedirectBaseUrl() string {
	return os.Getenv("CR_REDIRECT_URL")
}
