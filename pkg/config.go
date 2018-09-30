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

// ConfigDir keeps setting for querying folders.

type Folder int

const (
	ConfigGlobal Folder = iota
	ConfigLocal
	Cache
	DataLocal
	DataGlobal
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
	globalConfig      string
	userTemp          string
	globalTemp        string
	userCache         string
	globalCache       string
	localProgramData  string
	GlobalProgramData string

	configGlobal []string
	configLocal  string
	cache        string
	dataLocal    string
	dataGlobal   []string

	varVendorName      string
	varApplicationName string
	varConfigGlobal    string
	varConfigLocal     string
	varCache           string
	varDataLocal       string
	varDataGlobal      string
}

func New(vendorName,
	applicationName string) Environ {
	ad := Environ{
		vendorName:      vendorName,
		applicationName: applicationName,
	}
	ad.configGlobal = ad.addAppInfoStruct(configGlobal)
	ad.configLocal = ad.addAppInfo(configLocal)
	ad.cache = ad.addAppInfo(cache)
	ad.dataLocal = ad.addAppInfo(dataLocal)
	ad.dataGlobal = ad.addAppInfoStruct(dataGlobal)

	ad.varConfigGlobal = ad.addAppInfo(CONFIG_SHARED)
	ad.varConfigLocal = ad.addAppInfo(CONFIG_LOCAL)
	ad.varCache = ad.addAppInfo(CACHE)
	ad.varDataLocal = ad.addAppInfo(DATA_LOCAL)
	ad.varDataGlobal = ad.addAppInfo(DATA_SHARED)
	return ad
}

func (ad Environ) EnsureExistence(folder Folder) Environ {
	var path string
	switch folder {
	case ConfigGlobal:
		path = ad.configGlobal[0]
		break
	case ConfigLocal:
		path = ad.configLocal
	case Cache:
		path = ad.cache
	case DataGlobal:
		path = ad.dataGlobal[0]
	case DataLocal:
		path = ad.dataLocal
	}
	os.MkdirAll(path, os.ModePerm)
	return ad
}

func (ad Environ) addAppInfo(path string) string {
	return filepath.Join(path, ad.vendorName, ad.applicationName)
}
func (ad Environ) addAppInfoStruct(path []string) []string {
	length := len(path)
	toReturn := make([]string, length)
	for i := 0; i < length; i++ {
		toReturn[i] = ad.addAppInfo(path[i])
	}
	return toReturn
}

func (ad Environ) VendorName() string {
	return ad.vendorName
}
func (ad Environ) ApplicationName() string {
	return ad.applicationName
}
func (ad Environ) ConfigGlobal() []string {
	return ad.configGlobal
}
func (ad Environ) ConfigLocal() string {
	return ad.configLocal
}
func (ad Environ) Cache() string {
	return ad.cache
}
func (ad Environ) DataLocal() string {
	return ad.dataLocal
}
func (ad Environ) DataGlobal() []string {
	return ad.dataGlobal
}

func (ad Environ) VarVendorName() string {
	return ad.varVendorName
}
func (ad Environ) VarApplicationName() string {
	return ad.varApplicationName
}
func (ad Environ) VarConfigGlobal() string {
	return ad.varConfigGlobal
}
func (ad Environ) VarConfigLocal() string {
	return ad.varConfigLocal
}
func (ad Environ) VarCache() string {
	return ad.varCache
}
func (ad Environ) VarDataLocal() string {
	return ad.varDataLocal
}
func (ad Environ) VarDataGlobal() string {
	return ad.varDataGlobal
}
