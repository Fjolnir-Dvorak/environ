package pkg

import (
	"golang.org/x/sys/windows/registry"
	"os"
)

const (
	EnvHome              = "UserProfile"
	EnvUserConfig        = "AppData"
	EnvGlobalConfig      = "AllUsersProfile"
	EnvUserTemp          = "Temp"
	EnvUserCache         = "LocalAppData"
	EnvGlobalCache       = "ProgramData"
	EnvUserProgramData   = "LocalAppData"
	EnvGlobalProgramData = "ProgramFiles"

	RegistryKeyPath    = `Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`
	RegistryDesktop    = `Desktop`
	RegistryDocuments  = `Personal`
	RegistryDownloads  = `{374DE290-123F-4565-9164-39C4925E467B}`
	RegistryPictures   = `My Pictures`
	RegistryMusic      = `My Music`
	RegistryVideos     = `My Videos`
	RegistrySavedGames = `{4C5C32FF-BB9D-43B0-B5B4-2D72E54EAAA4}`
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
	userCache         string
	globalCache       string
	userProgramData   string
	globalProgramData []string
)

func create() Environ {
	home = getenv(EnvHome)
	userConfig = getenv(EnvUserConfig)
	globalConfig = []string{getenv(EnvGlobalConfig)}
	userTemp = getenv(EnvUserTemp)
	userCache = getenv(EnvUserCache)
	globalCache = getenv(EnvGlobalCache)
	userProgramData = getenv(EnvUserProgramData)
	globalProgramData = []string{getenv(EnvGlobalProgramData)}

	key, err := registry.OpenKey(registry.CURRENT_USER,
		RegistryKeyPath,
		registry.QUERY_VALUE)
	if err != nil {
		// TODO handle the error
	}
	defer key.Close()

	desktop = getRegistry(key, RegistryDesktop, "")
	documents = getRegistry(key, RegistryDocuments, "")
	downloads = getRegistry(key, RegistryDownloads, "")
	pictures = getRegistry(key, RegistryPictures, "")
	music = getRegistry(key, RegistryMusic, "")
	videos = getRegistry(key, RegistryVideos, "")
	saveGames = getRegistry(key, RegistrySavedGames, "")

	return Environ{
		home:      home,
		desktop:   desktop,
		documents: documents,
		downloads: downloads,
		pictures:  pictures,
		music:     music,
		videos:    videos,
		saveGames: saveGames,

		userConfig:        userConfig,
		globalConfig:      globalConfig,
		userTemp:          userTemp,
		userCache:         userCache,
		globalCache:       globalCache,
		userProgramData:   userProgramData,
		globalProgramData: globalProgramData,
	}
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
