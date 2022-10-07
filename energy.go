package hs110

import (
	"encoding/json"
	"fmt"
)

type EnergyReadout struct {
	VoltageMV int `json:"voltage_mv"`
	CurrentMA int `json:"current_ma"`
	PowerMW   int `json:"power_mw"`
	TotalWH   int `json:"total_wh"`
	ErrCode   int `json:"err_code"`
}

func (r *EnergyReadout) String() string {
	res := fmt.Sprintf("voltage mv:\t%d", r.VoltageMV)
	res += fmt.Sprintf("\ncurrent ma:\t%d", r.CurrentMA)
	res += fmt.Sprintf("\npower mw:\t%d", r.PowerMW)
	res += fmt.Sprintf("\ntotal wh:\t%d", r.TotalWH)
	return res
}

func (p *Plug) Energy() (*EnergyReadout, error) {
	res, err := p.sendCmd(cmdEnergy)
	if err != nil {
		return nil, err
	}
	var data struct {
		Emeter struct {
			GetRealtime EnergyReadout `json:"get_realtime"`
		} `json:"emeter"`
	}
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, err
	}
	readout := data.Emeter.GetRealtime
	if readout.ErrCode > 0 {
		return nil, fmt.Errorf("got error code %d", readout.ErrCode)
	}
	return &data.Emeter.GetRealtime, nil
}
