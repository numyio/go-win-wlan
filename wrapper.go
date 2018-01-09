package winWlan

import (
	"syscall"
	"unsafe"
)

var (
	hapi                         = syscall.NewLazyDLL("Wlanapi.dll")
	hWlanOpenHandle              = hapi.NewProc("WlanOpenHandle")
	hWlanCloseHandle             = hapi.NewProc("WlanCloseHandle")
	hWlanEnumInterfaces          = hapi.NewProc("WlanEnumInterfaces")
	hWlanGetInterfaceCapability  = hapi.NewProc("WlanGetInterfaceCapability")
	hWlanSetInterface            = hapi.NewProc("WlanSetInterface")
	hWlanQueryInterface          = hapi.NewProc("WlanQueryInterface")
	hWlanScan                    = hapi.NewProc("WlanScan")
	hWlanGetAvailableNetworkList = hapi.NewProc("WlanGetAvailableNetworkList")
	hWlanGetNetworkBssList       = hapi.NewProc("WlanGetNetworkBssList")
	hWlanGetProfileList          = hapi.NewProc("WlanGetProfileList")
	hWlanGetProfile              = hapi.NewProc("WlanGetProfile")
	hWlanFreeMemory              = hapi.NewProc("WlanFreeMemory")
)

func wlanOpenHandle(
	dwClientVersion uint32,
	pReserved uintptr,
	pdwNegotiatedVersion *uint32,
	phClientHandle *syscall.Handle) syscall.Errno {
	e, _, _ := hWlanOpenHandle.Call(
		uintptr(dwClientVersion),
		pReserved,
		uintptr(unsafe.Pointer(pdwNegotiatedVersion)),
		uintptr(unsafe.Pointer(phClientHandle)))
	return syscall.Errno(e)
}

func wlanCloseHandle(
	hClientHandle syscall.Handle,
	pReserved uintptr) syscall.Errno {
	e, _, _ := hWlanCloseHandle.Call(uintptr(hClientHandle),
		pReserved)

	return syscall.Errno(e)
}

func wlanFreeMemory(pMemory uintptr) {
	_, _, _ = hWlanFreeMemory.Call(pMemory)
}

func wlanEnumInterfaces(
	hClientHandle syscall.Handle,
	pReserved uintptr,
	ppInterfaceList **WLAN_INTERFACE_INFO_LIST) syscall.Errno {
	e, _, _ := hWlanEnumInterfaces.Call(uintptr(hClientHandle),
		pReserved,
		uintptr(unsafe.Pointer(ppInterfaceList)))

	return syscall.Errno(e)
}

func wlanQueryInterface(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	OpCode WLAN_INTF_OPCODE,
	pReserved uintptr,
	pdwDataSize *uint32,
	ppData **WLAN_CONNECTION_ATTRIBUTES, // Pointer to WLAN_CONNECTION_ATTRIBUTES, WLAN_RADIO_STATE
	pWlanOpcodeValueType uintptr) syscall.Errno {
	e, _, _ := hWlanQueryInterface.Call(
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		uintptr(OpCode),
		pReserved,
		uintptr(unsafe.Pointer(pdwDataSize)),
		uintptr(unsafe.Pointer(ppData)),
		pWlanOpcodeValueType)
	return syscall.Errno(e)
}

func wlanGetInterfaceCapability(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	pReserved uintptr,
	ppCapability **WLAN_INTERFACE_CAPABILITY) syscall.Errno {
	e, _, _ := hWlanGetInterfaceCapability.Call(
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		pReserved,
		uintptr(unsafe.Pointer(ppCapability)))
	return syscall.Errno(e)
}

func wlanSetInterface(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	OpCode WLAN_INTF_OPCODE,
	dwDataSize uint32,
	pData *WLAN_PHY_RADIO_STATE, // Pointer to WLAN_PHY_RADIO_STATE
	pReserved uintptr) syscall.Errno {
	e, _, _ := hWlanSetInterface.Call(
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		uintptr(OpCode),
		uintptr(dwDataSize),
		uintptr(unsafe.Pointer(pData)),
		pReserved)
	return syscall.Errno(e)
}

func wlanScan(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	pDot11Ssid *DOT11_SSID,
	pIeData uintptr,
	pReserved uintptr) syscall.Errno {
	e, _, _ := hWlanScan.Call(
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		uintptr(unsafe.Pointer(pDot11Ssid)),
		pIeData,
		pReserved)
	return syscall.Errno(e)
}

func wlanGetAvailableNetworkList(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	dwFlags uint32,
	pReserved uintptr,
	ppAvailableNetworkList **WLAN_AVAILABLE_NETWORK_LIST) syscall.Errno {
	e, _, _ := hWlanGetAvailableNetworkList.Call(
		uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		uintptr(dwFlags),
		pReserved,
		uintptr(unsafe.Pointer(ppAvailableNetworkList)))

	return syscall.Errno(e)
}

func wlanGetNetworkBssList(hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	pDot11Ssid *DOT11_SSID,
	dot11BssType DOT11_BSS_TYPE,
	bSecurityEnabled int32,
	pReserved uintptr,
	ppWlanBssList **WLAN_BSS_LIST) syscall.Errno {
	e, _, _ := hWlanGetNetworkBssList.Call(uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		uintptr(unsafe.Pointer(pDot11Ssid)),
		uintptr(dot11BssType),
		uintptr(bSecurityEnabled),
		pReserved,
		uintptr(unsafe.Pointer(ppWlanBssList)))

	return syscall.Errno(e)
}

func wlanGetProfileList(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	pReserved uintptr,
	ppProfileList **WLAN_PROFILE_INFO_LIST) syscall.Errno {
	e, _, _ := hWlanGetProfileList.Call(uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		pReserved,
		uintptr(unsafe.Pointer(ppProfileList)))
	return syscall.Errno(e)
} // Pointer to WLAN_PROFILE_INFO_LIST

func wlanGetProfile(
	hClientHandle syscall.Handle,
	pInterfaceGUID *syscall.GUID,
	strProfileName string,
	pReserved uintptr,
	pstrProfileXML *string,
	pdwFlags uint32,
	pdwGrantedAccess *uint32) syscall.Errno {
	e, _, _ := hWlanGetProfile.Call(uintptr(hClientHandle),
		uintptr(unsafe.Pointer(pInterfaceGUID)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(strProfileName))),
		pReserved,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(*pstrProfileXML))),
		uintptr(pdwFlags),
		uintptr(unsafe.Pointer(pdwGrantedAccess)))
	return syscall.Errno(e)
}
