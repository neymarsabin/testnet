package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IspInfo struct {
	Name     string `json:"isp"`
	Timezone string `json:"timezone"`
}

const IP_INFO_URL = "http://ip-api.com/json"

func IspDetails() *IspInfo {
	var isp IspInfo
	resp, err := http.Get(IP_INFO_URL)

	if err != nil {
		log.Fatal("Error while requesting ", IP_INFO_URL, " ", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&isp); err != nil {
		log.Fatal("Error while decoding JSON: ", err)
	}

	return &isp
}