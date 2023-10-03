package main

import (
	"fmt"
	"os"

	"./Examples"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Operation not found. Optional operations are rule/discard/add_host/show_hosts/dup_ip/auto_publish")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "discard":
		Examples.DiscardSessions()
	case "rule":
		Examples.AddAccessRule()
	case "add_host":
		Examples.AddHost()
	case "show_hosts":
		Examples.ShowHosts()
	case "dup_ip":
		Examples.DupIp()
	case "auto_publish":
		Examples.AutoPublish()
	default:
		fmt.Println("Operation not supported. Optional operations are rule/discard/add_host/show_hosts/dup_ip/auto_publish")
	}
}
