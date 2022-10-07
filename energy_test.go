package hs110

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_parseEnergyPayload(t *testing.T) {
	for _, tc := range []struct {
		inp       string
		shouldErr bool
		exp       *EnergyReadout
	}{
		{
			inp:       ``,
			shouldErr: true,
		},
		{
			inp: `{
			  "emeter": {
				"get_realtime": {
				  "voltage_mv": 0,
				  "current_ma": 0,
				  "power_mw": 0,
				  "total_wh": 0,
				  "err_code": 1
				}
			  }
			}`,
			shouldErr: true,
		},
		{
			inp: `{
			  "emeter": {
				"get_realtime": {
				  "voltage_mv": 241217,
				  "current_ma": 1917,
				  "power_mw": 462498,
				  "total_wh": 1596,
				  "err_code": 0
				}
			  }
			}`,
			exp: &EnergyReadout{
				VoltageMV: 241217,
				CurrentMA: 1917,
				PowerMW:   462498,
				TotalWH:   1596,
				ErrCode:   0,
			},
		},
		{
			inp: `{
			  "emeter": {
				"get_realtime": {
				  "current": 0.083255,
				  "voltage": 241.225492,
				  "power": 10.21876,
				  "total": 0.054,
				  "err_code": 0
				}
			  }
			}`,
			exp: &EnergyReadout{
				VoltageMV: 241225,
				CurrentMA: 83,
				PowerMW:   10218,
				TotalWH:   54,
				ErrCode:   0,
			},
		},
	} {
		got, err := parseEnergyPayload([]byte(tc.inp))
		if tc.shouldErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		assert.Equal(t, tc.exp, got)
	}
}
