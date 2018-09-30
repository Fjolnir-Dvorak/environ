package main

import (
	"fmt"
	"github.com/Fjolnir-Dvorak/environ/pkg"
)

func main() {
	pkg.CacheResults = false
	pkg.CreateTemp = false
	pkg.CreateUserConfig = false

	env := pkg.New("Fjolnir-Dvorak", "environ")

	fmt.Printf("vendorName: %s\n", env.VendorName())
	fmt.Printf("applicationName: %s\n", env.ApplicationName())
	fmt.Printf("home: %s\n", env.Home())
	fmt.Printf("desktop: %s\n", env.Desktop())
	fmt.Printf("documents: %s\n", env.Documents())
	fmt.Printf("downloads: %s\n", env.Downloads())
	fmt.Printf("pictures: %s\n", env.Pictures())
	fmt.Printf("music: %s\n", env.Music())
	fmt.Printf("videos: %s\n", env.Videos())
	fmt.Printf("saveGames: %s\n", env.SaveGames())
	fmt.Printf("userConfig: %s\n", env.UserConfig())
	fmt.Printf("globalConfig: %s\n", env.GlobalConfig())
	fmt.Printf("userTemp: %s\n", env.UserTemp())
	fmt.Printf("userCache: %s\n", env.UserCache())
	fmt.Printf("globalCache: %s\n", env.GlobalCache())
	fmt.Printf("userProgramData: %s\n", env.UserProgramData())
	fmt.Printf("globalProgramData: %s\n", env.GlobalProgramData())
}
