package vita_test

import (
	"testing"

	keepertest "vita/testutil/keeper"
	"vita/testutil/nullify"
	vita "vita/x/vita/module"
	"vita/x/vita/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VitaKeeper(t)
	vita.InitGenesis(ctx, k, genesisState)
	got := vita.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
