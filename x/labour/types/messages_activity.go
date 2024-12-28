package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateActivity{}

func NewMsgCreateActivity(creator string, worker string, beginWork uint64, finishWork uint64, workingTime int32, taskId uint64) *MsgCreateActivity {
	return &MsgCreateActivity{
		Creator:     creator,
		Worker:      worker,
		BeginWork:   beginWork,
		FinishWork:  finishWork,
		WorkingTime: workingTime,
		TaskId:      taskId,
	}
}

func (msg *MsgCreateActivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateActivity{}

func NewMsgUpdateActivity(creator string, id uint64, worker string, beginWork uint64, finishWork uint64, workingTime int32, taskId uint64) *MsgUpdateActivity {
	return &MsgUpdateActivity{
		Id:          id,
		Creator:     creator,
		Worker:      worker,
		BeginWork:   beginWork,
		FinishWork:  finishWork,
		WorkingTime: workingTime,
		TaskId:      taskId,
	}
}

func (msg *MsgUpdateActivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteActivity{}

func NewMsgDeleteActivity(creator string, id uint64) *MsgDeleteActivity {
	return &MsgDeleteActivity{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteActivity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
