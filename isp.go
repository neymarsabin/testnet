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
const ISP_INFO_APP_URL = "http://localhost:8000/api/v1/fastness"

func IspDetails() IspInfo {
	var isp IspInfo
	resp, err := http.Get(IP_INFO_URL)

	if err != nil {
		log.Println("Error while requesting ", IP_INFO_URL, " ", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&isp); err != nil {
		log.Println("Error while decoding JSON: ", err)
	}

	return isp
}
