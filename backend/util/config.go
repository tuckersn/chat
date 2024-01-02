package util

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var _storageDirInitialized = false
var StorageDir []string = []string{}

type ConfigTomlFile struct {
	Production      bool   `toml:"production"`
	MainNode        bool   `toml:"main_node"`
	Timezone        string `toml:"timezone"`
	StorageDir      string `toml:"storage_dir"`
	RedirectBaseUrl string `toml:"redirect_base_url"`

	Auth     ConfigAuth     `toml:"Auth"`
	Notes    ConfigNotes    `toml:"Notes"`
	Http     ConfigHTTP     `toml:"HTTP"`
	Database ConfigDatabase `toml:"Database"`
	GitLab   ConfigGitLab   `toml:"GitLab"`
	Google   ConfigGoogle   `toml:"Google"`
	OpenAI   ConfigOpenAI   `toml:"OpenAI"`
}

type ConfigAuth struct {
	TokenSecret        string `toml:"token_secret"`
	TokenIssuer        string `toml:"token_issuer"`
	TokenAudience      string `toml:"token_audience"`
	TokenExpirySeconds int    `toml:"token_expiry_seconds"`
	CookieMaxAge       int    `toml:"cookie_max_age"`
	/// Default: true
	CookieSecure bool `toml:"cookie_secure"`
	/// Default: true
	CookieHttpOnly bool `toml:"cookie_http_only"`
}

type ConfigNotes struct {
	Directory string `toml:"directory"`
}

type ConfigHTTP struct {
	Scheme         string   `toml:"scheme"` // http or https
	Host           string   `toml:"host"`
	Port           int      `toml:"port"`
	CertFile       string   `toml:"cert_file"`
	KeyFile        string   `toml:"key_file"`
	CorsRegexs     []string `toml:"cors_regexs"`
	OpensslCommand string   `toml:"openssl_command"`
}

type ConfigDatabase struct {
	Username string                 `toml:"username"`
	Password string                 `toml:"password"`
	Host     string                 `toml:"host"`
	Port     int                    `toml:"port"`
	Database string                 `toml:"database"`
	Schema   string                 `toml:"schema"`
	SSLMode  string                 `toml:"sslmode"`
	PGVector ConfigDatabasePGVector `toml:"PGVector"`
}

type ConfigDatabasePGVector struct {
	Enabled bool `toml:"enabled"`
}

type ConfigGitLab struct {
	AuthEnabled   bool   `toml:"auth_enabled"`
	AuthAppId     string `toml:"auth_app_id"`
	AuthAppSecret string `toml:"auth_app_secret"`
	InstanceUrl   string `toml:"instance_url"`
}

type ConfigGoogle struct {
	AuthEnabled   bool   `toml:"auth_enabled"`
	AuthAppId     string `toml:"auth_app_id"`
	AuthAppSecret string `toml:"auth_app_secret"`
}

type ConfigOpenAI struct {
	APIKey string `toml:"api_key"`
}

var Config ConfigTomlFile = ConfigTomlFile{}

func GetStorageDir(pathStr string) string {
	if !_storageDirInitialized {
		StorageDir = SmartPathFromStr(Config.StorageDir)
	}
	return path.Join(append(StorageDir, pathStr)...)
}

func CreateStorageDirectoryIfNotExists() {
	if Config.StorageDir == "" {
		panic(errors.New("Config.StorageDir not set"))
	}
	if _, err := os.Stat(GetStorageDir("")); os.IsNotExist(err) {
		err := os.MkdirAll(GetStorageDir(""), 0755)
		if err != nil {
			panic(err)
		}
	}
}

func GetRedirectBaseUrl() string {
	return Config.RedirectBaseUrl
}

func LoadConfigOnStartup() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)
	configFilePath := filepath.Join(exeDir, "Config.toml")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		panic(err)
	}

	err = toml.NewDecoder(configFile).Decode(&Config)
	if err != nil {
		panic(err)
	}

	http := &Config.Http
	{
		if http.CertFile == "" {
			http.CertFile = GetStorageDir("cert.crt")
		}
		if http.KeyFile == "" {
			http.KeyFile = GetStorageDir("cert.key")
		}
	}

	db := &Config.Database
	{
		if db.Username == "" {
			db.Username = "postgres"
		}
		if db.Password == "" {
			db.Password = "postgres"
		}
		if db.Database == "" {
			db.Database = "chatroom"
		}
		if db.Schema == "" {
			db.Schema = "public"
		}
		if db.Host == "" {
			db.Host = "localhost"
		}
		if db.Port == 0 {
			db.Port = 5432
		}
		if db.SSLMode == "" {
			db.SSLMode = "require"
		}
	}

	auth := &Config.Auth
	{
		if auth.CookieHttpOnly == false {
			auth.CookieHttpOnly = true
		}
		if auth.CookieSecure == false {
			auth.CookieSecure = true
		}
	}
}
