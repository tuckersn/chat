package api

import (
	"io"

	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/db"
)

type UserGetResponse struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

func ValidateUsername(username string) bool {
	return db.UserNameRegex.MatchString(username)
}

func ValidateDisplayName(displayName string) bool {
	return db.DisplayNameRegex.MatchString(displayName)
}

// HttpUserGet godoc
// @Summary get a user's info
// @Schemes
// @Description retrieves information about the user specified by the username
// @Tags User API
// @Accept json
// @Produce json
// @Param username path string true "Username of the user to retrieve"
// @Success 200 {object} UserGetResponse "User information"
// @Failure 400 {string} string "Invalid username"
// @Failure 500 {string} string "User not found"
// @Router /api/user/id/:username [get]
func HttpUserGet(c *gin.Context) {
	username := c.Param("username")
	if !(ValidateUsername(username)) {
		c.JSON(400, gin.H{
			"error": "Invalid username",
		})
		return
	}

	user, err := db.GetUser(username)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(200, user)
}

// HttpUserGet godoc
// @Summary creates a page
// @Schemes
// @Description creates a page based on it's path
// @Tags User API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/user/id/:userId [post]
func HttpUserUpdate(r *gin.Context) {

}

type UserCreateRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Admin       *bool  `json:"admin"`
}

func UserCreate(username string, displayName string) (db.RecordUser, error) {
	user, err := db.InsertUser(username, displayName)
	return user, err
}

// HttpUserCreate godoc
// @Summary create a new user
// @Schemes
// @Description creates a new user
// @Tags User API
// @Accept json
// @Produce json
// @Param body body UserCreateRequest true "User creation request"
// @Success 200 {string} HttpUserCreateRequest
// @Router /api/user [post]
func HttpUserCreate(c *gin.Context) {
	input, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	var username string
	username, err = jsonparser.GetString(input, "username")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid username",
		})
		return
	}
	if !(ValidateUsername(username)) {
		c.JSON(400, gin.H{
			"error": "Invalid username",
		})
		return
	}

	var displayName string
	displayName, err = jsonparser.GetString(input, "displayName")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid display name",
		})
		return
	}
	if !(ValidateDisplayName(displayName)) {
		c.JSON(400, gin.H{
			"error": "Invalid display name",
		})
		return
	}

	user, err := UserCreate(username, displayName)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, user)
}

// HttpUserDelete godoc
// @Summary deletes a user
// @Schemes
// @Description deletes a user based on it's path
// @Tags User API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/user/id/:username [delete]
func HttpUserDelete(r *gin.Context) {

}
