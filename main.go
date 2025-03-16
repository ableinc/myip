package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func getIPv4() {
	var url string = "https://ifconfig.me" // HTTPS
	cmd := exec.Command("curl", "-4", url)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to fetch IPv4")
		os.Exit(2)
	}
	fmt.Println(string(output))
}

func getIPv6() {
	var url string = "http://ifconfig.me" // HTTP
	cmd := exec.Command("curl", "-4", url)
	output, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to fetch IPv6")
		os.Exit(2)
	}
	fmt.Println(string(output))
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: myip v4 (or v6)\n")
		os.Exit(1)
	}
	// Parse command line arguments
	flag.Parse()
	// Work with arguments
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
	}
	ipVersion := flag.Arg(0)
	switch ipVersion {
	case "v4":
		getIPv4()
		break
	case "v6":
		getIPv6()
		break
	default:
		flag.Usage()
	}
}
