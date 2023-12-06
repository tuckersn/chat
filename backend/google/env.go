package google

import "os"

func GetGoogleAppId() string {
	return os.Getenv("CR_GOOGLE_APP_ID")
}

func GetGoogleAppSecret() string {
	return os.Getenv("CR_GOOGLE_APP_SECRET")
}

func GetGoogleAuthEnabled() bool {
	return os.Getenv("CR_GOOGLE_AUTH_ENABLED") == "true"
}
