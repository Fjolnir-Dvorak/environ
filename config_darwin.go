package environ

import "os"

const (
	// Sorry. I do know nothing about that system. That was the best I could find.
	CONFIG_LOCAL_DEF  = "${home}/Library/Application Support"
	CONFIG_SHARED_DEF = "/Library/Application Support"
	CACHE_DEF         = "${home}/Library/Caches"
	DATA_LOCAL_DEF    = CONFIG_LOCAL_DEF
	DATA_SHARED_DEF   = CONFIG_SHARED_DEF

	CONFIG_LOCAL  = CONFIG_LOCAL_DEF
	CONFIG_SHARED = CONFIG_SHARED_DEF
	CACHE         = CACHE_DEF
	DATA_LOCAL    = DATA_LOCAL_DEF
	DATA_SHARED   = DATA_SHARED_DEF
)

var (
	configGlobal []string
	configLocal  string
	cache        string
	dataLocal    string
	dataGlobal   []string
)

func init() {
	configLocal = os.ExpandEnv(CONFIG_LOCAL_DEF)
	configGlobal = []string{os.ExpandEnv(CONFIG_SHARED_DEF)}
	cache = os.ExpandEnv(CACHE_DEF)
	dataLocal = os.ExpandEnv(DATA_LOCAL)
	dataGlobal = []string{os.ExpandEnv(DATA_SHARED_DEF)}
}
