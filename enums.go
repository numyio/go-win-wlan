package winWlan

type WLAN_INTERFACE_STATE uint32

const (
	wlan_interface_state_not_ready             WLAN_INTERFACE_STATE = iota
	wlan_interface_state_connected                                  = iota
	wlan_interface_state_ad_hoc_network_formed                      = iota
	wlan_interface_state_disconnecting                              = iota
	wlan_interface_state_disconnected                               = iota
	wlan_interface_state_associating                                = iota
	wlan_interface_state_discovering                                = iota
	wlan_interface_state_authenticating                             = iota
)

type WLAN_CONNECTION_MODE uint32

const (
	wlan_connection_mode_profile            WLAN_CONNECTION_MODE = iota
	wlan_connection_mode_temporary_profile                       = iota
	wlan_connection_mode_discovery_secure                        = iota
	wlan_connection_mode_discovery_unsecure                      = iota
	wlan_connection_mode_auto                                    = iota
	wlan_connection_mode_invalid                                 = iota
)

type WLAN_INTERFACE_TYPE uint32

const (
	wlan_interface_type_emulated_802_11 WLAN_INTERFACE_TYPE = iota
	wlan_interface_type_native_802_11                       = iota
	wlan_interface_type_invalid                             = iota
)

type DOT11_RADIO_STATE uint32

const (
	dot11_radio_state_unknown DOT11_RADIO_STATE = iota
	dot11_radio_state_on                        = iota
	dot11_radio_state_off                       = iota
)

type DOT11_BSS_TYPE uint32

const (
	dot11_BSS_type_infrastructure DOT11_BSS_TYPE = 1
	dot11_BSS_type_independent                   = 2
	dot11_BSS_type_any                           = 3
)

type DOT11_PHY_TYPE uint32

const (
	dot11_phy_type_unknown    DOT11_PHY_TYPE = 0
	dot11_phy_type_any                       = 0
	dot11_phy_type_fhss                      = 1
	dot11_phy_type_dsss                      = 2
	dot11_phy_type_irbaseband                = 3
	dot11_phy_type_ofdm                      = 4
	dot11_phy_type_hrdsss                    = 5
	dot11_phy_type_erp                       = 6
	dot11_phy_type_ht                        = 7
	dot11_phy_type_vht                       = 8
	dot11_phy_type_IHV_start                 = 0x80000000
	dot11_phy_type_IHV_end                   = 0xffffffff
)

type DOT11_AUTH_ALGORITHM uint32

const (
	DOT11_AUTH_ALGO_80211_OPEN       DOT11_AUTH_ALGORITHM = 1
	DOT11_AUTH_ALGO_80211_SHARED_KEY                      = 2
	DOT11_AUTH_ALGO_WPA                                   = 3
	DOT11_AUTH_ALGO_WPA_PSK                               = 4
	DOT11_AUTH_ALGO_WPA_NONE                              = 5
	DOT11_AUTH_ALGO_RSNA                                  = 6
	DOT11_AUTH_ALGO_RSNA_PSK                              = 7
	DOT11_AUTH_ALGO_IHV_START                             = 0x80000000
	DOT11_AUTH_ALGO_IHV_END                               = 0xffffffff
)

type DOT11_CIPHER_ALGORITHM uint32

const (
	DOT11_CIPHER_ALGO_NONE          DOT11_CIPHER_ALGORITHM = 0x00
	DOT11_CIPHER_ALGO_WEP40                                = 0x01
	DOT11_CIPHER_ALGO_TKIP                                 = 0x02
	DOT11_CIPHER_ALGO_CCMP                                 = 0x04
	DOT11_CIPHER_ALGO_WEP104                               = 0x05
	DOT11_CIPHER_ALGO_WPA_USE_GROUP                        = 0x100
	DOT11_CIPHER_ALGO_RSN_USE_GROUP                        = 0x100
	DOT11_CIPHER_ALGO_WEP                                  = 0x101
	DOT11_CIPHER_ALGO_IHV_START                            = 0x80000000
	DOT11_CIPHER_ALGO_IHV_END                              = 0xffffffff
)

type WLAN_INTF_OPCODE uint32

