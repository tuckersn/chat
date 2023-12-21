package api

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/util"
)

var notesDirectory string = func() string {
	dirPath := util.Config.Notes.Directory
	if dirPath == "" {
		working_dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return path.Join(working_dir, "notes")
	}
	return dirPath
}()

func ValidateNotePath(notePath string) byte {
	if !db.NotePathRegex.MatchString(notePath) && fs.ValidPath(notePath) {
		return util.FILE_INVALID_PATH
	}
	fileSystemPath := path.Join(notesDirectory, notePath)
	fileSystemPath = path.Clean(fileSystemPath)
	file, err := os.Stat(fileSystemPath)
	if errors.Is(err, fs.ErrNotExist) {
		return util.FILE_NOT_FOUND
	}
	if err != nil {
		return util.FILE_UNKNOWN_ERROR
	}
	if file.IsDir() {
		return util.FILE_IS_DIR
	}
	return util.FILE_EXISTS
}

// HttpNoteGet godoc
// @Summary Returns the content of a Note
// @Schemes
// @Description Returns the content of a Note based on it's path
// @Tags Note API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/note/:path [get]
func HttpNoteGet(c *gin.Context) {

	notePath := c.Param("path")
	notePathStatus := ValidateNotePath(notePath)
	switch notePathStatus {
	case util.FILE_NOT_FOUND:
		c.JSON(404, gin.H{
			"message": "Not Found, does not exist",
			"path":    notePath,
		})
		return
	case util.FILE_INVALID_PATH:
		c.JSON(404, gin.H{
			"message": "Not Found, invalid path",
			"path":    notePath,
		})
		return
	case util.FILE_IS_DIR:
	case util.FILE_EXISTS:
		break
	default:
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"path":    notePath,
		})
		return
	}

	if notePathStatus == util.FILE_IS_DIR {
		files := make([]string, 0)
		err := fs.WalkDir(os.DirFS(notePath), ".", func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			files = append(files, path)
			return nil
		})
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error reading directory",
				"path":    notePath,
			})
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"path":  notePath,
			"files": files,
		})
	} else {
		content, err := os.ReadFile(notePath)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error reading file",
				"path":    notePath,
			})
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": "pong",
			"path":    notePath,
			"content": content,
		})
	}
}

// HttpNotePost godoc
// @Summary Writes a Note
// @Schemes
// @Description Writes the provided content to the Note path provided in the url
// @Tags Note API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /note/:path [post]
func HttpNotePost(c *gin.Context) {
}

// HttpPageDelete godoc
// @Summary Deletes a note
// @Schemes
// @Description Deletes a note based on it's path
// @Tags Note API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /note/:path [delete]
func HttpNoteDelete(c *gin.Context) {
}
