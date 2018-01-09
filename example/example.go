package main

import (
	"fmt"

	"github.com/xilixsys/go-win-wlan"
)

func main() {
	wlan := winWlan.CreateWLAN()
	defer wlan.Close()

	wlan.GetInterfaces()
	for index, iface := range wlan.Interfaces {
		fmt.Printf("interface description: %s index: %d \n", iface.Description, index)
		iface.GetAvailableNetworks()
		for _, net := range iface.AvailableNetworks {
			fmt.Printf("interface description: %b signal: %d profilename: %s\n", net.IsConnectable, net.SignalQuality, net.ProfileName)
		}
		iface.GetProfiles()
		for _, profile := range iface.Profiles {
			fmt.Printf("profile name: %s flags: %d xml %s \n", profile.Name, profile.Flags, profile.XML)
		}
	}

}
