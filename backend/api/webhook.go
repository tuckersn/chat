package api

import "github.com/gin-gonic/gin"

// HttpWebhookList godoc
// @Summary list all webhooks controlled by you
// @Schemes
// @Description returns information about a webhook
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld[]
// @Router /api/webhook [get]
func HttpWebhookList(r *gin.Context) {

}

// HttpWebhookCreate godoc
// @Summary create a new webhook, if 2FA is enabled it's required
// @Schemes
// @Description returns information about a webhook
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook [post]
func HttpWebhookCreate(r *gin.Context) {

}

// HttpWebhookGet godoc
// @Summary get's the details of a given webhook
// @Schemes
// @Description returns information about a webhook
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook/:webhookKey [get]
func HttpWebhookGet(r *gin.Context) {

}

// HttpWebhookDelete godoc
// @Summary delete's a webhook, if 2FA is set to strict it's required
// @Schemes
// @Description returns information about a webhook
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook/:webhookKey [delete]
func HttpWebhookDelete(r *gin.Context) {

}

// HttpWebhookUpdate godoc
// @Summary provided details overwrite the existing webhook
// @Schemes
// @Tags Webhook API
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/webhook/:webhookKey [post]
func HttpWebhookUpdate(r *gin.Context) {

}
