package keeper_test

import (
	"context"
	"testing"

	keepertest "vita/testutil/keeper"
	"vita/testutil/nullify"
	"vita/x/gainsharing/keeper"
	"vita/x/gainsharing/types"

	"github.com/stretchr/testify/require"
)

func createNMechanism(keeper keeper.Keeper, ctx context.Context, n int) []types.Mechanism {
	items := make([]types.Mechanism, n)
	for i := range items {
		items[i].Id = keeper.AppendMechanism(ctx, items[i])
	}
	return items
}

func TestMechanismGet(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNMechanism(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMechanism(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMechanismRemove(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNMechanism(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMechanism(ctx, item.Id)
		_, found := keeper.GetMechanism(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMechanismGetAll(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNMechanism(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMechanism(ctx)),
	)
}

func TestMechanismCount(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNMechanism(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMechanismCount(ctx))
}
