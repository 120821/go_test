package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func main() {
	ipAddress := "118.181.151.191" // 假设这是客户端的 IP 地址

	// 发送 HTTP GET 请求到第三方接口获取 IP 地址的位置信息
	resp, err := http.Get("http://ip-api.com/json/" + ipAddress + "?lang=zh-CN")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 解析 JSON 响应
	location := ""
	if resp.StatusCode == http.StatusOK {
		location = string(body)
	}

	fmt.Println("=== Location:", location)

//Location: {"status":"success","country":"China","countryCode":"CN","region":"GS","regionName":"Gansu","city":"Yuzhong Chengguanzhen","zip":"","lat":35.7518,"lon":104.286,"timezone":"Asia/Shanghai","isp":"Chinanet","org":"Chinanet GS","as":"AS4134 CHINANET-BACKBONE","query":"118.181.151.191"}

}
