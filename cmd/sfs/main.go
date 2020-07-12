package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/bykovme/goconfig"
	"github.com/bykovme/sfs/internal/sfsapp"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // For random folders generation
	usrHomePath, err := goconfig.GetUserHomePath()
	var loadedConfig sfsapp.Config
	if err == nil {
		err = goconfig.LoadConfig(usrHomePath+sfsapp.CConfigPath, &loadedConfig)
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}
	} else {
		log.Println(err.Error())
		panic(err)
	}

	app := &sfsapp.App{
		LoadedConfig: loadedConfig,
	}

	log.Println("File service is running on the port: " + app.GetPort())

	err = http.ListenAndServe(":"+app.GetPort(), app)
	log.Fatal(err)
}
