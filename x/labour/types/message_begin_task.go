package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBeginTask{}

func NewMsgBeginTask(creator string, taskid uint64, assigner string, beginTask uint64, deadline uint64) *MsgBeginTask {
	return &MsgBeginTask{
		Creator:   creator,
		Taskid:    taskid,
		Assigner:  assigner,
		BeginTask: beginTask,
		Deadline:  deadline,
	}
}
func (msg *MsgBeginTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBeginTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Assigner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid assigner address (%s)", err)
	}
	//check dates are correct timestamp format
	if (len(strconv.Itoa(int(msg.BeginTask))) != 13) || (len(strconv.Itoa(int(msg.Deadline))) != 13) {
		return errorsmod.Wrapf(ErrInvalidTimestampFormat, "invalid timestamp format of BeginTask or Deadline (%d, %d)", msg.BeginTask, msg.Deadline)

	}
	return nil
}