const (
	wlan_intf_opcode_autoconf_start                             WLAN_INTF_OPCODE = 0x000000000
	wlan_intf_opcode_autoconf_enabled                                            = wlan_intf_opcode_autoconf_start + 1
	wlan_intf_opcode_background_scan_enabled                                     = wlan_intf_opcode_autoconf_start + 2
	wlan_intf_opcode_media_streaming_mode                                        = wlan_intf_opcode_autoconf_start + 3
	wlan_intf_opcode_radio_state                                                 = wlan_intf_opcode_autoconf_start + 4
	wlan_intf_opcode_bss_type                                                    = wlan_intf_opcode_autoconf_start + 5
	wlan_intf_opcode_interface_state                                             = wlan_intf_opcode_autoconf_start + 6
	wlan_intf_opcode_current_connection                                          = wlan_intf_opcode_autoconf_start + 7
	wlan_intf_opcode_channel_number                                              = wlan_intf_opcode_autoconf_start + 8
	wlan_intf_opcode_supported_infrastructure_auth_cipher_pairs                  = wlan_intf_opcode_autoconf_start + 9
	wlan_intf_opcode_supported_adhoc_auth_cipher_pairs                           = wlan_intf_opcode_autoconf_start + 10
	wlan_intf_opcode_supported_country_or_region_string_list                     = wlan_intf_opcode_autoconf_start + 11
	wlan_intf_opcode_current_operation_mode                                      = wlan_intf_opcode_autoconf_start + 12
	wlan_intf_opcode_supported_safe_mode                                         = wlan_intf_opcode_autoconf_start + 13
	wlan_intf_opcode_certified_safe_mode                                         = wlan_intf_opcode_autoconf_start + 14
	wlan_intf_opcode_hosted_network_capable                                      = wlan_intf_opcode_autoconf_start + 15
	wlan_intf_opcode_management_frame_protection_capable                         = wlan_intf_opcode_autoconf_start + 16
	wlan_intf_opcode_autoconf_end                                                = 0x0fffffff
	wlan_intf_opcode_msm_start                                                   = 0x10000100
	wlan_intf_opcode_statistics                                                  = wlan_intf_opcode_msm_start + 1
	wlan_intf_opcode_rssi                                                        = wlan_intf_opcode_msm_start + 2
	wlan_intf_opcode_msm_end                                                     = 0x1fffffff
	wlan_intf_opcode_security_start                                              = 0x20010000
	wlan_intf_opcode_security_end                                                = 0x2fffffff
	wlan_intf_opcode_ihv_start                                                   = 0x30000000
	wlan_intf_opcode_ihv_end                                                     = 0x3fffffff
)

type WLAN_NOTIFICATION_ACM uint32

const (
	wlan_notification_acm_start                      WLAN_NOTIFICATION_ACM = iota
	wlan_notification_acm_autoconf_enabled                                 = iota
	wlan_notification_acm_autoconf_disabled                                = iota
	wlan_notification_acm_background_scan_enabled                          = iota
	wlan_notification_acm_background_scan_disabled                         = iota
	wlan_notification_acm_bss_type_change                                  = iota
	wlan_notification_acm_power_setting_change                             = iota
	wlan_notification_acm_scan_complete                                    = iota
	wlan_notification_acm_scan_fail                                        = iota
	wlan_notification_acm_connection_start                                 = iota
	wlan_notification_acm_connection_complete                              = iota
	wlan_notification_acm_connection_attempt_fail                          = iota
	wlan_notification_acm_filter_list_change                               = iota
	wlan_notification_acm_interface_arrival                                = iota
	wlan_notification_acm_interface_removal                                = iota
	wlan_notification_acm_profile_change                                   = iota
	wlan_notification_acm_profile_name_change                              = iota
	wlan_notification_acm_profiles_exhausted                               = iota
	wlan_notification_acm_network_not_available                            = iota
	wlan_notification_acm_network_available                                = iota
	wlan_notification_acm_disconnecting                                    = iota
	wlan_notification_acm_disconnected                                     = iota
	wlan_notification_acm_adhoc_network_state_change                       = iota
	wlan_notification_acm_profile_unblocked                                = iota
	wlan_notification_acm_screen_power_change                              = iota
	wlan_notification_acm_profile_blocked                                  = iota
	wlan_notification_acm_scan_list_refresh                                = iota
	wlan_notification_acm_end                                              = iota
)
const (
	WLAN_AVAILABLE_NETWORK_INCLUDE_ALL_ADHOC_PROFILES         = 0x00000001
	WLAN_AVAILABLE_NETWORK_INCLUDE_ALL_MANUAL_HIDDEN_PROFILES = 0x00000002
	WLAN_AVAILABLE_NETWORK_CONNECTED                          = 1
	MAX_INDEX                                                 = 1000
	WLAN_MAX_NAME_LENGTH                                      = 256
	WLAN_PROFILE_GET_PLAINTEXT_KEY                            = 4

	ERROR_SUCCESS                        = 0
	ERROR_INVALID_PARAMETER              = 87
	ERROR_INVALID_HANDLE                 = 6
	ERROR_INVALID_STATE                  = 5023
	ERROR_NOT_FOUND                      = 1168
	ERROR_NOT_ENOUGH_MEMORY              = 8
	ERROR_ACCESS_DENIED                  = 5
	ERROR_NOT_SUPPORTED                  = 50
	ERROR_SERVICE_NOT_ACTIVE             = 1062
	ERROR_NDIS_DOT11_AUTO_CONFIG_ENABLED = 0x80342000
	ERROR_NDIS_DOT11_MEDIA_IN_USE        = 0x80342001
	ERROR_NDIS_DOT11_POWER_STATE_INVALID = 0x80342002
	ERROR_ALREADY_EXISTS                 = 183
	ERROR_BAD_PROFILE                    = 1206
	ERROR_NO_MATCH                       = 1169
	ERROR_GEN_FAILURE                    = 31

	WLAN_NOTIFICATION_SOURCE_NONE     uint32 = 0
	WLAN_NOTIFICATION_SOURCE_ALL             = 0x0000FFFF
	WLAN_NOTIFICATION_SOURCE_ACM             = 0x00000008
	WLAN_NOTIFICATION_SOURCE_HNWK            = 0x00000080
	WLAN_NOTIFICATION_SOURCE_ONEX            = 0x00000004
	WLAN_NOTIFICATION_SOURCE_MSM             = 0x00000010
	WLAN_NOTIFICATION_SOURCE_SECURITY        = 0x00000020
	WLAN_NOTIFICATION_SOURCE_IHV             = 0x00000040

	FORMAT_MESSAGE_FROM_SYSTEM uint32 = 0x00001000
)
