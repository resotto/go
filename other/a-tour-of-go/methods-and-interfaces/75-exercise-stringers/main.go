package main

import "fmt"

// IPAddr is type
type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipa IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ipa[0], ipa[1], ipa[2], ipa[3])
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
