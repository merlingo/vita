package labour_test

import (
	"testing"

	keepertest "vita/testutil/keeper"
	"vita/testutil/nullify"
	labour "vita/x/labour/module"
	"vita/x/labour/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LabourKeeper(t)
	labour.InitGenesis(ctx, k, genesisState)
	got := labour.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	require.Equal(t, genesisState.TaskCount, got.TaskCount)
	require.ElementsMatch(t, genesisState.ActivityList, got.ActivityList)
	require.Equal(t, genesisState.ActivityCount, got.ActivityCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
