package ipinfo

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

var regIsHostname *regexp.Regexp

func init() {
	regIsHostname, _ = regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
}

// get ip list of given hostname
func IpList(host string) []net.IP {
	var addrs []net.IP
	if isHostname(host) {
		addrs, _ = net.LookupIP(host)
	} else {
		addrs = []net.IP{net.ParseIP(host)}
	}

	return addrs
}

// check if given string is a hostname
func isHostname(host string) bool {
	if regIsHostname.Match([]byte(host)) {
		return false
	}
	return true
}

// fetch API result as map
func IpInfo(ip net.IP) (result map[string]interface{}) {
	var url string
	if ip == nil {
		url = "https://ipinfo.io/json"
	} else {
		url = fmt.Sprintf("https://ipinfo.io/%s/json", ip)
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error while fetching ipinfo.io")
		return
	}
	defer func() { _ = resp.Body.Close() }()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatal("error while parsing result")
		return
	}
	return
}

// get spaces left for others
func calcSpace(word string, maxLength int) int {
	targetL := len(word)
	if maxLength > targetL {
		return maxLength - targetL
	}
	return 0
}

// output beautified result to console
func ParseInfo(info map[string]interface{}) {
	ip, ok := info["ip"]
	if !ok {
		return
	}
	if runtime.GOOS == "windows" {
		fmt.Println(strings.Repeat(" ", 5) + ip.(string))
	} else {
		fmt.Printf("\n\033[01;32m%s\033[00m\n\n",
			strings.Repeat(" ", 5)+ip.(string))
	}
	var keys []string
	for k := range info {
		if k != "ip" && k != "readme" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s%s%s\n", strings.Title(k),
			strings.Repeat(" ", calcSpace(k, MaxOutputWidth)),
			info[k])
	}
}

// combine Ipinfo & ParseInfo together, lazy function
func Display(ip net.IP) {
	ParseInfo(IpInfo(ip))
}
