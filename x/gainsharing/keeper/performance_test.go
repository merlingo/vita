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

func createNPerformance(keeper keeper.Keeper, ctx context.Context, n int) []types.Performance {
	items := make([]types.Performance, n)
	for i := range items {
		items[i].Id = keeper.AppendPerformance(ctx, items[i])
	}
	return items
}

func TestPerformanceGet(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNPerformance(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPerformance(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPerformanceRemove(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNPerformance(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePerformance(ctx, item.Id)
		_, found := keeper.GetPerformance(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPerformanceGetAll(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNPerformance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPerformance(ctx)),
	)
}

func TestPerformanceCount(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	items := createNPerformance(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPerformanceCount(ctx))
}
