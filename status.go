package hs110

import (
	"encoding/json"
	"fmt"
)

// Status returns the status of the plug's relay. There are many additional fields returned by the plug, though most
// of these have been omitted. Below is an example of the returned payload.
//
//	{
//	  "system": {
//	    "get_sysinfo": {
//	      "sw_ver": "1.5.4 Build 180815 Rel.121440",
//	      "hw_ver": "2.0",
//	      "type": "IOT.SMARTPLUGSWITCH",
//	      "model": "HS110(EU)",
//	      "mac": "AA:AA:AA:AA:AA:AA",
//	      "dev_name": "Smart Wi-Fi Plug With Energy Monitoring",
//	      "alias": "plug-1",
//	      "relay_state": 1,
//	      "on_time": 26,
//	      "active_mode": "none",
//	      "feature": "TIM:ENE",
//	      "updating": 0,
//	      "icon_hash": "",
//	      "rssi": -50,
//	      "led_off": 0,
//	      "longitude_i": 100000,
//	      "latitude_i": 500000,
//	      "hwId": "344A516FE63C275F9458DA25C2CCC5A0",
//	      "fwId": "00000000000000000000000000000000",
//	      "deviceId": "80061B5970839C523B9EABB16F14C76828E2C220",
//	      "oemId": "2807A14DAA86E4E001FD7CAF42868B5F",
//	      "next_action": {
//	        "type": -1
//	      },
//	      "err_code": 0
//	    }
//	  }
//	}
func (p *Plug) Status() (bool, error) {
	res, err := p.sendCmd(cmdInfo)
	if err != nil {
		return false, fmt.Errorf("sending request: %w", err)
	}
	fmt.Println(string(res))

	var state struct {
		System struct {
			Sysinfo struct {
				RelayState int `json:"relay_state"`
			} `json:"get_sysinfo"`
		} `json:"system"`
	}
	if err := json.Unmarshal(res, &state); err != nil {
		return false, fmt.Errorf("unmarshaling response: %w", err)
	}
	if st := state.System.Sysinfo.RelayState; st == 1 {
		return true, nil
	} else if st == 0 {
		return false, nil
	} else {
		return false, fmt.Errorf("unknown state %d", st)
	}
}
