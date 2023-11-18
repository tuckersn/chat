package api

import "github.com/gin-gonic/gin"

// HttpUserGet godoc
// @Summary creates a page
// @Schemes
// @Description creates a page based on it's path
// @Tags basic
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/:userId [get]
func HttpUserGet(r *gin.Context) {
}

// HttpUserGet godoc
// @Summary creates a page
// @Schemes
// @Description creates a page based on it's path
// @Tags basic
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/:userId [post]
func HttpUserUpdate(r *gin.Context) {
}

// HttpUserCreate godoc
// @Summary creates a page
// @Schemes
// @Description creates a page based on it's path
// @Tags basic
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user [post]
func HttpUserCreate(r *gin.Context) {
}

// HttpUserDelete godoc
// @Summary deletes a user
// @Schemes
// @Description deletes a user based on it's path
// @Tags basic
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/:userId [delete]
func HttpUserDelete(r *gin.Context) {
}
