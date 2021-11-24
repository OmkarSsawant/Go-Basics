package main

import (
	"fmt"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	ipstr := ""

	for i := 0; i < len(ip); i++ {
		ipstr += fmt.Sprint(ip[i])
		if i < len(ip)-1 {
			ipstr += "."
		}
	}
	return ipstr
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
