package types_test

import (
	"testing"

	"vita/x/gainsharing/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				MechanismList: []types.Mechanism{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MechanismCount: 2,
				PerformanceList: []types.Performance{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				PerformanceCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated mechanism",
			genState: &types.GenesisState{
				MechanismList: []types.Mechanism{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid mechanism count",
			genState: &types.GenesisState{
				MechanismList: []types.Mechanism{
					{
						Id: 1,
					},
				},
				MechanismCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated performance",
			genState: &types.GenesisState{
				PerformanceList: []types.Performance{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid performance count",
			genState: &types.GenesisState{
				PerformanceList: []types.Performance{
					{
						Id: 1,
					},
				},
				PerformanceCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
