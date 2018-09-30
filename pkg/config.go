// configdir provides access to configuration folder in each platforms.
//
// System wide configuration folders:
//
//   - Windows: %PROGRAMDATA% (C:\ProgramData)
//   - Linux/BSDs: ${XDG_CONFIG_DIRS} (/etc/xdg)
//   - MacOSX: "/Library/Application Support"
//
// User wide configuration folders:
//
//   - Windows: %APPDATA% (C:\Users\<User>\Environ\Roaming)
//   - Linux/BSDs: ${XDG_CONFIG_HOME} (${HOME}/.config)
//   - MacOSX: "${HOME}/Library/Application Support"
//
// User wide cache folders:
//
//   - Windows: %LOCALAPPDATA% (C:\Users\<User>\Environ\Local)
//   - Linux/BSDs: ${XDG_CACHE_HOME} (${HOME}/.cache)
//   - MacOSX: "${HOME}/Library/Caches"
//
// environ returns paths inside the above folders.

package pkg

import (
	"os"
	"path/filepath"
)

type Environ struct {
	vendorName      string
	applicationName string

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
}

type adInfo struct {
	vendorName      string
	applicationName string
}

var (
	CacheResults            = true
	CreateTemp              = true
	CreateUserConfig        = true
	CreateGlobalConfig      = false
	CreateUserCache         = false
	CreateGlobalCache       = false
	CreateUserProgramData   = false
	CreateGlobalProgramData = false
	CreateSavedGames        = false

	cachedEnviron *Environ = nil
)

func New(vendorName, applicationName string) Environ {
	var ad Environ
	if cachedEnviron != nil {
		ad = *cachedEnviron
	} else {
		ad = create()
		if CacheResults {
			cachedEnviron = &ad
		}
	}

	ai := adInfo{
		vendorName:      vendorName,
		applicationName: applicationName,
	}

	return Environ{
		vendorName:      ai.vendorName,
		applicationName: ai.applicationName,

		home:      ad.home,
		desktop:   ad.desktop,
		documents: ad.documents,
		downloads: ad.downloads,
		pictures:  ad.pictures,
		music:     ad.music,
		videos:    ad.videos,
		saveGames: ai.addAppInfo(ad.saveGames, CreateSavedGames, false),

		userConfig:        ai.addAppInfo(ad.userConfig, CreateUserConfig, false),
		globalConfig:      ai.addAppInfoStruct(ad.globalConfig, CreateGlobalConfig, true),
		userTemp:          ai.addAppInfo(ad.userTemp, CreateTemp, false),
		userCache:         ai.addAppInfo(ad.userCache, CreateUserCache, false),
		globalCache:       ai.addAppInfo(ad.globalCache, CreateGlobalCache, true),
		userProgramData:   ai.addAppInfo(ad.userProgramData, CreateUserProgramData, false),
		globalProgramData: ai.addAppInfoStruct(ad.globalProgramData, CreateGlobalProgramData, true),
	}
}

func createDir(path string, create bool, isGlobal bool) {
	if create {
		filemode := 0500
		if isGlobal {
			filemode = 0555
		}
		os.MkdirAll(path, os.FileMode(filemode))
	}
}

func (ad adInfo) addAppInfo(path string, create bool, isGlobal bool) string {
	joined := filepath.Join(path, ad.vendorName, ad.applicationName)
	createDir(joined, create, isGlobal)
	return joined
}

func (ad adInfo) addAppInfoStruct(paths []string, create bool, isGlobal bool) []string {
	length := len(paths)
	toReturn := make([]string, length)
	for i := 0; i < length; i++ {
		toReturn[i] = ad.addAppInfo(paths[i], false, false)
	}
	createDir(toReturn[0], create, isGlobal)
	return toReturn
}

func (ad Environ) Home() string {
	return ad.home
}

func (ad Environ) Desktop() string {
	return ad.desktop
}

func (ad Environ) Documents() string {
	return ad.documents
}

func (ad Environ) Downloads() string {
	return ad.downloads
}

func (ad Environ) Pictures() string {
	return ad.pictures
}

func (ad Environ) Music() string {
	return ad.music
}

func (ad Environ) Videos() string {
	return ad.videos
}

func (ad Environ) SaveGames() string {
	return ad.saveGames
}

func (ad Environ) UserConfig() string {
	return ad.userConfig
}

func (ad Environ) GlobalConfigAll() []string {
	return ad.globalConfig
}

func (ad Environ) GlobalConfig() string {
	return ad.globalConfig[0]
}

func (ad Environ) UserTemp() string {
	return ad.userTemp
}

func (ad Environ) UserCache() string {
	return ad.userCache
}

func (ad Environ) GlobalCache() string {
	return ad.globalCache
}

func (ad Environ) UserProgramData() string {
	return ad.userProgramData
}

func (ad Environ) GlobalProgramDataAll() []string {
	return ad.globalProgramData
}

func (ad Environ) GlobalProgramData() string {
	return ad.globalProgramData[0]
}

func (ad Environ) VendorName() string {
	return ad.vendorName
}
func (ad Environ) ApplicationName() string {
	return ad.applicationName
}
