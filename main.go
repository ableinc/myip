package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	// Define command line args
	ipVersion := flag.Int("v", 4, "Which IP version to display")
	// Parse command line arguments
	flag.Parse()
	// Validate IP version
	if *ipVersion != 4 && *ipVersion != 6 {
		fmt.Fprintf(os.Stderr, "[error] IP version must be 4 or 6. You provided: %d\n", *ipVersion)
		flag.Usage()
		os.Exit(1)
	}
	// Execute operation
	getIPAddress(ipVersion)
	// Exit successfully
	os.Exit(0)
}

func getIPAddress(ipVersion *int) {
	var url string = "https://ifconfig.me" // HTTPS
	if *ipVersion == 6 {
		url = "http://ifconfig.me" // HTTP
	}
	cmd := exec.Command("curl", "-4", url)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to fetch IPv%d", ipVersion)
		os.Exit(2)
	}
	fmt.Println(string(output))
}
