package keeper

import (
	"context"

	"vita/x/gainsharing/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PerformanceAll(ctx context.Context, req *types.QueryAllPerformanceRequest) (*types.QueryAllPerformanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var performances []types.Performance

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	performanceStore := prefix.NewStore(store, types.KeyPrefix(types.PerformanceKey))

	pageRes, err := query.Paginate(performanceStore, req.Pagination, func(key []byte, value []byte) error {
		var performance types.Performance
		if err := k.cdc.Unmarshal(value, &performance); err != nil {
			return err
		}

		performances = append(performances, performance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPerformanceResponse{Performance: performances, Pagination: pageRes}, nil
}

func (k Keeper) Performance(ctx context.Context, req *types.QueryGetPerformanceRequest) (*types.QueryGetPerformanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	performance, found := k.GetPerformance(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPerformanceResponse{Performance: performance}, nil
}
