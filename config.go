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
//   - Windows: %APPDATA% (C:\Users\<User>\AppData\Roaming)
//   - Linux/BSDs: ${XDG_CONFIG_HOME} (${HOME}/.config)
//   - MacOSX: "${HOME}/Library/Application Support"
//
// User wide cache folders:
//
//   - Windows: %LOCALAPPDATA% (C:\Users\<User>\AppData\Local)
//   - Linux/BSDs: ${XDG_CACHE_HOME} (${HOME}/.cache)
//   - MacOSX: "${HOME}/Library/Caches"
//
// environ returns paths inside the above folders.

package environ

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

type AppData struct {
	vendorName      string
	applicationName string
	configGlobal    []string
	configLocal     string
	cache           string
	dataLocal       string
	dataGlobal      []string
}

func New(vendorName, applicationName string) AppData {
	ad := AppData{
		vendorName:      vendorName,
		applicationName: applicationName,
	}
	ad.configGlobal = ad.addAppInfoStruct(configGlobal)
	ad.configLocal = ad.addAppInfo(configLocal)
	ad.cache = ad.addAppInfo(cache)
	ad.dataLocal = ad.addAppInfo(dataLocal)
	ad.dataGlobal = ad.addAppInfoStruct(dataGlobal)
	return ad
}
func (ad AppData) EnsureExistence(folder Folder) AppData {
	var path string
	switch folder {
	case ConfigGlobal:
		path = ad.configGlobal[0]
		break
	case CONFIG_LOCAL:
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

func (ad AppData) addAppInfo(path string) string {
	return filepath.Join(path, ad.vendorName, ad.applicationName)
}
func (ad AppData) addAppInfoStruct(path []string) []string {
	length := len(path)
	toReturn := make([]string, length)
	for i := 0; i < length; i++ {
		toReturn[i] = ad.addAppInfo(path[i])
	}
}
func (ad AppData) VendorName() string {
	return ad.vendorName
}
func (ad AppData) ApplicationName() string {
	return ad.applicationName
}
func (ad AppData) ConfigGlobal() []string {
	return ad.configGlobal
}
func (ad AppData) Cache() string {
	return ad.cache
}
func (ad AppData) DataLocal() string {
	return ad.dataLocal
}
func (ad AppData) DataGlobal() []string {
	return ad.dataGlobal
}
