package main

import (
	"fmt"
	"net"
	"os"

	"github.com/dstotijn/go-bunq"
)

func main() {
	client := bunq.NewClient()
	client.APIKey = os.Getenv("BUNQ_API_KEY")
	client.BaseURL = "https://sandbox.public.api.bunq.com"

	file, err := os.Open("/Users/david/.keys/bunq.key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	if err := client.SetPrivateKey(file); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	fmt.Printf("* Creating installation...\n\n")

	installation, err := client.CreateInstallation()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Created installation: %#v\n", installation)

	fmt.Printf("\n* Creating DeviceServer...\n\n")

	client.Token = installation.Token.Token
	deviceServer, err := client.CreateDeviceServer("Foobar", []net.IP{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	fmt.Printf("Created DeviceServer: %#v\n", deviceServer)
}