package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMechanism{}

func NewMsgCreateMechanism(creator string, metrics string, coefficients string, convergeLimit string, slope string) *MsgCreateMechanism {
	return &MsgCreateMechanism{
		Creator:       creator,
		Metrics:       metrics,
		Coefficients:  coefficients,
		ConvergeLimit: convergeLimit,
		Slope:         slope,
	}
}

func (msg *MsgCreateMechanism) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMechanism{}

func NewMsgUpdateMechanism(creator string, id uint64, metrics string, coefficients string, convergeLimit string, slope string) *MsgUpdateMechanism {
	return &MsgUpdateMechanism{
		Id:            id,
		Creator:       creator,
		Metrics:       metrics,
		Coefficients:  coefficients,
		ConvergeLimit: convergeLimit,
		Slope:         slope,
	}
}

func (msg *MsgUpdateMechanism) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMechanism{}

func NewMsgDeleteMechanism(creator string, id uint64) *MsgDeleteMechanism {
	return &MsgDeleteMechanism{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteMechanism) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
