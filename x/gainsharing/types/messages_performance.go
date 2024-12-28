package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePerformance{}

func NewMsgCreatePerformance(creator string, mid uint64, tid string, wager sdk.Coin, reward sdk.Coin, earner string) *MsgCreatePerformance {
	return &MsgCreatePerformance{
		Creator: creator,
		Mid:     mid,
		Tid:     tid,
		Wager:   wager,
		Reward:  reward,
		Earner:  earner,
	}
}

func (msg *MsgCreatePerformance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePerformance{}

func NewMsgUpdatePerformance(creator string, id uint64, mid uint64, tid string, wager sdk.Coin, reward sdk.Coin, earner string) *MsgUpdatePerformance {
	return &MsgUpdatePerformance{
		Id:      id,
		Creator: creator,
		Mid:     mid,
		Tid:     tid,
		Wager:   wager,
		Reward:  reward,
		Earner:  earner,
	}
}

func (msg *MsgUpdatePerformance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePerformance{}

func NewMsgDeletePerformance(creator string, id uint64) *MsgDeletePerformance {
	return &MsgDeletePerformance{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeletePerformance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
