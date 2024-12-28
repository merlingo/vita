package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCalculatePerformance{}

func NewMsgCalculatePerformance(creator string, mid uint64, wager sdk.Coin, earner string, taskid uint64) *MsgCalculatePerformance {
	return &MsgCalculatePerformance{
		Creator: creator,
		Mid:     mid,
		Wager:   wager,
		Earner:  earner,
		Taskid:  taskid,
	}
}

func (msg *MsgCalculatePerformance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
