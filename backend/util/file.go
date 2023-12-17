package util

import (
	"errors"
	"os/user"
	"path"
	"regexp"
)

const (
	FILE_EXISTS        byte = 0
	FILE_NOT_FOUND     byte = 1
	FILE_IS_DIR        byte = 2
	FILE_INVALID_PATH  byte = 3
	FILE_UNKNOWN_ERROR byte = 255
)

var SlashesRegex *regexp.Regexp = regexp.MustCompile(`[/\\]+`)

func SmartPathSeg(pathSeg []string) []string {
	if pathSeg[0] == "~" {
		usr, err := user.Current()
		if err != nil {
			panic(errors.New("Error getting current user"))
		}
		pathSeg = append(SlashesRegex.Split(usr.HomeDir, -1), pathSeg[1:]...)
	}
	return pathSeg
}

func SmartPath(pathStr string) string {
	return path.Join(SmartPathSeg(SlashesRegex.Split(pathStr, -1))...)
}

func SmartPathFromStr(pathStr string) []string {
	return SmartPathSeg(SlashesRegex.Split(pathStr, -1))
}
