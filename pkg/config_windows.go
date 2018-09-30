package pkg

import (
	"golang.org/x/sys/windows/registry"
	"os"
)

const (
	envHome              = "UserProfile"
	envUserConfig        = "AppData"
	envGlobalConfig      = "AllUsersProfile"
	envUserTemp          = "Temp"
	envUserCache         = "LocalAppData"
	envGlobalCache       = "ProgramData"
	envUserProgramData   = "LocalAppData"
	envGlobalProgramData = "ProgramFiles"

	registryKeyPath    = `Software\Microsoft\Windows\CurrentVersion\Explorer\Shell Folders`
	registryDesktop    = `Desktop`
	registryDocuments  = `Personal`
	registryDownloads  = `{374DE290-123F-4565-9164-39C4925E467B}`
	registryPictures   = `My Pictures`
	registryMusic      = `My Music`
	registryVideos     = `My Videos`
	registrySavedGames = `{4C5C32FF-BB9D-43B0-B5B4-2D72E54EAAA4}`
)

func create() Environ {
	home := getenv(envHome)
	userConfig := getenv(envUserConfig)
	globalConfig := []string{getenv(envGlobalConfig)}
	userTemp := getenv(envUserTemp)
	userCache := getenv(envUserCache)
	globalCache := getenv(envGlobalCache)
	userProgramData := getenv(envUserProgramData)
	globalProgramData := []string{getenv(envGlobalProgramData)}

	key, err := registry.OpenKey(registry.CURRENT_USER,
		registryKeyPath,
		registry.QUERY_VALUE)
	if err != nil {
		// TODO handle the error
	}
	defer key.Close()

	desktop := getRegistry(key, registryDesktop, "")
	documents := getRegistry(key, registryDocuments, "")
	downloads := getRegistry(key, registryDownloads, "")
	pictures := getRegistry(key, registryPictures, "")
	music := getRegistry(key, registryMusic, "")
	videos := getRegistry(key, registryVideos, "")
	saveGames := getRegistry(key, registrySavedGames, "")

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
