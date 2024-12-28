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

func (k Keeper) MechanismAll(ctx context.Context, req *types.QueryAllMechanismRequest) (*types.QueryAllMechanismResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var mechanisms []types.Mechanism

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	mechanismStore := prefix.NewStore(store, types.KeyPrefix(types.MechanismKey))

	pageRes, err := query.Paginate(mechanismStore, req.Pagination, func(key []byte, value []byte) error {
		var mechanism types.Mechanism
		if err := k.cdc.Unmarshal(value, &mechanism); err != nil {
			return err
		}

		mechanisms = append(mechanisms, mechanism)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMechanismResponse{Mechanism: mechanisms, Pagination: pageRes}, nil
}

func (k Keeper) Mechanism(ctx context.Context, req *types.QueryGetMechanismRequest) (*types.QueryGetMechanismResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	mechanism, found := k.GetMechanism(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMechanismResponse{Mechanism: mechanism}, nil
}
