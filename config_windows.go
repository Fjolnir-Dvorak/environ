package environ

import "os"
import "strings"

const (
	CONFIG_LOCAL  = "%APPDATA%"
	CONFIG_SHARED = "%PROGRAMDATA%"
	CACHE         = "%TEMP%"
	DATA_LOCAL    = "%LOCALAPPDATA%"
	DATA_SHARED   = "%PROGRAMDATA%"
)

var (
	configGlobal []string
	configLocal  string
	cache        string
	dataLocal    string
	dataGlobal   []string
)

func init() {
	configLocal = getenv(CONFIG_LOCAL)
	configGlobal = []string{getenv(CONFIG_SHARED)}
	cache = getenv(CACHE)
	dataLocal = getenv(DATA_LOCAL)
	dataGlobal = []string{getenv(DATA_SHARED)}
}

func getenv(variable string) string {
	return os.Getenv(strings.Trim(variable, "%"))
}
