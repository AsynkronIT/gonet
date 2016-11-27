package gonet

import (
	"log"
	"net"
	"strconv"
)

func findFreePort() int {
	l, _ := net.Listen("tcp", ":0")
	defer l.Close()
	s := l.Addr().String()
	_, p := getAddress(s)
	return p
}

func getAddress(addr string) (string, int) {
	h, p, err := net.SplitHostPort(addr)
	if err != nil {
		log.Fatalf("[CLUSTER] Failed parsing address '%v', error %v", addr, err)
	}
	p2, err := strconv.Atoi(p)
	if err != nil {
		log.Fatalf("[CLUSTER] Failed parsing host for address '%v', error %v", addr, err)
	}
	return h, p2
}
