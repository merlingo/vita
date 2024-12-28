package keeper_test

import (
	"context"
	"testing"

	keepertest "vita/testutil/keeper"
	"vita/testutil/nullify"
	"vita/x/labour/keeper"
	"vita/x/labour/types"

	"github.com/stretchr/testify/require"
)

func createNActivity(keeper keeper.Keeper, ctx context.Context, n int) []types.Activity {
	items := make([]types.Activity, n)
	for i := range items {
		items[i].Id = keeper.AppendActivity(ctx, items[i])
	}
	return items
}

func TestActivityGet(t *testing.T) {
	keeper, ctx := keepertest.LabourKeeper(t)
	items := createNActivity(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetActivity(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestActivityRemove(t *testing.T) {
	keeper, ctx := keepertest.LabourKeeper(t)
	items := createNActivity(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveActivity(ctx, item.Id)
		_, found := keeper.GetActivity(ctx, item.Id)
		require.False(t, found)
	}
}

func TestActivityGetAll(t *testing.T) {
	keeper, ctx := keepertest.LabourKeeper(t)
	items := createNActivity(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllActivity(ctx)),
	)
}

func TestActivityCount(t *testing.T) {
	keeper, ctx := keepertest.LabourKeeper(t)
	items := createNActivity(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetActivityCount(ctx))
}
