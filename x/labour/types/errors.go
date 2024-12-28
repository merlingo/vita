package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/labour module sentinel errors
var (
	ErrInvalidSigner          = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                 = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidTimestampFormat = sdkerrors.Register(ModuleName, 1102, "timestamp format is wrong. It should be unix.milli")
	ErrDeadlineConfliction    = sdkerrors.Register(ModuleName, 1103, "Deadline cannot be same with or earlier than beginning: %d")
	ErrTaskNotFound           = sdkerrors.Register(ModuleName, 1104, "There is no task with given ID")
	ErrInvalidAssignerAddress = sdkerrors.Register(ModuleName, 1105, "Invalid assigner address %s")
)
