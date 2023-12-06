package api

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tuckersn/chatbackend/gitlab"
	"github.com/tuckersn/chatbackend/util"
)

type gitlabLoginRequest struct {
	csrfToken     string
	verifier      string
	timeOfRequest time.Time
}

// TODO: use an async go routine to cull this every 5 minutes if it's older than 5 minutes
// TODO: move to SQL
var gitLabLoginRequests = make(map[string]gitlabLoginRequest)

// HttpLoginGitlabRedirect godoc
// @Summary handles GitLab oauth2 callback
// @Schemes
// @Description where the user's browser is sent by GitLab after completing the oauth2 flow
// @Tags Login
// @Accept json
// @Produce json
// @Success 302
// @Router /login/gitlab [get]
func HttpLoginGitlabRedirect(c *gin.Context) {

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
		protocol+c.Request.Host+c.Request.URL.Path+"/receive",
		"code",
		csrfToken,
		"read_user",
		challenge,
		"S256",
	)
	gitLabLoginRequests[csrfToken] = gitlabLoginRequest{
		csrfToken:     csrfToken,
		verifier:      verifier,
		timeOfRequest: time.Now(),
	}
	c.Redirect(302, url)
}

type GitLabReceiveTokenReponse struct {
}

// HttpGitLabLoginReceive godoc
// @Summary handles GitLab oauth2 callback
// @Schemes
// @Description where the user's browser is sent by GitLab after completing the oauth2 flow
// @Tags Login
// @Accept json
// @Param code query string true "The code returned by GitLab"
// @Param state query string true "The state returned by GitLab"
// @Produce json
// @Success 200 {object} GitLabReceiveTokenReponse
// @Router /login/gitlab/receive [get]
// func HttpGitLabLoginReceive(c *gin.Context) {
// 	code := c.Query("code")
// 	csrfToken := c.Query("state")
// 	session, ok := gitLabLoginRequests[csrfToken]
// 	if !ok {
// 		c.JSON(400, gin.H{"error": "Invalid csrf token"})
// 		return
// 	}

// 	protocol := "https://"
// 	if c.Request.TLS == nil {
// 		protocol = "http://"
// 	}
// 	url := protocol + c.Request.Host + c.Request.URL.Path
// 	// already contains the /receiveToken ending
// 	access_token, err := gitlab.GetAccessToken(code, url, session.verifier)
// 	if err != nil {
// 		//TODO: handle this better
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	fmt.Println(access_token)
// 	c.JSON(200, gin.H{"access_token": access_token})
// }

func HttpLoginGitlabReceive(c *gin.Context) {
	code := c.Query("code")
	user, err := gitlab.GetUserInfo(code)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
