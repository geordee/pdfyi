package config

import (
	"os"
	"strings"
)

type appConfig struct {
	ListenPort    int
	AllowInsecure bool
}

// App configuration from environment
var App appConfig

const (
	appAllowInsecure = "APP_ALLOW_INSECURE"
)

const (
	defaultListenPort  = 9090
	defaultUploadLimit = 10
)

// InitializeApp Configuration
func InitializeApp() {
	// Listen Port
	App.ListenPort = defaultListenPort

	// Allow Insecure
	App.AllowInsecure = false
	insecure, ok := os.LookupEnv(appAllowInsecure)
	if ok && (strings.ToLower(insecure) == "true") {
		App.AllowInsecure = true
	}
}
