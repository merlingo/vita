package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "vita/testutil/keeper"
	"vita/x/vita/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VitaKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
