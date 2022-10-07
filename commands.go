package hs110

type Cmd string

const (
	// System Commands
	// ========================================

	// Get System Info (Software & Hardware Versions, MAC, deviceID, hwID etc.)
	cmdInfo = `{"system":{"get_sysinfo":null}}`

	// Reboot
	cmdReboot = `{"system":{"reboot":{"delay":1}}}`

	// Reset (To Factory Settings)
	cmdReset = `{"system":{"reset":{"delay":1}}}`

	// Turn On
	cmdOn = `{"system":{"set_relay_state":{"state":1}}}`

	// Turn Off
	cmdOff = `{"system":{"set_relay_state":{"state":0}}}`

	// Turn On Device LED (Night mode)
	cmdLEDOn = `{"system":{"set_led_off":{"off":0}}}`

	// Turn Off Device LED (Night mode)
	cmdLEDOff = `{"system":{"set_led_off":{"off":1}}}`

	// Set Device Alias (use with fmt.Sprintf)
	cmdSetAlias = `{"system":{"set_dev_alias":{"alias":"%s"}}}`

	// Set MAC Address (use with fmt.Sprintf, format all upper case with "-" separator)
	cmdSetMAC = `{"system":{"set_mac_addr":{"mac":"%s"}}}`

	// Set Device ID (use with fmt.Sprintf)
	cmdSetDeviceID = `{"system":{"set_device_id":{"deviceId":"%s"}}}`

	// Set Hardware ID (use with fmt.Sprintf)
	cmdSetHardwareID = `{"system":{"set_hw_id":{"hwId":"%s"}}}`

	// Set Location (use with fmt.Sprintf)
	cmdSetLocation = `{"system":{"set_dev_location":{"longitude":%f,"latitude":%f}}}`

	// Perform uBoot Bootloader Check
	cmdBootloaderCheck = `{"system":{"test_check_uboot":null}}`

	// Get Device Icon
	cmdDeviceIcon = `{"system":{"get_dev_icon":null}}`

	// Set Device Icon (use with fmt.Sprintf)
	cmdSetDeviceIcon = `{"system":{"set_dev_icon":{"icon":"%s","hash":"%s"}}}`

	// Set Test Mode (command only accepted coming from IP 192.168.1.100)
	cmdSetTestMode = `{"system":{"set_test_mode":{"enable":1}}}`

	// Download Firmware from URL (use with fmt.Sprintf)
	cmdDownloadFirmwareFromURL = `{"system":{"download_firmware":{"url":"%s"}}}`

	// Get Download State
	cmdDownloadState = `{"system":{"get_download_state":{}}}`

	// Flash Downloaded Firmware
	cmdFlashFirmware = `{"system":{"flash_firmware":{}}}`

	// Check Config
	cmdCheckConfig = `{"system":{"check_new_config":null}}`

	// WLAN Commands
	// ========================================

	// Scan for list of available APs
	cmdScanAccessPoints = `{"netif":{"get_scaninfo":{"refresh":1}}}`

	// Connect to AP with given SSID and Password (use with fmt.Sprintf)
	cmdConnectAccessPoint = `{"netif":{"set_stainfo":{"ssid":"%s","password":"%s","key_type":3}}}`

	// Cloud Commands
	// ========================================

	// Get Cloud Info (Server, Username, Connection Status)
	cmdCloudInfo = `{"cnCloud":{"get_info":null}}`

	// Get Firmware List from Cloud Server
	cmdListFirmware = `{"cnCloud":{"get_intl_fw_list":{}}}`

	// Set Server URL (use with fmt.Sprintf)
	cmdSetServerURL = `{"cnCloud":{"set_server_url":{"server":"%s"}}}`

	// Connect with Cloud username & Password (use with fmt.Sprintf)
	cmdConnectToCloud = `{"cnCloud":{"bind":{"username":"%s", "password":"%s"}}}`

	// Unregister Device from Cloud Account
	cmdUnregister = `{"cnCloud":{"unbind":null}}`

	// Time Commands
	// ========================================

	// Get Time
	cmdTime = `{"time":{"get_time":null}}`

	// Get Timezone
	cmdTimezone = `{"time":{"get_timezone":null}}`

	// Set Timezone (use with fmt.Sprintf)
	cmdSetTimezone = `{"time":{"set_timezone":{"year":%d,"month":%d,"mday":%d,"hour":%d,"min":%d,"sec":%d,"index":%d}}}`

	// EMeter Energy Usage Statistics Commands
	// (for TP-Link HS110)
	// ========================================

	// Get Realtime Current and Voltage Reading
	cmdEnergy = `{"emeter":{"get_realtime":{}}}`

	// Get EMeter VGain and IGain Settings
	cmdEmeter = `{"emeter":{"get_vgain_igain":{}}}`

	// Set EMeter VGain and Igain
	cmdSetEmeter = `{"emeter":{"set_vgain_igain":{"vgain":%d,"igain":%d}}}`

	// Start EMeter Calibration
	cmdCalibrateEmeter = `{"emeter":{"start_calibration":{"vtarget":%d,"itarget":%d}}}`

	// Get Daily Statistic for given Month (use with fmt.Sprintf)
	cmdEmeterDailyStatistics = `{"emeter":{"get_daystat":{"month":%d,"year":%d}}}`

	// Get Montly Statistic for given Year (use with fmt.Sprintf
	cmdEmeterMonthlyStatistics = `{"emeter":{"get_monthstat":{"year":%d}}}`

	// Erase All EMeter Statistics
	cmdEraseEmeterStatistics = `{"emeter":{"erase_emeter_stat":null}}`

	// Schedule Commands
	// (action to perform regularly on given weekdays)
	// ========================================

	// Get Next Scheduled Action
	cmdNextScheduledAction = `{"schedule":{"get_next_action":null}}`

	// Get Schedule Rules List
	cmdScheduleRules = `{"schedule":{"get_rules":null}}`

	// Add New Schedule Rule (use with fmt.Sprintf)
	cmdNewScheduleRule = `{"schedule":{"add_rule":{"stime_opt":%d,"wday":[%d,%d,%d,%d,%d,%d,%d],"smin":%d,"enable":1,"repeat":%d,"etime_opt":-1,"name":"lights on","eact":-1,"month":%d,"sact":%d,"year":%d,"longitude":%d,"day":%d,"force":%d,"latitude":%d,"emin":%d},"set_overall_enable":{"enable":1}}}`

	// Edit Schedule Rule with given ID (use with fmt.Sprintf)
	cmdEditScheduleRule = `{"schedule":{"edit_rule":{"stime_opt":%d,"wday":[%d,%d,%d,%d,%d,%d,0],"smin":%d,"enable":1,"repeat":%d,"etime_opt":-1,"id":"%s","name":"lights on","eact":-1,"month":%d,"sact":%d,"year":%d,"longitude":%d,"day":%d,"force":%d,"latitude":%d,"emin":%d}}}`

	// Delete Schedule Rule with given ID (use with fmt.Sprintf)
	cmdDeleteScheduleRule = `{"schedule":{"delete_rule":{"id":"%s"}}}`

	// Delete All Schedule Rules and Erase Statistics
	cmdDeleteAllScheduleRules = `{"schedule":{"delete_all_rules":null,"erase_runtime_stat":null}}`

	// Countdown Rule Commands
	// (action to perform after number of seconds)
	// ========================================

	// Get Rule (only one allowed)
	cmdCountdownRule = `{"count_down":{"get_rules":null}}`

	// Add New Countdown Rule (use with fmt.Sprintf)
	cmdNewCountdownRule = `{"count_down":{"add_rule":{"enable":1,"delay":%d,"act":%d,"name":"%s"}}}`

	// Edit Countdown Rule with given ID (use with fmt.Sprintf)
	cmdEditCountdownRule = `{"count_down":{"edit_rule":{"enable":1,"id":"%s","delay":%d,"act":%d,"name":"%s"}}}`

	// Delete Countdown Rule with given ID (use with fmt.Sprintf)
	cmdDeleteCountdownRule = `{"count_down":{"delete_rule":{"id":"%s"}}}`

	// Delete All Coundown Rules
	cmdDeleteAllCountdownRules = `{"count_down":{"delete_all_rules":null}}`

	// Anti-Theft Rule Commands (aka Away Mode)
	// (period of time during which device will be randomly turned on and off to deter thieves)
	// ========================================

	// Get Anti-Theft Rules List
	cmdAntiTheftRules = `{"anti_theft":{"get_rules":null}}`

	// Add New Anti-Theft Rule (use with fmt.Sprintf)
	cmdNewAntiTheftRule = `{"anti_theft":{"add_rule":{"stime_opt":%d,"wday":[%d,%d,%d,%d,%d,%d,%d],"smin":%d,"enable":1,"frequency":%d,"repeat":%d,"etime_opt":%d,"duration":%d,"name":"test","lastfor":%d,"month":%d,"year":%d,"longitude":%d,"day":%d,"latitude":%d,"force":%d,"emin":1047},"set_overall_enable":1}}`

	// Edit Anti-Theft Rule with given ID (use with fmt.Sprintf)
	cmdEditAntiTheftRule = `{"anti_theft":{"edit_rule":{"stime_opt":%d,"wday":[%d,%d,%d,%d,%d,%d,%d],"smin":%d,"enable":1,"frequency":%d,"repeat":%d,"etime_opt":%d,"id":"%s","duration":%d,"name":"test","lastfor":%d,"month":%d,"year":%d,"longitude":%d,"day":%d,"latitude":%d,"force":%d,"emin":%d},"set_overall_enable":1}}`

	// Delete Anti-Theft Rule with given ID (use with fmt.Sprintf)
	cmdDeleteAntiTheftRule = `{"anti_theft":{"delete_rule":{"id":"%s"}}}`

	// Delete All Anti-Theft Rules
	cmdDeleteAllAntiTheftRules = `{"anti_theft":{"delete_all_rules":null}}`
)
