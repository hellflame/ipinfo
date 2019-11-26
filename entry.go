package main

import (
	"fmt"
	"github.com/hellflame/ipinfo/ipinfo"
	"github.com/spf13/cobra"
)

func runner() *cobra.Command {
	//var target []string
	cmd := &cobra.Command{
		Use: "ipinfo",
		Short: "tool for ip info lookup",
		Version: ipinfo.Version,
		Run: func(cmd *cobra.Command, args []string) {
			for _, target := range args{
				fmt.Println(target)
			}
		},

	}
	//targets := &cobra.Command{
	//	Use: "target",
	//	Short: "ip/host to look up",
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println(args)
	//	},
	//}
	//cmd.AddCommand(targets)
	return cmd
}

func main() {
	fmt.Println(ipinfo.IsHostname("1.1.1.1"))
	//runner().Execute()
	//iplocate.IpInfo(net.IP{127, 0, 0, 1})
	//fmt.Println(ipinfo.IpInfo(nil))
}
