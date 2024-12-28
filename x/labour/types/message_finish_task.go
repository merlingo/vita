package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgFinishTask{}

func NewMsgFinishTask(creator string, taskid uint64, finishTask uint64) *MsgFinishTask {
	return &MsgFinishTask{
		Creator:    creator,
		Taskid:     taskid,
		FinishTask: finishTask,
	}
}

func (msg *MsgFinishTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(strconv.Itoa(int(msg.FinishTask))) != 13 {
		return errorsmod.Wrapf(ErrInvalidTimestampFormat, "invalid timestamp format of BeginTask or Deadline (%d)", msg.FinishTask)

	}
	return nil
}
