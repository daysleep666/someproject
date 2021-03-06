package main

import (
	"os"

	"github.com/hashicorp/mdns"
)

func main() {
	host, _ := os.Hostname()
	info := []string{"My awesome service"}
	service, _ := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})

	defer server.Shutdown()
	select {}
}
