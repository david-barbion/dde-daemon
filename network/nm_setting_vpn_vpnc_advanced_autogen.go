// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnVpncAdvancedKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_VPNC_KEY_DOMAIN:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_VENDOR:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_APP_VERSION:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_SINGLE_DES:
		t = ktypeBoolean
	case NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION:
		t = ktypeBoolean
	case NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_DHGROUP:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD:
		t = ktypeString
	case NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT:
		t = ktypeUint32
	case NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT:
		t = ktypeUint32
	case NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT:
		t = ktypeUint32
	}
	return
}

// Check is key in current setting section
func isKeyInSettingVpnVpncAdvanced(key string) bool {
	switch key {
	case NM_SETTING_VPN_VPNC_KEY_DOMAIN:
		return true
	case NM_SETTING_VPN_VPNC_KEY_VENDOR:
		return true
	case NM_SETTING_VPN_VPNC_KEY_APP_VERSION:
		return true
	case NM_SETTING_VPN_VPNC_KEY_SINGLE_DES:
		return true
	case NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION:
		return true
	case NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE:
		return true
	case NM_SETTING_VPN_VPNC_KEY_DHGROUP:
		return true
	case NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD:
		return true
	case NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT:
		return true
	case NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT:
		return true
	case NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnVpncAdvancedDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_VPNC_KEY_DOMAIN:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_VENDOR:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_APP_VERSION:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_SINGLE_DES:
		value = false
	case NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION:
		value = false
	case NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_DHGROUP:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD:
		value = ""
	case NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT:
		value = 0
	case NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT:
		value = 0
	case NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT:
		value = 0
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnVpncAdvancedKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnVpncAdvancedKeyJSON: invalide key", key)
	case NM_SETTING_VPN_VPNC_KEY_DOMAIN:
		value = getSettingVpnVpncKeyDomainJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_VENDOR:
		value = getSettingVpnVpncKeyVendorJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_APP_VERSION:
		value = getSettingVpnVpncKeyAppVersionJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_SINGLE_DES:
		value = getSettingVpnVpncKeySingleDesJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION:
		value = getSettingVpnVpncKeyNoEncryptionJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE:
		value = getSettingVpnVpncKeyNatTraversalModeJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_DHGROUP:
		value = getSettingVpnVpncKeyDhgroupJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD:
		value = getSettingVpnVpncKeyPerfectForwardJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT:
		value = getSettingVpnVpncKeyLocalPortJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT:
		value = getSettingVpnVpncKeyDpdIdleTimeoutJSON(data)
	case NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT:
		value = getSettingVpnVpncKeyCiscoUdpEncapsPortJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnVpncAdvancedKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnVpncAdvancedKeyJSON: invalide key", key)
	case NM_SETTING_VPN_VPNC_KEY_DOMAIN:
		err = setSettingVpnVpncKeyDomainJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_VENDOR:
		err = setSettingVpnVpncKeyVendorJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_APP_VERSION:
		err = setSettingVpnVpncKeyAppVersionJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_SINGLE_DES:
		err = setSettingVpnVpncKeySingleDesJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION:
		err = setSettingVpnVpncKeyNoEncryptionJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE:
		err = setSettingVpnVpncKeyNatTraversalModeJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_DHGROUP:
		err = setSettingVpnVpncKeyDhgroupJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD:
		err = setSettingVpnVpncKeyPerfectForwardJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT:
		err = setSettingVpnVpncKeyLocalPortJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT:
		err = setSettingVpnVpncKeyDpdIdleTimeoutJSON(data, valueJSON)
	case NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT:
		err = setSettingVpnVpncKeyCiscoUdpEncapsPortJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnVpncKeyDomainExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN)
}
func isSettingVpnVpncKeyVendorExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR)
}
func isSettingVpnVpncKeyAppVersionExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION)
}
func isSettingVpnVpncKeySingleDesExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES)
}
func isSettingVpnVpncKeyNoEncryptionExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION)
}
func isSettingVpnVpncKeyNatTraversalModeExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE)
}
func isSettingVpnVpncKeyDhgroupExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP)
}
func isSettingVpnVpncKeyPerfectForwardExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD)
}
func isSettingVpnVpncKeyLocalPortExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT)
}
func isSettingVpnVpncKeyDpdIdleTimeoutExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT)
}
func isSettingVpnVpncKeyCiscoUdpEncapsPortExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT)
}

