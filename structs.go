package winWlan

import "syscall"

type WLAN_INTERFACE_INFO struct {
	InterfaceGuid           syscall.GUID
	strInterfaceDescription [256]uint16
	isState                 WLAN_INTERFACE_STATE
}

type WLAN_CONNECTION_ATTRIBUTES struct {
	isState                   WLAN_INTERFACE_STATE
	wlanConnectionMode        WLAN_CONNECTION_MODE
	strProfileName            [256]uint16
	wlanAssociationAttributes WLAN_ASSOCIATION_ATTRIBUTES
	wlanSecurityAttributes    WLAN_SECURITY_ATTRIBUTES
}

type WLAN_INTERFACE_CAPABILITY struct {
	interfaceType             WLAN_INTERFACE_TYPE
	bDot11DSupported          bool
	dwMaxDesiredSsidListSize  uint32
	dwMaxDesiredBssidListSize uint32
	dwNumberOfSupportedPhys   uint32
	dot11PhyTypes             [64]DOT11_PHY_TYPE
}

type WLAN_PHY_RADIO_STATE struct {
	dwPhyIndex              uint32
	dot11SoftwareRadioState DOT11_RADIO_STATE
	dot11HardwareRadioState DOT11_RADIO_STATE
}

type WLAN_RADIO_STATE struct {
	dwNumberOfPhys uint32
	PhyRadioState  [MAX_INDEX + 1]WLAN_PHY_RADIO_STATE
}

type WLAN_INTERFACE_INFO_LIST struct {
	dwNumberOfItems uint32
	dwIndex         uint32
	InterfaceInfo   [MAX_INDEX + 1]WLAN_INTERFACE_INFO
}

type WLAN_AVAILABLE_NETWORK struct {
	strProfileName              [256]uint16
	dot11Ssid                   DOT11_SSID
	dot11BssType                DOT11_BSS_TYPE
	uNumberOfBssids             uint32
	bNetworkConnectable         bool
	wlanNotConnectableReason    uint32
	uNumberOfPhyTypes           uint32
	dot11PhyTypes               [8]DOT11_PHY_TYPE
	bMorePhyTypes               bool
	wlanSignalQuality           uint32
	bSecurityEnabled            bool
	dot11DefaultAuthAlgorithm   DOT11_AUTH_ALGORITHM
	dot11DefaultCipherAlgorithm DOT11_CIPHER_ALGORITHM
	dwFlags                     uint32
	dwReserved                  uint32
}

type WLAN_AVAILABLE_NETWORK_LIST struct {
	dwNumberOfItems uint32
	dwIndex         uint32
	Network         [MAX_INDEX + 1]WLAN_AVAILABLE_NETWORK
}

type WLAN_BSS_ENTRY struct {
	dot11Ssid               DOT11_SSID
	uPhyId                  uint32
	dot11Bssid              DOT11_MAC_ADDRESS
	dot11BssType            DOT11_BSS_TYPE
	dot11BssPhyType         DOT11_PHY_TYPE
	lRssi                   int32
	uLinkQuality            uint32
	bInRegDomain            int32
	usBeaconPeriod          uint16
	ullTimestamp            uint64
	ullHostTimestamp        uint64
	usCapabilityInformation uint16
	ulChCenterFrequency     uint32
	wlanRateSet             WLAN_RATE_SET
	ulIeOffset              uint32
	ulIeSize                uint32
}

type WLAN_BSS_LIST struct {
	dwTotalSize     uint32
	dwNumberOfItems uint32
	wlanBssEntries  [MAX_INDEX + 1]WLAN_BSS_ENTRY
}

type WLAN_PROFILE_INFO struct {
	strProfileName [256]uint16
	dwFlags        uint32
}

type WLAN_PROFILE_INFO_LIST struct {
	dwNumberOfItems uint32
	dwIndex         uint32
	ProfileInfo     [MAX_INDEX + 1]WLAN_PROFILE_INFO
}

type DOT11_SSID struct {
	uSSIDLength uint32
	ucSSID      [32]byte
}

type DOT11_MAC_ADDRESS [6]byte

type WLAN_ASSOCIATION_ATTRIBUTES struct {
	dot11Ssid         DOT11_SSID
	dot11BssType      DOT11_BSS_TYPE
	dot11Bssid        DOT11_MAC_ADDRESS
	dot11PhyType      DOT11_PHY_TYPE
	uDot11PhyIndex    uint32
	wlanSignalQuality uint32
	ulRxRate          uint32
	ulTxRate          uint32
}

type WLAN_SECURITY_ATTRIBUTES struct {
	bSecurityEnabled     bool
	bOneXEnabled         bool
	dot11AuthAlgorithm   DOT11_AUTH_ALGORITHM
	dot11CipherAlgorithm DOT11_CIPHER_ALGORITHM
}

type WLAN_RATE_SET struct {
	uRateSetLength uint32
	usRateSet      [126]uint16
}

type WLAN_CONNECTION_PARAMETERS struct {
	wlanConnectionMode WLAN_CONNECTION_MODE
	strProfile         string
	pDot11Ssid         uintptr
	pDesiredBssidList  uintptr
	dot11BssType       DOT11_BSS_TYPE
	dwFlags            uint32
}

type DOT11_BSSID_LIST struct {
	Header             NDIS_OBJECT_HEADER
	uNumOfEntries      uint32
	uTotalNumOfEntries uint32
	BSSIDs             DOT11_MAC_ADDRESS
}

type NDIS_OBJECT_HEADER struct {
	Type     byte
	Revision byte
	Size     uint16
}

type WLAN_NOTIFICATION_DATA struct {
	NotificationSource uint32
	NotificationCode   uint32
	InterfaceGuid      syscall.GUID
	dwDataSize         uint32
	pData              uintptr
}

type WLAN_CONNECTION_NOTIFICATION_DATA struct {
	wlanConnectionMode WLAN_CONNECTION_MODE
	strProfileName     [WLAN_MAX_NAME_LENGTH]uint16
	dot11Ssid          DOT11_SSID
	dot11BssType       DOT11_BSS_TYPE
	bSecurityEnabled   bool
	wlanReasonCode     uint32
	dwFlags            uint32
	strProfileXml      [1]uint16
}
