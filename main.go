// author: neymarsabin
// IspInfo: settings banaune -> API key dine/user based
// yo binary le chai tyo API key use garera every minute ko data chai server lai pathaidinxa
// server ko data lai chai open-data handine
// send the speed value to the server, http post request with Ispinfo's provided API key

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

type Speed struct {
	value     int
	timestamp int64
}

func SpeedDetails() int {
	var speedValue string

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://fast.com"),
		chromedp.Sleep(20*time.Second),
		chromedp.InnerHTML("#speed-value", &speedValue),
	)

	if err != nil {
		log.Fatal("Error while running fast.com with chrome: ", err)
	}

	speed, err := strconv.Atoi(speedValue)

	if err != nil {
		log.Fatal("Error while converting: ", speed)
	}

	return speed
}

func saveToDatabase(speed Speed, isp IspInfo) error {
	type postBody map[string]interface{}
	ISP_INFO_APP_URL := os.Getenv("ISP_INFO_APP_URL")
	body := postBody{
		"speed":     speed.value,
		"timezone":  isp.Timezone,
		"ispName":   isp.Name,
		"timestamp": speed.timestamp,
	}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)

	_, err := http.Post(ISP_INFO_APP_URL, "application/json", bytes.NewBuffer(reqBodyBytes.Bytes()))

	if err != nil {
		log.Fatal("Error while saving speed to database", err)
		return err
	}

	return nil
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading envs")
	}

	for range time.Tick(time.Second * 120) {
		fmt.Println("OS ENVS: ", os.Getenv("ISP_INFO_APP_URL"))

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
