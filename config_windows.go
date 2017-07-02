package environ

import "os"

const (
	CONFIG_LOCAL  = "APPDATA"
	CONFIG_SHARED = "PROGRAMDATA"
	CACHE         = "TEMP"
	DATA_LOCAL    = "LOCALAPPDATA"
	DATA_SHARED   = "PROGRAMDATA"
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
	configGlobal = []string{os.Getenv(CONFIG_SHARED)}
	cache = os.Getenv(CACHE)
	dataLocal = os.Getenv(DATA_LOCAL)
	dataGlobal = []string{os.Getenv(DATA_SHARED)}
}
