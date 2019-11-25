package ipinfo

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

func getIpList(host string) []net.IP {
	addrs, _ := net.LookupIP(host)
	return addrs
}

func IpInfo(ip net.IP)(result map[string]interface{}) {
	var url string
	if ip == nil {
		url = "https://ipinfo.io/json"
	}else {
		url = fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error while fetching ipinfo.io")
		return
	}
	defer func() {_ = resp.Body.Close()}()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal("error while parsing result")
		return
	}
	return
}

