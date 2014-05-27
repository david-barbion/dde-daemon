// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnOpenvpnSecurityKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_OPENVPN_KEY_CIPHER:
		t = ktypeString
	case NM_SETTING_VPN_OPENVPN_KEY_AUTH:
		t = ktypeString
	}
	return
}

// Check is key in current setting section
func isKeyInSettingVpnOpenvpnSecurity(key string) bool {
	switch key {
	case NM_SETTING_VPN_OPENVPN_KEY_CIPHER:
		return true
	case NM_SETTING_VPN_OPENVPN_KEY_AUTH:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnOpenvpnSecurityDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_VPN_OPENVPN_KEY_CIPHER:
		value = ""
	case NM_SETTING_VPN_OPENVPN_KEY_AUTH:
		value = ""
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnOpenvpnSecurityKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingVpnOpenvpnSecurityKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENVPN_KEY_CIPHER:
		value = getSettingVpnOpenvpnKeyCipherJSON(data)
	case NM_SETTING_VPN_OPENVPN_KEY_AUTH:
		value = getSettingVpnOpenvpnKeyAuthJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnOpenvpnSecurityKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingVpnOpenvpnSecurityKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENVPN_KEY_CIPHER:
		err = setSettingVpnOpenvpnKeyCipherJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENVPN_KEY_AUTH:
		err = setSettingVpnOpenvpnKeyAuthJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnOpenvpnKeyCipherExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER)
}
func isSettingVpnOpenvpnKeyAuthExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH)
}

// Ensure section and key exists and not empty
func ensureSectionSettingVpnOpenvpnSecurityExists(data connectionData, errs sectionErrors, relatedKey string) {
	if !isSettingSectionExists(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME))
	}
	sectionData, _ := data[NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME]
	if len(sectionData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME))
	}
}
func ensureSettingVpnOpenvpnKeyCipherNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyCipherExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenvpnKeyCipher(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenvpnKeyAuthNoEmpty(data connectionData, errs sectionErrors) {
	if !isSettingVpnOpenvpnKeyAuthExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenvpnKeyAuth(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingVpnOpenvpnKeyCipher(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER)
	value = interfaceToString(ivalue)
	return
}
func getSettingVpnOpenvpnKeyAuth(data connectionData) (value string) {
	ivalue := getSettingKey(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH)
	value = interfaceToString(ivalue)
	return
}

// Setter
func setSettingVpnOpenvpnKeyCipher(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER, value)
}
func setSettingVpnOpenvpnKeyAuth(data connectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH, value)
}

// JSON Getter
func getSettingVpnOpenvpnKeyCipherJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER, getSettingVpnOpenvpnSecurityKeyType(NM_SETTING_VPN_OPENVPN_KEY_CIPHER))
	return
}
func getSettingVpnOpenvpnKeyAuthJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH, getSettingVpnOpenvpnSecurityKeyType(NM_SETTING_VPN_OPENVPN_KEY_AUTH))
	return
}

// JSON Setter
func setSettingVpnOpenvpnKeyCipherJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER, valueJSON, getSettingVpnOpenvpnSecurityKeyType(NM_SETTING_VPN_OPENVPN_KEY_CIPHER))
}
func setSettingVpnOpenvpnKeyAuthJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH, valueJSON, getSettingVpnOpenvpnSecurityKeyType(NM_SETTING_VPN_OPENVPN_KEY_AUTH))
}

// Logic JSON Setter

// Remover
func removeSettingVpnOpenvpnKeyCipher(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_CIPHER)
}
func removeSettingVpnOpenvpnKeyAuth(data connectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME, NM_SETTING_VPN_OPENVPN_KEY_AUTH)
}
