package keeper

import (
	"context"
	"fmt"

	"vita/x/gainsharing/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePerformance(goCtx context.Context, msg *types.MsgCreatePerformance) (*types.MsgCreatePerformanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var performance = types.Performance{
		Creator: msg.Creator,
		Mid:     msg.Mid,
		Tid:     msg.Tid,
		Wager:   msg.Wager,
		Reward:  msg.Reward,
		Earner:  msg.Earner,
	}

	id := k.AppendPerformance(
		ctx,
		performance,
	)

	return &types.MsgCreatePerformanceResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdatePerformance(goCtx context.Context, msg *types.MsgUpdatePerformance) (*types.MsgUpdatePerformanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var performance = types.Performance{
		Creator: msg.Creator,
		Id:      msg.Id,
		Mid:     msg.Mid,
		Tid:     msg.Tid,
		Wager:   msg.Wager,
		Reward:  msg.Reward,
		Earner:  msg.Earner,
	}

	// Checks that the element exists
	val, found := k.GetPerformance(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetPerformance(ctx, performance)

	return &types.MsgUpdatePerformanceResponse{}, nil
}

func (k msgServer) DeletePerformance(goCtx context.Context, msg *types.MsgDeletePerformance) (*types.MsgDeletePerformanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetPerformance(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePerformance(ctx, msg.Id)

	return &types.MsgDeletePerformanceResponse{}, nil
}
