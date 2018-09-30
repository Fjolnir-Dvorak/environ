package pkg

import (
	"golang.org/x/sys/windows/registry"
	"os"
)

const (
	ENV_HOME                = "UserProfile"
	ENV_USER_CONFIG         = "AppData"
	ENV_GLOBAL_CONFIG       = "AllUsersProfile"
	ENV_USER_TEMP           = "Temp"
	ENV_GLOBAL_TEMP         = "Temp"
	ENV_USER_CACHE          = "LocalAppData"
	ENV_GLOBAL_CACHE        = "ProgramData"
	ENV_USER_PROGRAM_DATA   = "LocalAppData"
	ENV_GLOBAL_PROGRAM_DATA = "ProgramFiles"

	REGISTRY_KEY_PATH    = `Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`
	REGISTRY_DESKTOP     = `Desktop`
	REGISTRY_DOCUMENTS   = `Personal`
	REGISTRY_DOWNLOADS   = `{374DE290-123F-4565-9164-39C4925E467B}`
	REGISTRY_PICTURES    = `My Pictures`
	REGISTRY_MUSIC       = `My Music`
	REGISTRY_VIDEOS      = `My Videos`
	REGISTRY_SAVED_GAMES = `{4C5C32FF-BB9D-43B0-B5B4-2D72E54EAAA4}`
)

var (
	home      string
	desktop   string
	documents string
	downloads string
	pictures  string
	music     string
	videos    string
	saveGames string

	userConfig        string
	globalConfig      []string
	userTemp          string
	globalTemp        string
	userCache         string
	globalCache       string
	userProgramData   string
	globalProgramData []string

	hasUserConfig        = true
	hasGlobalConfig      = true
	hasUserTmp           = true
	hasGlobalTmp         = false
	hasUserCache         = true
	hasGlobalCache       = true
	hasUserProgramData   = true
	hasGlobalProgramData = true
)

func init() {
	home = getenv(ENV_HOME)
	userConfig = getenv(ENV_USER_CONFIG)
	globalConfig = []string{getenv(ENV_GLOBAL_CONFIG)}
	userTemp = getenv(ENV_USER_TEMP)
	globalTemp = getenv(ENV_GLOBAL_TEMP)
	userCache = getenv(ENV_USER_CACHE)
	globalCache = getenv(ENV_GLOBAL_CACHE)
	userProgramData = getenv(ENV_USER_PROGRAM_DATA)
	globalProgramData = []string{getenv(ENV_GLOBAL_PROGRAM_DATA)}

	key, err := registry.OpenKey(registry.CURRENT_USER,
		REGISTRY_KEY_PATH,
		registry.QUERY_VALUE)
	if err != nil {
		// TODO handle the error
	}
	defer key.Close()

	desktop = getRegistry(key, REGISTRY_DESKTOP, "")
	documents = getRegistry(key, REGISTRY_DOCUMENTS, "")
	downloads = getRegistry(key, REGISTRY_DOWNLOADS, "")
	pictures = getRegistry(key, REGISTRY_PICTURES, "")
	music = getRegistry(key, REGISTRY_MUSIC, "")
	videos = getRegistry(key, REGISTRY_VIDEOS, "")
	saveGames = getRegistry(key, REGISTRY_SAVED_GAMES, "")
}

func getRegistry(key registry.Key, name string, def string) string {
	value, _, err := key.GetStringValue(name)
	if err != nil {
		value = def
		// TODO handle the error
	}
	return value
}

func getenv(variable string) string {
	return os.Getenv(variable)
}
