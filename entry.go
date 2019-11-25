package main

import (
	"fmt"
	"github.com/hellflame/ipinfo/ipinfo"
	"github.com/spf13/cobra"
)

func runner() *cobra.Command {
	cmd := &cobra.Command{
		Use: "iplocate",
		Short: "",

	}
	return cmd
}

func main() {
	//iplocate.IpInfo(net.IP{127, 0, 0, 1})
	fmt.Println(ipinfo.IpInfo(nil))
}
