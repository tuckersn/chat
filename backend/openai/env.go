package openai

import "github.com/tuckersn/chatbackend/util"

func APIKey() string {
	return util.Config.OpenAI.APIKey
}
