#!/bin/bash

go get github.com/vishvananda/netlink
go build main.go
ip route del default
ip route add default nexthop via 172.17.0.4 nexthop via 172.17.0.5
./main
