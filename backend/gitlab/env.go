package gitlab

import (
	"github.com/tuckersn/chatbackend/util"
)

func Enabled() bool {
	return util.Config.GitLab.AuthEnabled
}

func HostURL() string {
	return util.Config.GitLab.InstanceUrl
}

func GetAppId() string {
	return util.Config.GitLab.AuthAppId
}

func GetAppSecret() string {
	return util.Config.GitLab.AuthAppSecret
}
