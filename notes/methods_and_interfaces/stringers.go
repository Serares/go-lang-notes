package methods_and_interfaces

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	to_strings := make([]string, len(i))

	for ind, num := range i {
		to_strings[ind] = fmt.Sprint(num)
	}
	return strings.Join(to_strings, string('.'))
}

func UseStringer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
