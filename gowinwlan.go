package winWlan

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

type WLANInterface struct {
	guid              syscall.GUID
	handle            *syscall.Handle
	Description       string
	AvailableNetworks []AvailableNetwork
	Profiles          []Profile
}

type Profile struct {
	Name  string
	Flags uint32
	XML   string
}

type AvailableNetwork struct {
	ProfileName   string
	IsConnectable bool
	SignalQuality int
}

type WLAN struct {
	handle     syscall.Handle
	Interfaces []WLANInterface
}

func getVersion() (apiVersion uint32) {
	switch v, _ := syscall.GetVersion(); {
	case v&0xff == 5:
		if (v>>8)&0xFF > 0 {
			apiVersion = 1
		} else {
			// Windows 2000
			fmt.Fprintln(os.Stderr, "Requires Windows XP or later")
			os.Exit(10)
		}
	case v&0xff > 5:
		apiVersion = 2
	default:
		// Earlier than Windows 2000. Its probably crashed before reaching here!
		fmt.Fprintln(os.Stderr, "Requires Windows XP or later")
		os.Exit(10)
	}
	return apiVersion
}

func CreateWLAN() WLAN {
	var localWLAN WLAN
	var negotiatedVersion uint32
	version := getVersion()
	wlanOpenHandle(version, uintptr(0), &negotiatedVersion, &localWLAN.handle)
	return localWLAN
}

func (wlan *WLAN) GetInterfaces() {
	var interfaceList *WLAN_INTERFACE_INFO_LIST
	defer wlanFreeMemory(uintptr(unsafe.Pointer(interfaceList)))

	wlanEnumInterfaces(wlan.handle, 0, &interfaceList)
	wlan.Interfaces = make([]WLANInterface, interfaceList.dwNumberOfItems)
	for index := 0; index < int(interfaceList.dwNumberOfItems); index++ {
		wlan.Interfaces[index].guid = interfaceList.InterfaceInfo[index].InterfaceGuid
		wlan.Interfaces[index].handle = &wlan.handle
		wlan.Interfaces[index].Description = syscall.UTF16ToString(interfaceList.InterfaceInfo[index].strInterfaceDescription[:])
	}
}

func (iface *WLANInterface) GetAvailableNetworks() {
	var aList *WLAN_AVAILABLE_NETWORK_LIST
	wlanGetAvailableNetworkList(
		*iface.handle,
		&iface.guid,
		0,
		uintptr(0),
		&aList)

	iface.AvailableNetworks = make([]AvailableNetwork, aList.dwNumberOfItems)
	for index := 0; index < int(aList.dwNumberOfItems); index++ {
		scannedNetwork := &aList.Network[index]
		localNetwork := &iface.AvailableNetworks[index]

		localNetwork.IsConnectable = scannedNetwork.bNetworkConnectable
		localNetwork.SignalQuality = int(scannedNetwork.wlanSignalQuality)
		localNetwork.ProfileName = syscall.UTF16ToString(scannedNetwork.strProfileName[:])
	}
	wlanFreeMemory(uintptr(unsafe.Pointer(aList)))
	return
}

func (iface *WLANInterface) GetProfiles() {
	var pList *WLAN_PROFILE_INFO_LIST
	defer wlanFreeMemory(uintptr(unsafe.Pointer(pList)))

	wlanGetProfileList(
		*iface.handle,
		&iface.guid,
		uintptr(0),
		&pList)
	iface.Profiles = make([]Profile, pList.dwNumberOfItems)
	for index := 0; index < int(pList.dwNumberOfItems); index++ {
		scannedProfile := &pList.ProfileInfo[index]
		localProfile := &iface.Profiles[index]
		localProfile.Name = syscall.UTF16ToString(scannedProfile.strProfileName[:])
		localProfile.Flags = scannedProfile.dwFlags

	}
}

func (wlan *WLAN) Close() {
	wlanCloseHandle(wlan.handle, uintptr(0))
}
