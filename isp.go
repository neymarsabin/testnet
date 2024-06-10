package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/chromedp"
)

type IspInfo struct {
	Name     string `json:"isp"`
	Timezone string `json:"timezone"`
}

type Speed struct {
	value     int
	timestamp int64
}

const IP_INFO_URL = "http://ip-api.com/json"

func IspDetails() IspInfo {
	var isp IspInfo
	ISP_INFO_URL := os.Getenv("ISP_INFO_URL")
	resp, err := http.Get(ISP_INFO_URL)

	if err != nil {
		log.Println("Error while requesting ", IP_INFO_URL, " ", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&isp); err != nil {
		log.Println("Error while decoding JSON: ", err)
	}

	return isp
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
