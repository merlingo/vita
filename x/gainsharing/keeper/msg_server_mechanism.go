package keeper

import (
	"context"
	"fmt"

	"vita/x/gainsharing/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMechanism(goCtx context.Context, msg *types.MsgCreateMechanism) (*types.MsgCreateMechanismResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var mechanism = types.Mechanism{
		Creator:       msg.Creator,
		Metrics:       msg.Metrics,
		Coefficients:  msg.Coefficients,
		ConvergeLimit: msg.ConvergeLimit,
		Slope:         msg.Slope,
	}

	id := k.AppendMechanism(
		ctx,
		mechanism,
	)

	return &types.MsgCreateMechanismResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMechanism(goCtx context.Context, msg *types.MsgUpdateMechanism) (*types.MsgUpdateMechanismResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var mechanism = types.Mechanism{
		Creator:       msg.Creator,
		Id:            msg.Id,
		Metrics:       msg.Metrics,
		Coefficients:  msg.Coefficients,
		ConvergeLimit: msg.ConvergeLimit,
		Slope:         msg.Slope,
	}

	// Checks that the element exists
	val, found := k.GetMechanism(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMechanism(ctx, mechanism)

	return &types.MsgUpdateMechanismResponse{}, nil
}

func (k msgServer) DeleteMechanism(goCtx context.Context, msg *types.MsgDeleteMechanism) (*types.MsgDeleteMechanismResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMechanism(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMechanism(ctx, msg.Id)

	return &types.MsgDeleteMechanismResponse{}, nil
}
