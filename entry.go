package main

import (
	"fmt"
	"github.com/hellflame/argparse"
)

// Version info of the program
const Version = "0.2.1"

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
		fmt.Println(Version)
		return
	}

	if target == nil || len(*target) == 0 {
		Display(nil)
	} else {
		for _, target := range *target {
			for _, ip := range IpList(target) {
				Display(ip)
			}
		}
	}
}
