package keeper

import (
	"context"
	"fmt"

	"vita/x/labour/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateActivity(goCtx context.Context, msg *types.MsgCreateActivity) (*types.MsgCreateActivityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var activity = types.Activity{
		Creator:     msg.Creator,
		Worker:      msg.Worker,
		BeginWork:   msg.BeginWork,
		FinishWork:  msg.FinishWork,
		WorkingTime: msg.WorkingTime,
		TaskId:      msg.TaskId,
	}

	id := k.AppendActivity(
		ctx,
		activity,
	)

	return &types.MsgCreateActivityResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateActivity(goCtx context.Context, msg *types.MsgUpdateActivity) (*types.MsgUpdateActivityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var activity = types.Activity{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Worker:      msg.Worker,
		BeginWork:   msg.BeginWork,
		FinishWork:  msg.FinishWork,
		WorkingTime: msg.WorkingTime,
		TaskId:      msg.TaskId,
	}

	// Checks that the element exists
	val, found := k.GetActivity(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetActivity(ctx, activity)

	return &types.MsgUpdateActivityResponse{}, nil
}

func (k msgServer) DeleteActivity(goCtx context.Context, msg *types.MsgDeleteActivity) (*types.MsgDeleteActivityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetActivity(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveActivity(ctx, msg.Id)

	return &types.MsgDeleteActivityResponse{}, nil
}
