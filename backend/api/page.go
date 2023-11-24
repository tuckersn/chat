package api

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
)

var pagesDirectory string = func() string {
	env := os.Getenv("PAGES_DIRECTORY")
	if env == "" {
		working_dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		env = path.Join(working_dir, "pages")
		return env
	}
	return env
}()

// HttpPageGet godoc
// @Summary returns the content of a page
// @Schemes
// @Description returns the content of a page based on it's path
// @Tags basic
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /page/:path [get]
func HttpPageGet(c *gin.Context) {

	filePath := c.Param("path")
	if !db.PagePathRegex.MatchString(filePath) && fs.ValidPath(filePath) {
		c.JSON(404, gin.H{
			"message": "Not Found, invalid path",
		})
		return
	}
	fileSystemPath := path.Join(pagesDirectory, filePath)
	fileSystemPath = path.Clean(fileSystemPath)
	file, err := os.Stat(fileSystemPath)

	if errors.Is(err, fs.ErrNotExist) {
		c.JSON(404, gin.H{
			"message": "Not Found, does not exist",
			"path":    filePath,
		})
		return
	}
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
			"path":    filePath,
		})
		log.Println(err)
		return
	}

	if file.IsDir() {
		// returns a list of all files in the directory
		files := make([]string, 0)
		err := fs.WalkDir(os.DirFS(filePath), ".", func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			files = append(files, path)
			return nil
		})
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error reading directory",
				"path":    filePath,
			})
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"path":  filePath,
			"files": files,
		})
	} else {
		content, err := os.ReadFile(fileSystemPath)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error reading file",
				"path":    filePath,
			})
			log.Println(err)
			return
		}
		c.JSON(200, gin.H{
			"message": "pong",
			"path":    filePath,
			"content": content,
		})
	}
}

// HttpPagePost godoc
// @Summary creates a page
// @Schemes
// @Description creates a page based on it's path
// @Tags basic
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /page/:path [post]
func HttpPagePost(c *gin.Context) {
}

// HttpPageDelete godoc
// @Summary deletes a page
// @Schemes
// @Description deletes a page based on it's path
// @Tags basic
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /page/:path [delete]
func HttpPageDelete(c *gin.Context) {
}
