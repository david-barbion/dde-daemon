package main

const NM_SETTING_PPPOE_SETTING_NAME = "pppoe"

const (
	NM_SETTING_PPPOE_SERVICE        = "service"
	NM_SETTING_PPPOE_USERNAME       = "username"
	NM_SETTING_PPPOE_PASSWORD       = "password"
	NM_SETTING_PPPOE_PASSWORD_FLAGS = "password-flags"
)

func newPppoeConnection(id, username string) (uuid string) {
	logger.Debugf("new pppoe connection, id=%s", id)
	uuid = newUUID()
	data := newPppoeConnectionData(id, uuid)
	setSettingPppoeUsername(data, username)
	nmAddConnection(data)
	return
}

func newPppoeConnectionData(id, uuid string) (data connectionData) {
	data = make(connectionData)

	addSettingSection(data, sectionConnection)
	setSettingConnectionId(data, id)
	setSettingConnectionUuid(data, uuid)
	setSettingConnectionType(data, NM_SETTING_PPPOE_SETTING_NAME)
	setSettingConnectionAutoconnect(data, false)

	initSettingSectionWired(data)

	addSettingSection(data, sectionPppoe)

	addSettingSection(data, sectionPpp)
	logicSetSettingVkPppEnableLcpEcho(data, true)

	initSettingSectionIpv4(data)
	return
}

// Get available keys
func getSettingPppoeAvailableKeys(data connectionData) (keys []string) {
	keys = appendAvailableKeys(data, keys, sectionPppoe, NM_SETTING_PPPOE_SERVICE)
	keys = appendAvailableKeys(data, keys, sectionPppoe, NM_SETTING_PPPOE_USERNAME)
	keys = appendAvailableKeys(data, keys, sectionPppoe, NM_SETTING_PPPOE_PASSWORD)
	return
}

// Get available values
func getSettingPppoeAvailableValues(data connectionData, key string) (values []kvalue) {
	return
}

// Check whether the values are correct
func checkSettingPppoeValues(data connectionData) (errs sectionErrors) {
	errs = make(map[string]string)

	// check username
	ensureSettingPppoeUsernameNoEmpty(data, errs)

	return
}
