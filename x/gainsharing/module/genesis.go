package gainsharing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"vita/x/gainsharing/keeper"
	"vita/x/gainsharing/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the mechanism
	for _, elem := range genState.MechanismList {
		k.SetMechanism(ctx, elem)
	}

	// Set mechanism count
	k.SetMechanismCount(ctx, genState.MechanismCount)
	// Set all the performance
	for _, elem := range genState.PerformanceList {
		k.SetPerformance(ctx, elem)
	}

	// Set performance count
	k.SetPerformanceCount(ctx, genState.PerformanceCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MechanismList = k.GetAllMechanism(ctx)
	genesis.MechanismCount = k.GetMechanismCount(ctx)
	genesis.PerformanceList = k.GetAllPerformance(ctx)
	genesis.PerformanceCount = k.GetPerformanceCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
