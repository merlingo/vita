package simulation

import (
	"math/rand"

	"vita/x/labour/keeper"
	"vita/x/labour/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgFinishTask(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgFinishTask{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the FinishTask simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "FinishTask simulation not implemented"), nil, nil
	}
}
