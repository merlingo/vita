package types

import (
	"context"
	"vita/x/labour/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type LabourKeeper interface {
// TODO Add methods imported from labour should be defined here
// }
type LabourKeeper interface {
	// TODO Add methods imported from labour should be defined here
	GetTaskActivities(ctx context.Context, worker string, taskid uint64) (list []types.Activity)
	GetTask(ctx context.Context, id uint64) (val types.Task, found bool)
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
