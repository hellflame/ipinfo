package ipinfo

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
)

var regIsHostname *regexp.Regexp

func init() {
	regIsHostname, _ = regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
}

func getIpList(host string) []net.IP {
	addrs, _ := net.LookupIP(host)
	return addrs
}

func IsHostname(host string) bool {
	if regIsHostname.Match([]byte(host)){
		return false
	}
	return true
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

