package gainsharing_test

import (
	"testing"

	keepertest "vita/testutil/keeper"
	"vita/testutil/nullify"
	gainsharing "vita/x/gainsharing/module"
	"vita/x/gainsharing/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GainsharingKeeper(t)
	gainsharing.InitGenesis(ctx, k, genesisState)
	got := gainsharing.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.MechanismList, got.MechanismList)
	require.Equal(t, genesisState.MechanismCount, got.MechanismCount)
	require.ElementsMatch(t, genesisState.PerformanceList, got.PerformanceList)
	require.Equal(t, genesisState.PerformanceCount, got.PerformanceCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
