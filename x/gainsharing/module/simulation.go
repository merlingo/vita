package gainsharing

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"vita/testutil/sample"
	gainsharingsimulation "vita/x/gainsharing/simulation"
	"vita/x/gainsharing/types"
)

// avoid unused import issue
var (
	_ = gainsharingsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateMechanism = "op_weight_msg_mechanism"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMechanism int = 100

	opWeightMsgUpdateMechanism = "op_weight_msg_mechanism"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMechanism int = 100

	opWeightMsgDeleteMechanism = "op_weight_msg_mechanism"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMechanism int = 100

	opWeightMsgCreatePerformance = "op_weight_msg_performance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePerformance int = 100

	opWeightMsgUpdatePerformance = "op_weight_msg_performance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePerformance int = 100

	opWeightMsgDeletePerformance = "op_weight_msg_performance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePerformance int = 100

	opWeightMsgCalculatePerformance = "op_weight_msg_calculate_performance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCalculatePerformance int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gainsharingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MechanismList: []types.Mechanism{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MechanismCount: 2,
		PerformanceList: []types.Performance{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PerformanceCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gainsharingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMechanism int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMechanism, &weightMsgCreateMechanism, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMechanism = defaultWeightMsgCreateMechanism
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMechanism,
		gainsharingsimulation.SimulateMsgCreateMechanism(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMechanism int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMechanism, &weightMsgUpdateMechanism, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMechanism = defaultWeightMsgUpdateMechanism
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMechanism,
		gainsharingsimulation.SimulateMsgUpdateMechanism(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMechanism int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMechanism, &weightMsgDeleteMechanism, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMechanism = defaultWeightMsgDeleteMechanism
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMechanism,
		gainsharingsimulation.SimulateMsgDeleteMechanism(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePerformance int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePerformance, &weightMsgCreatePerformance, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePerformance = defaultWeightMsgCreatePerformance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePerformance,
		gainsharingsimulation.SimulateMsgCreatePerformance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePerformance int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePerformance, &weightMsgUpdatePerformance, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePerformance = defaultWeightMsgUpdatePerformance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePerformance,
		gainsharingsimulation.SimulateMsgUpdatePerformance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePerformance int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePerformance, &weightMsgDeletePerformance, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePerformance = defaultWeightMsgDeletePerformance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePerformance,
		gainsharingsimulation.SimulateMsgDeletePerformance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCalculatePerformance int
	simState.AppParams.GetOrGenerate(opWeightMsgCalculatePerformance, &weightMsgCalculatePerformance, nil,
		func(_ *rand.Rand) {
			weightMsgCalculatePerformance = defaultWeightMsgCalculatePerformance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCalculatePerformance,
		gainsharingsimulation.SimulateMsgCalculatePerformance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMechanism,
			defaultWeightMsgCreateMechanism,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgCreateMechanism(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMechanism,
			defaultWeightMsgUpdateMechanism,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgUpdateMechanism(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMechanism,
			defaultWeightMsgDeleteMechanism,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgDeleteMechanism(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePerformance,
			defaultWeightMsgCreatePerformance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgCreatePerformance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePerformance,
			defaultWeightMsgUpdatePerformance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgUpdatePerformance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePerformance,
			defaultWeightMsgDeletePerformance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgDeletePerformance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCalculatePerformance,
			defaultWeightMsgCalculatePerformance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				gainsharingsimulation.SimulateMsgCalculatePerformance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
