// +build !windows,!darwin

package environ

import (
	"os"
	"path/filepath"
	"strings"
)

// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

const (
	CONFIG_LOCAL  = "XDG_CONFIG_HOME"
	CONFIG_SHARED = "XDG_CONFIG_DIRS"
	CACHE         = "XDG_CACHE_HOME"
	DATA_LOCAL    = "XDG_DATA_HOME"
	DATA_SHARED   = "XDG_DATA_DIRS"

	// The environment variables must not be set. Defaults have to be hard coded.
	CONFIG_LOCAL_DEF  = "${HOME}/.config/"
	CONFIG_SHARED_DEF = "/etc/xdg/"
	CACHE_DEF         = "${home}/.cache/"
	DATA_LOCAL_DEF    = "${home]/.local/share/"
	DATA_SHARED_DEF   = "/usr/local/share/:/usr/share/"

	PATH_SEPARATOR = ":"
)

var (
	configGlobal  []string
	configLocal   string
	cache         string
	dataLocal     string
	dataGlobal    []string
)

func init() {
	configLocal = os.Getenv(CONFIG_LOCAL)
	if os.Getenv(configLocal) == "" {
		configLocal = os.ExpandEnv(CONFIG_LOCAL_DEF)
	}

	configGlobal = os.Getenv(CONFIG_SHARED)
	if configGlobal == "" {
		configGlobal = CONFIG_SHARED_DEF
	}
	configGlobal = strings.Split(configGlobal, PATH_SEPARATOR)

	cache = os.Getenv(CACHE)
	if os.Getenv(configLocal) == "" {
		cache = os.ExpandEnv(CACHE_DEF)
	}

	dataLocal = os.Getenv(DATA_LOCAL)
	if os.Getenv(configLocal) == "" {
		dataLocal = os.ExpandEnv(DATA_LOCAL_DEF)
	}

	dataGlobal = os.Getenv(DATA_SHARED)
	if dataGlobal == "" {
		dataGlobal = DATA_SHARED_DEF
	}
	dataGlobal = strings.Split(dataGlobal, PATH_SEPARATOR)
}
