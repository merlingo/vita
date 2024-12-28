package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/gainsharing module sentinel errors
var (
	ErrInvalidSigner                     = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                            = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrCoefficientParsing                = sdkerrors.Register(ModuleName, 1102, "Coefficients should be seperated by comma(,)")
	ErrMetricsParsing                    = sdkerrors.Register(ModuleName, 1103, "Metrics should be seperated by comma(,)")
	ErrMetricAndCoefficientNotCompatible = sdkerrors.Register(ModuleName, 1104, "Number of Coefficients and Metrics should be same)")
	ErrConvergeLimitParsing              = sdkerrors.Register(ModuleName, 1105, "Converge Limit should be in float type format ")
	ErrOutOfConvergeLimit                = sdkerrors.Register(ModuleName, 1106, "Converge Limit should be between 0 and 1 ")
	ErrSlopeParsing                      = sdkerrors.Register(ModuleName, 1107, "Slope should be in float type format ")
	ErrMechanismNotFound                 = sdkerrors.Register(ModuleName, 1108, "Mechanism with given ID is not found in the store ")
	ErrActivitiesNotFound                = sdkerrors.Register(ModuleName, 1109, "Activities with given task ID is not found in the store. Worker doesn't work on the task. ")
	ErrTaskNotFound                      = sdkerrors.Register(ModuleName, 1110, "Task with given task ID is not found in the store.")
	ErrRewardFunding                     = sdkerrors.Register(ModuleName, 1111, "Error raised when Coins send from module to Account as a Reward")
)
