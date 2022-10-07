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
	bs, err := p.sendCmd(cmdEnergy)
	if err != nil {
		return nil, err
	}
	return parseEnergyPayload(bs)
}

func parseEnergyPayload(bs []byte) (*EnergyReadout, error) {
	var data struct {
		Emeter struct {
			GetRealtime struct {
				VoltageMV *int `json:"voltage_mv"`
				CurrentMA *int `json:"current_ma"`
				PowerMW   *int `json:"power_mw"`
				TotalWH   *int `json:"total_wh"`

				Voltage *float64 `json:"voltage"`
				Current *float64 `json:"current"`
				Power   *float64 `json:"power"`
				Total   *float64 `json:"total"`

				ErrCode int `json:"err_code"`
			} `json:"get_realtime"`
		} `json:"emeter"`
	}
	if err := json.Unmarshal(bs, &data); err != nil {
		return nil, err
	}
	readout := data.Emeter.GetRealtime
	if readout.ErrCode > 0 {
		return nil, fmt.Errorf("got error code %d", readout.ErrCode)
	}

	var res EnergyReadout
	if readout.VoltageMV == nil {
		if readout.Voltage == nil || readout.Current == nil || readout.Power == nil || readout.Total == nil {
			return nil, fmt.Errorf("expected whole unit floats but did not get all expected fields")
		}
		res.VoltageMV = int(*readout.Voltage * 1000)
		res.CurrentMA = int(*readout.Current * 1000)
		res.PowerMW = int(*readout.Power * 1000)
		res.TotalWH = int(*readout.Total * 1000)
	} else if readout.CurrentMA == nil || readout.PowerMW == nil || readout.TotalWH == nil {
		return nil, fmt.Errorf("expected milli unit ints didn't get expected fields")
	} else {
		res.VoltageMV = *readout.VoltageMV
		res.CurrentMA = *readout.CurrentMA
		res.PowerMW = *readout.PowerMW
		res.TotalWH = *readout.TotalWH
	}
	return &res, nil
}
