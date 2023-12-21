package api

import "github.com/gin-gonic/gin"

// HttpWebhookList godoc
// @Summary List WebHooks that you own
// @Schemes
// @Description Returns a list of WebHooks this user owns.
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld[]
// @Router /api/webhook [get]
func HttpWebhookList(r *gin.Context) {

}

// HttpWebhookCreate godoc
// @Summary [2FA] Create a new WebHook
// @Schemes
// @Description [Two Factor Authentication Required] Creates a new WebHook with the provided subscription details.
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook [post]
func HttpWebhookCreate(r *gin.Context) {

}

// HttpWebhookGet godoc
// @Summary Get's the details of a WebHook
// @Schemes
// @Description Provided either the key, name, or id returns the details of the webhook.
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook/:webhook [get]
func HttpWebhookGet(r *gin.Context) {

}

// HttpWebhookDelete godoc
// @Summary [2FA] Deletes a WebHook
// @Schemes
// @Description [Two Factor Authentication Required] Deletes a WebHook at the given WebHook key or id
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook/:webhook [delete]
func HttpWebhookDelete(r *gin.Context) {

}

// HttpWebhookUpdate godoc
// @Summary [2FA] Update the details of an existing WebHook
// @Schemes
// @Description [Two Factor Authentication Required] Updates the details of an existing WebHook at the given WebHook key or id
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook/:webhook [post]
func HttpWebhookUpdate(r *gin.Context) {

}
