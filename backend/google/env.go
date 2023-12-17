package google

import (
	"github.com/tuckersn/chatbackend/util"
)

var Config = &util.Config.Google

func AuthEnabled() bool {
	return util.Config.Google.AuthEnabled
}
