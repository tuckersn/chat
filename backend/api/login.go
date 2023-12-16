package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/auth"
	"github.com/tuckersn/chatbackend/db"
)

// HttpLoginRecent godoc
// @Summary gets a list of recent logins for the user's account
// @Description gets a list of recent logins for the user's account
// @Tags Login
// @Accept json
// @Produce json
// @Success 200
// @Router /login/recent [get]
func HttpLoginRecent(r *gin.Context) {
	user, err := auth.HttpAuthResponseHandled(r)
	if err != nil {
		return // error handled in auth.HttpAuthWithResponse
	}

	logins, err := db.GetLoginsRecent(user.UserId, 10)
	if err != nil {
		panic(err)
	}

	r.JSON(200, logins)
}