// Ensure section and key exists and not empty
func ensureSectionSettingVpnVpncAdvancedExists(data connectionData, errs sectionErrors, relatedKey string) {
	if !isSettingSectionExists(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME))
	}
	sectionData, _ := data[NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME]
	if len(sectionData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME))
	}
}
func ensureSettingVpnVpncKeyDomainNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyDomainExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyDomain(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyVendorNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyVendorExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyVendor(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyAppVersionNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyAppVersionExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyAppVersion(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeySingleDesNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeySingleDesExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnVpncKeyNoEncryptionNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyNoEncryptionExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnVpncKeyNatTraversalModeNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyNatTraversalModeExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyNatTraversalMode(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyDhgroupNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyDhgroupExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyDhgroup(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyPerfectForwardNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyPerfectForwardExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnVpncKeyPerfectForward(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnVpncKeyLocalPortNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyLocalPortExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnVpncKeyDpdIdleTimeoutNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyDpdIdleTimeoutExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnVpncKeyCiscoUdpEncapsPortNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnVpncKeyCiscoUdpEncapsPortExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingVpnVpncKeyDomain(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyVendor(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyAppVersion(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeySingleDes(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnVpncKeyNoEncryption(data connectionData) (value bool) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION)
	value = interfaceToBoolean(ivalue)
	return
}
func getSettingVpnVpncKeyNatTraversalMode(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyDhgroup(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyPerfectForward(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnVpncKeyLocalPort(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnVpncKeyDpdIdleTimeout(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT)
	value = interfaceToUint32(ivalue)
	return
}
func getSettingVpnVpncKeyCiscoUdpEncapsPort(data connectionData) (value uint32) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT)
	value = interfaceToUint32(ivalue)
	return
}

// Setter
func setSettingVpnVpncKeyDomain(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN, value)
}
func setSettingVpnVpncKeyVendor(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR, value)
}
func setSettingVpnVpncKeyAppVersion(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION, value)
}
func setSettingVpnVpncKeySingleDes(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES, value)
}
func setSettingVpnVpncKeyNoEncryption(data connectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION, value)
}
func setSettingVpnVpncKeyNatTraversalMode(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE, value)
}
func setSettingVpnVpncKeyDhgroup(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP, value)
}
func setSettingVpnVpncKeyPerfectForward(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD, value)
}
func setSettingVpnVpncKeyLocalPort(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT, value)
}
func setSettingVpnVpncKeyDpdIdleTimeout(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT, value)
}
func setSettingVpnVpncKeyCiscoUdpEncapsPort(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT, value)
}

// JSON Getter
func getSettingVpnVpncKeyDomainJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_DOMAIN))
	return
}
func getSettingVpnVpncKeyVendorJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_VENDOR))
	return
}
func getSettingVpnVpncKeyAppVersionJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_APP_VERSION))
	return
}
func getSettingVpnVpncKeySingleDesJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_SINGLE_DES))
	return
}
func getSettingVpnVpncKeyNoEncryptionJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION))
	return
}
func getSettingVpnVpncKeyNatTraversalModeJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE))
	return
}
func getSettingVpnVpncKeyDhgroupJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_DHGROUP))
	return
}
func getSettingVpnVpncKeyPerfectForwardJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD))
	return
}
func getSettingVpnVpncKeyLocalPortJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT))
	return
}
func getSettingVpnVpncKeyDpdIdleTimeoutJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT))
	return
}
func getSettingVpnVpncKeyCiscoUdpEncapsPortJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT))
	return
}

// JSON Setter
func setSettingVpnVpncKeyDomainJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_DOMAIN))
}
func setSettingVpnVpncKeyVendorJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_VENDOR))
}
func setSettingVpnVpncKeyAppVersionJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_APP_VERSION))
}
func setSettingVpnVpncKeySingleDesJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_SINGLE_DES))
}
func setSettingVpnVpncKeyNoEncryptionJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION))
}
func setSettingVpnVpncKeyNatTraversalModeJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE))
}
func setSettingVpnVpncKeyDhgroupJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_DHGROUP))
}
func setSettingVpnVpncKeyPerfectForwardJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD))
}
func setSettingVpnVpncKeyLocalPortJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT))
}
func setSettingVpnVpncKeyDpdIdleTimeoutJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT))
}
func setSettingVpnVpncKeyCiscoUdpEncapsPortJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT, valueJSON, getSettingVpnVpncAdvancedKeyType(NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT))
}

// Logic JSON Setter

// Remover
func removeSettingVpnVpncKeyDomain(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DOMAIN)
}
func removeSettingVpnVpncKeyVendor(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_VENDOR)
}
func removeSettingVpnVpncKeyAppVersion(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_APP_VERSION)
}
func removeSettingVpnVpncKeySingleDes(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_SINGLE_DES)
}
func removeSettingVpnVpncKeyNoEncryption(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NO_ENCRYPTION)
}
func removeSettingVpnVpncKeyNatTraversalMode(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_NAT_TRAVERSAL_MODE)
}
func removeSettingVpnVpncKeyDhgroup(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DHGROUP)
}
func removeSettingVpnVpncKeyPerfectForward(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_PERFECT_FORWARD)
}
func removeSettingVpnVpncKeyLocalPort(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_LOCAL_PORT)
}
func removeSettingVpnVpncKeyDpdIdleTimeout(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_DPD_IDLE_TIMEOUT)
}
func removeSettingVpnVpncKeyCiscoUdpEncapsPort(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME, NM_SETTING_VPN_VPNC_KEY_CISCO_UDP_ENCAPS_PORT)
}
