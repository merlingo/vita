package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTask{}

func NewMsgCreateTask(creator string, taskId uint64, assigner string, state int32, beginTask uint64, deadline uint64, finishTask uint64, wager string) *MsgCreateTask {
	return &MsgCreateTask{
		Creator:    creator,
		TaskId:     taskId,
		Assigner:   assigner,
		State:      state,
		BeginTask:  beginTask,
		Deadline:   deadline,
		FinishTask: finishTask,
		Wager:      wager,
	}
}

func (msg *MsgCreateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTask{}

func NewMsgUpdateTask(creator string, id uint64, taskId uint64, assigner string, state int32, beginTask uint64, deadline uint64, finishTask uint64, wager string) *MsgUpdateTask {
	return &MsgUpdateTask{
		Id:         id,
		Creator:    creator,
		TaskId:     taskId,
		Assigner:   assigner,
		State:      state,
		BeginTask:  beginTask,
		Deadline:   deadline,
		FinishTask: finishTask,
		Wager:      wager,
	}
}

func (msg *MsgUpdateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTask{}

func NewMsgDeleteTask(creator string, id uint64) *MsgDeleteTask {
	return &MsgDeleteTask{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
