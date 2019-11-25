package main

import (
	"github.com/hellflame/iplocate-go/iplocate"
)

func main() {
	//iplocate.IpInfo(net.IP{127, 0, 0, 1})
	iplocate.IpInfo(nil)
}
