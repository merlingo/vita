package labour

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"vita/testutil/sample"
	laboursimulation "vita/x/labour/simulation"
	"vita/x/labour/types"
)

// avoid unused import issue
var (
	_ = laboursimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTask = "op_weight_msg_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTask int = 100

	opWeightMsgUpdateTask = "op_weight_msg_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTask int = 100

	opWeightMsgDeleteTask = "op_weight_msg_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTask int = 100

	opWeightMsgCreateActivity = "op_weight_msg_activity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateActivity int = 100

	opWeightMsgUpdateActivity = "op_weight_msg_activity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateActivity int = 100

	opWeightMsgDeleteActivity = "op_weight_msg_activity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteActivity int = 100

	opWeightMsgBeginTask = "op_weight_msg_begin_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBeginTask int = 100

	opWeightMsgWork = "op_weight_msg_work"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWork int = 100

	opWeightMsgFinishTask = "op_weight_msg_finish_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinishTask int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	labourGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TaskList: []types.Task{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TaskCount: 2,
		ActivityList: []types.Activity{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ActivityCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&labourGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTask int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTask, &weightMsgCreateTask, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTask = defaultWeightMsgCreateTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTask,
		laboursimulation.SimulateMsgCreateTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTask int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTask, &weightMsgUpdateTask, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTask = defaultWeightMsgUpdateTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTask,
		laboursimulation.SimulateMsgUpdateTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTask int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTask, &weightMsgDeleteTask, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTask = defaultWeightMsgDeleteTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTask,
		laboursimulation.SimulateMsgDeleteTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateActivity int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateActivity, &weightMsgCreateActivity, nil,
		func(_ *rand.Rand) {
			weightMsgCreateActivity = defaultWeightMsgCreateActivity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateActivity,
		laboursimulation.SimulateMsgCreateActivity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateActivity int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateActivity, &weightMsgUpdateActivity, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateActivity = defaultWeightMsgUpdateActivity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateActivity,
		laboursimulation.SimulateMsgUpdateActivity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteActivity int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteActivity, &weightMsgDeleteActivity, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteActivity = defaultWeightMsgDeleteActivity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteActivity,
		laboursimulation.SimulateMsgDeleteActivity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBeginTask int
	simState.AppParams.GetOrGenerate(opWeightMsgBeginTask, &weightMsgBeginTask, nil,
		func(_ *rand.Rand) {
			weightMsgBeginTask = defaultWeightMsgBeginTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBeginTask,
		laboursimulation.SimulateMsgBeginTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWork int
	simState.AppParams.GetOrGenerate(opWeightMsgWork, &weightMsgWork, nil,
		func(_ *rand.Rand) {
			weightMsgWork = defaultWeightMsgWork
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWork,
		laboursimulation.SimulateMsgWork(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinishTask int
	simState.AppParams.GetOrGenerate(opWeightMsgFinishTask, &weightMsgFinishTask, nil,
		func(_ *rand.Rand) {
			weightMsgFinishTask = defaultWeightMsgFinishTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinishTask,
		laboursimulation.SimulateMsgFinishTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTask,
			defaultWeightMsgCreateTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgCreateTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTask,
			defaultWeightMsgUpdateTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgUpdateTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTask,
			defaultWeightMsgDeleteTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgDeleteTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateActivity,
			defaultWeightMsgCreateActivity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgCreateActivity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateActivity,
			defaultWeightMsgUpdateActivity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgUpdateActivity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteActivity,
			defaultWeightMsgDeleteActivity,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgDeleteActivity(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBeginTask,
			defaultWeightMsgBeginTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgBeginTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWork,
			defaultWeightMsgWork,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgWork(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWork,
			defaultWeightMsgWork,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgWork(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgFinishTask,
			defaultWeightMsgFinishTask,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				laboursimulation.SimulateMsgFinishTask(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
