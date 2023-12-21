package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/auth"
	"github.com/tuckersn/chatbackend/db"
)

// HttpLoginRecent godoc
// @Summary Gets your recent logins (up to 10)
// @Description gets a list of recent logins for provided token's associated user.
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
