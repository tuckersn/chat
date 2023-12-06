package api

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/gitlab"
	"github.com/tuckersn/chatbackend/util"
)

func HttpLogin(c *gin.Context) {

}

var gitLabLoginRequests = make(map[string]string)

// HttpGitLabLoginRedirect godoc
// @Summary handles GitLab oauth2 callback
// @Schemes
// @Description where the user's browser is sent by GitLab after completing the oauth2 flow
// @Tags Login
// @Accept json
// @Produce json
// @Success 302
// @Router /login/gitlab [get]
func HttpGitLabLoginRedirect(c *gin.Context) {

	csrfToken := util.RandomString(32, []string{util.ALPHABET_URL_SAFE})
	verifier := util.RandomString(128, []string{util.ALPHABET_URL_SAFE})
	h := sha256.New()
	h.Write([]byte(verifier))
	challenge := base64.StdEncoding.EncodeToString(h.Sum(nil))

	protocol := "https://"
	if c.Request.TLS == nil {
		protocol = "http://"
	}
	url := fmt.Sprintf(
		"%s/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=%s&state=%s&scope=%s&code_challenge=%s&code_challenge_method=%s",
		gitlab.HostURL(),
		gitlab.GetAppId(),
		protocol+c.Request.Host+c.Request.URL.Path+"/receiveToken",
		"code",
		csrfToken,
		"read_user+api+profile",
		challenge,
		"S256",
	)
	c.Redirect(302, url)
}

type GitLabReceiveTokenReponse struct {
}

// HttpGitLabLoginReceiveToken godoc
// @Summary handles GitLab oauth2 callback
// @Schemes
// @Description where the user's browser is sent by GitLab after completing the oauth2 flow
// @Tags Login
// @Accept json
// @Param code query string true "The code returned by GitLab"
// @Param state query string true "The state returned by GitLab"
// @Produce json
// @Success 200 {object} GitLabReceiveTokenReponse
// @Router /login/gitlab/receiveToken [get]
func HttpGitLabLoginReceiveToken(c *gin.Context) {
	code := c.Query("code")
	//csrfToken := c.Query("state")
	protocol := "https://"
	if c.Request.TLS == nil {
		protocol = "http://"
	}
	// already contains the /receiveToken ending
	access_token, err := gitlab.GetAccessToken(code, protocol+c.Request.Host+c.Request.URL.Path)
	if err != nil {
		//TODO: handle this better
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(access_token)
	c.JSON(200, gin.H{"access_token": access_token})
}
