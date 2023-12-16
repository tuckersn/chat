package util

import (
	"errors"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"

	"github.com/pelletier/go-toml/v2"
)

var _storageDirInitialized = false
var StorageDir []string = []string{}
var SlashesRegex *regexp.Regexp = regexp.MustCompile(`[/\\]+`)

type ConfigTomlFile struct {
	Production      bool   `toml:"production"`
	MainNode        bool   `toml:"main_node"`
	StorageDir      string `toml:"storage_dir"`
	RedirectBaseUrl string `toml:"redirect_base_url"`
	ApiBaseUrl      string `toml:"api_base_url"`

	Auth     ConfigAuth     `toml:"Auth"`
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
}

type ConfigHTTP struct {
	Port           int      `toml:"port"`
	CertFile       string   `toml:"cert_file"`
	KeyFile        string   `toml:"key_file"`
	CorsRegexs     []string `toml:"cors_regexs"`
	OpensslCommand string   `toml:"openssl_command"`
}

type ConfigDatabase struct {
	PGVector ConfigDatabasePGVector `toml:"PGVector"`
}

type ConfigDatabasePGVector struct {
}

type ConfigGitLab struct {
}

type ConfigGoogle struct {
}

type ConfigOpenAI struct {
}

var Config ConfigTomlFile = ConfigTomlFile{}

func GetStorageDir(pathStr string) string {
	if !_storageDirInitialized {
		StorageDir = SlashesRegex.Split(Config.StorageDir, -1)
		if StorageDir[0] == "~" {
			usr, err := user.Current()
			if err != nil {
				panic(errors.New("Error getting current user"))
			}
			StorageDir = append(SlashesRegex.Split(usr.HomeDir, -1), StorageDir[1:]...)
		}
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

func IsMainNode() bool {
	return os.Getenv("MAIN_NODE") == "true"
}

func GetRedirectBaseUrl() string {
	return os.Getenv("CR_REDIRECT_URL")
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
}
