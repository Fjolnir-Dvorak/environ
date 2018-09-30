//+build !windows,!darwin

package pkg

import (
	"os"
	"strings"
)

// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

const (
	envHome              = "HOME"
	envUserConfig        = "XDG_CONFIG_HOME"
	envGlobalConfig      = "XDG_CONFIG_DIRS"
	envUserTemp          = "XDG_CACHE_HOME"
	envUserCache         = "XDG_CACHE_HOME"
	envUserProgramData   = "XDG_DATA_HOME"
	envGlobalProgramData = "XDG_DATA_DIRS"
	envDesktop           = "XDG_DESKTOP_DIR"
	envDocuments         = "XDG_DOCUMENTS_DIR"
	envDownloads         = "XDG_DOWNLOAD_DIR"
	envPictures          = "XDG_PICTURES_DIR"
	envMusic             = "XDG_MUSIC_DIR"
	envVideos            = "XDG_VIDEOS_DIR"
	envSaveGames         = "XDG_DATA_HOME"

	// The environment variables must not be set. Defaults have to be hard coded.
	defUserConfig        = "$HOME/.config"
	defGlobalConfig      = "/etc/xdg"
	defUserCache         = "$HOME/.cache"
	defGlobalCache       = "/var/cache"
	defUserProgramData   = "$HOME/.local/share"
	defGlobalProgramData = "/usr/share"
	defDesktop           = "$HOME/Desktop"
	defDocuments         = "$HOME/Documents"
	defDownloads         = "$HOME/Downloads"
	defPictures          = "$HOME/Pictures"
	defMusic             = "$HOME/Music"
	defVideos            = "$HOME/Videos"
	defSaveGames         = "$HOME/.local/share"

	pathSeparator = ":"
)

func create() Environ {
	home := getenv(envHome)
	userConfig := getenvWithDef(envUserConfig, defUserConfig)
	globalConfig := splitPaths(getenvWithDef(envGlobalConfig, defGlobalConfig))
	userTemp := getenv(envUserTemp)
	userCache := getenvWithDef(envUserCache, defUserCache)
	globalCache := defGlobalCache
	userProgramData := getenvWithDef(envUserProgramData, defUserProgramData)
	globalProgramData := splitPaths(getenvWithDef(envGlobalProgramData, defGlobalProgramData))
	desktop := getenvWithDef(envDesktop, defDesktop)
	documents := getenvWithDef(envDocuments, defDocuments)
	downloads := getenvWithDef(envDownloads, defDownloads)
	pictures := getenvWithDef(envPictures, defPictures)
	music := getenvWithDef(envMusic, defMusic)
	videos := getenvWithDef(envVideos, defVideos)
	saveGames := getenvWithDef(envSaveGames, defSaveGames)

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

func splitPaths(variable string) []string {
	return strings.Split(variable, pathSeparator)
}

func getenv(variable string) string {
	return os.Getenv(variable)
}
func getenvWithDef(variable string, def string) string {
	path := os.Getenv(variable)
	if path == "" {
		return os.ExpandEnv(def)
	} else {
		return path
	}
}
