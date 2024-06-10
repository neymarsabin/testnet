// author: neymarsabin
// IspInfo: settings banaune -> API key dine/user based
// yo binary le chai tyo API key use garera every minute ko data chai server lai pathaidinxa
// server ko data lai chai open-data handine
// send the speed value to the server, http post request with Ispinfo's provided API key

package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get current directory")
	}

	err = godotenv.Load(filepath.Join(pwd, ".env"))

	if err != nil {
		log.Fatal("Error while loading envs")
	}

	for range time.Tick(time.Second * 180) {
		var speedUnit Speed
		var speedUnitInt = SpeedDetails()
		isp := IspDetails()

		speedUnit.value = speedUnitInt
		speedUnit.timestamp = time.Now().Unix()

		// save to Database
		err := saveToDatabase(speedUnit, isp)
		if err != nil {
			log.Fatal("Error while saving to database: ", err)
		}
	}
}
