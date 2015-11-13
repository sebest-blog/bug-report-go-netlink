package main

import "github.com/vishvananda/netlink"
import "fmt"
import "net"

func main() {
	routes, err := netlink.RouteList(nil, 0)
	if err != nil {
		fmt.Println("netlink.RouteList: ", err)
		return
	}
	for _, r := range routes {
		// a nil Dst means that this is the default route.
		if r.Dst == nil {
			fmt.Println(r)
			i, err := net.InterfaceByIndex(r.LinkIndex)
			if err != nil {
				fmt.Println("net.InterfaceByIndex", r.LinkIndex, err)
				continue
			}
			fmt.Println(i)
			return
		}
	}
	return
}
