package main

import (
	"fmt"
	"github.com/hellflame/argparse"
	"github.com/hellflame/ipinfo/ipinfo"
)

func main() {
	parser := argparse.NewParser("ipinfo", "tool for ip info lookup",
		&argparse.ParserConfig{DisableDefaultShowHelp: true})
	showVersion := parser.Flag("v", "version", &argparse.Option{Help: "show version info"})
	target := parser.Strings("ip", "host", &argparse.Option{Positional: true, Help: "target ip/host"})
	if e := parser.Parse(nil); e != nil {
		fmt.Println(e.Error())
		return
	}
	if *showVersion {
		fmt.Println(ipinfo.Version)
		return
	}

	if target == nil || len(*target) == 0 {
		ipinfo.Display(nil)
	} else {
		for _, target := range *target {
			for _, ip := range ipinfo.IpList(target) {
				ipinfo.Display(ip)
			}
		}
	}
}
