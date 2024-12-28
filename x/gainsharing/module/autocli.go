package gainsharing

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "vita/api/vita/gainsharing"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "MechanismAll",
					Use:       "list-mechanism",
					Short:     "List all mechanism",
				},
				{
					RpcMethod:      "Mechanism",
					Use:            "show-mechanism [id]",
					Short:          "Shows a mechanism by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "PerformanceAll",
					Use:       "list-performance",
					Short:     "List all performance",
				},
				{
					RpcMethod:      "Performance",
					Use:            "show-performance [id]",
					Short:          "Shows a performance by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateMechanism",
					Use:            "create-mechanism [metrics] [coefficients] [convergeLimit] [slope]",
					Short:          "Create mechanism",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "metrics"}, {ProtoField: "coefficients"}, {ProtoField: "convergeLimit"}, {ProtoField: "slope"}},
				},
				{
					RpcMethod:      "UpdateMechanism",
					Use:            "update-mechanism [id] [metrics] [coefficients] [convergeLimit] [slope]",
					Short:          "Update mechanism",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "metrics"}, {ProtoField: "coefficients"}, {ProtoField: "convergeLimit"}, {ProtoField: "slope"}},
				},
				{
					RpcMethod:      "DeleteMechanism",
					Use:            "delete-mechanism [id]",
					Short:          "Delete mechanism",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreatePerformance",
					Use:            "create-performance [mid] [tid] [wager] [reward] [earner]",
					Short:          "Create performance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "mid"}, {ProtoField: "tid"}, {ProtoField: "wager"}, {ProtoField: "reward"}, {ProtoField: "earner"}},
				},
				{
					RpcMethod:      "UpdatePerformance",
					Use:            "update-performance [id] [mid] [tid] [wager] [reward] [earner]",
					Short:          "Update performance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "mid"}, {ProtoField: "tid"}, {ProtoField: "wager"}, {ProtoField: "reward"}, {ProtoField: "earner"}},
				},
				{
					RpcMethod:      "DeletePerformance",
					Use:            "delete-performance [id]",
					Short:          "Delete performance",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CalculatePerformance",
					Use:            "calculate-performance [mid] [wager] [earner] [taskid]",
					Short:          "Send a calculate-performance tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "mid"}, {ProtoField: "wager"}, {ProtoField: "earner"}, {ProtoField: "taskid"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
