package main

import (
	"github.com/hellflame/ipinfo/ipinfo"
	"github.com/spf13/cobra"
)

func runner() *cobra.Command {
	//var target []string
	cmd := &cobra.Command{
		Use:     "ipinfo",
		Short:   "tool for ip info lookup",
		Version: ipinfo.Version,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				ipinfo.Display(nil)
				return
			}
			for _, target := range args {
				for _, ip := range ipinfo.IpList(target) {
					ipinfo.Display(ip)
				}
			}
		},
	}
	return cmd
}

func main() {
	_ = runner().Execute()
	//fmt.Println(ipinfo.IpList("1283.12.31.1")[0] == nil)
	//fmt.Println(ipinfo.IsHostname("1.1.1.1"))
	//runner().Execute()
	//iplocate.IpInfo(net.IP{127, 0, 0, 1})
	//fmt.Println(ipinfo.IpInfo(nil))
}
