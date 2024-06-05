// author: neymarsabin
// IspInfo: settings banaune -> API key dine/user based
// yo binary le chai tyo API key use garera every minute ko data chai server lai pathaidinxa
// server ko data lai chai open-data handine
// send the speed value to the server, http post request with Ispinfo's provided API key

package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
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
		chromedp.Sleep(10*time.Second),
		chromedp.InnerHTML("#speed-value", &speedValue),
	)

	if err != nil {
		log.Fatal(err)
	}

	speed, err := strconv.Atoi(speedValue)

	if err != nil {
		log.Fatal("Error while converting: ", speed)
	}

	return speed
}

func main() {
	for range time.Tick(time.Second * 60) {
		var speedUnit Speed
		var speedUnitInt = SpeedDetails()
		isp := IspDetails()
		log.Printf("Name is: %v, Timezone is: %v", isp.Name, isp.Timezone)

		speedUnit.value = speedUnitInt
		speedUnit.timestamp = time.Now().Unix()
		log.Printf("Speed Value at: %v is %d Mbps", time.Unix(speedUnit.timestamp, 0), speedUnit.value)
	}
}
