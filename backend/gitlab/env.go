package gitlab

import "os"

func Enabled() bool {
	return os.Getenv("CR_GITLAB_ENABLED") == "true"
}

func HostURL() string {
	return os.Getenv("CR_GITLAB_URL")
}

func GetAppId() string {
	return os.Getenv("CR_GITLAB_APP_ID")
}

func GetAppSecret() string {
	return os.Getenv("CR_GITLAB_APP_SECRET")
}
