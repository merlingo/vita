package types_test

import (
	"testing"

	"vita/x/labour/types"

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

				TaskList: []types.Task{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				TaskCount: 2,
				ActivityList: []types.Activity{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ActivityCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated task",
			genState: &types.GenesisState{
				TaskList: []types.Task{
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
			desc: "invalid task count",
			genState: &types.GenesisState{
				TaskList: []types.Task{
					{
						Id: 1,
					},
				},
				TaskCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated activity",
			genState: &types.GenesisState{
				ActivityList: []types.Activity{
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
			desc: "invalid activity count",
			genState: &types.GenesisState{
				ActivityList: []types.Activity{
					{
						Id: 1,
					},
				},
				ActivityCount: 0,
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
