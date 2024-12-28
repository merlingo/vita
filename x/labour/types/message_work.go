package types

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWork{}

func NewMsgWork(creator string, worker string, taskid uint64, beginWork uint64, finishWork uint64, workingTime int32) *MsgWork {
	return &MsgWork{
		Creator:     creator,
		Worker:      worker,
		Taskid:      taskid,
		BeginWork:   beginWork,
		FinishWork:  finishWork,
		WorkingTime: workingTime,
	}
}

func (msg *MsgWork) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Worker)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if (len(strconv.Itoa(int(msg.BeginWork))) != 13) || (len(strconv.Itoa(int(msg.FinishWork))) != 13) {
		return errorsmod.Wrapf(ErrInvalidTimestampFormat, "invalid timestamp format of BeginTask or Deadline (%d, %d)", msg.BeginWork, msg.FinishWork)

	}
	return nil
}
