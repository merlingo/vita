package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "vita/testutil/keeper"
	"vita/testutil/nullify"
	"vita/x/gainsharing/types"
)

func TestMechanismQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	msgs := createNMechanism(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMechanismRequest
		response *types.QueryGetMechanismResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMechanismRequest{Id: msgs[0].Id},
			response: &types.QueryGetMechanismResponse{Mechanism: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetMechanismRequest{Id: msgs[1].Id},
			response: &types.QueryGetMechanismResponse{Mechanism: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetMechanismRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Mechanism(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMechanismQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.GainsharingKeeper(t)
	msgs := createNMechanism(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMechanismRequest {
		return &types.QueryAllMechanismRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MechanismAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mechanism), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Mechanism),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MechanismAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Mechanism), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Mechanism),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MechanismAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Mechanism),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MechanismAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
