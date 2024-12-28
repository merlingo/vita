package labour

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"vita/x/labour/keeper"
	"vita/x/labour/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the task
	for _, elem := range genState.TaskList {
		k.SetTask(ctx, elem)
	}

	// Set task count
	k.SetTaskCount(ctx, genState.TaskCount)
	// Set all the activity
	for _, elem := range genState.ActivityList {
		k.SetActivity(ctx, elem)
	}

	// Set activity count
	k.SetActivityCount(ctx, genState.ActivityCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TaskList = k.GetAllTask(ctx)
	genesis.TaskCount = k.GetTaskCount(ctx)
	genesis.ActivityList = k.GetAllActivity(ctx)
	genesis.ActivityCount = k.GetActivityCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
