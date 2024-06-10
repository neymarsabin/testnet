package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type IspInfo struct {
	Name     string `json:"isp"`
	Timezone string `json:"timezone"`
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
