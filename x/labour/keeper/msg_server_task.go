package keeper

import (
	"context"
	"fmt"

	"vita/x/labour/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var task = types.Task{
		Creator:    msg.Creator,
		TaskId:     msg.TaskId,
		Assigner:   msg.Assigner,
		State:      msg.State,
		BeginTask:  msg.BeginTask,
		Deadline:   msg.Deadline,
		FinishTask: msg.FinishTask,
		Wager:      msg.Wager,
	}

	id := k.AppendTask(
		ctx,
		task,
	)

	return &types.MsgCreateTaskResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTask(goCtx context.Context, msg *types.MsgUpdateTask) (*types.MsgUpdateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var task = types.Task{
		Creator:    msg.Creator,
		Id:         msg.Id,
		TaskId:     msg.TaskId,
		Assigner:   msg.Assigner,
		State:      msg.State,
		BeginTask:  msg.BeginTask,
		Deadline:   msg.Deadline,
		FinishTask: msg.FinishTask,
		Wager:      msg.Wager,
	}

	// Checks that the element exists
	val, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTask(ctx, task)

	return &types.MsgUpdateTaskResponse{}, nil
}

func (k msgServer) DeleteTask(goCtx context.Context, msg *types.MsgDeleteTask) (*types.MsgDeleteTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTask(ctx, msg.Id)

	return &types.MsgDeleteTaskResponse{}, nil
}
