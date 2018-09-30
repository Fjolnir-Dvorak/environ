package pkg

import (
	"os"
)

const (
	// Sorry. I do know nothing about that system. That was the best I could find.
	envHome     = "HOME"
	envUserTemp = "TMPDATA"

	defUserConfig        = "$HOME//Library/Application Support"
	defGlobalConfig      = "/Library/Application Support"
	defUserCache         = "$HOME/Library/Cache"
	defGlobalCache       = "/Library/Cache"
	defUserProgramData   = "$HOME//Library/Application\\ Support"
	defGlobalProgramData = "/Library/Application\\ Support"
	defDesktop           = "$HOME/Desktop"
	defDocuments         = "$HOME/Documents"
	defDownloads         = "$HOME/Downloads"
	defPictures          = "$HOME/Pictures"
	defMusic             = "$HOME/Music"
	defVideos            = "$HOME/Videos"
	defSaveGames         = "$HOME/Library/Application\\ Support"
)

func create() Environ {
	home := getenv(envHome)
	userConfig := expandEnv(defUserConfig)
	globalConfig := []string{expandEnv(defGlobalConfig)}
	userTemp := getenv(envUserTemp)
	userCache := expandEnv(defUserCache)
	globalCache := defGlobalCache
	userProgramData := expandEnv(defUserProgramData)
	globalProgramData := []string{expandEnv(defGlobalProgramData)}
	desktop := expandEnv(defDesktop)
	documents := expandEnv(defDocuments)
	downloads := expandEnv(defDownloads)
	pictures := expandEnv(defPictures)
	music := expandEnv(defMusic)
	videos := expandEnv(defVideos)
	saveGames := expandEnv(defSaveGames)

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

func getenv(variable string) string {
	return os.Getenv(variable)
}
func expandEnv(variable string) string {
	return os.ExpandEnv(variable)
}
