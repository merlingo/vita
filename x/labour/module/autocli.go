package labour

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "vita/api/vita/labour"
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
					RpcMethod: "TaskAll",
					Use:       "list-task",
					Short:     "List all task",
				},
				{
					RpcMethod:      "Task",
					Use:            "show-task [id]",
					Short:          "Shows a task by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ActivityAll",
					Use:       "list-activity",
					Short:     "List all activity",
				},
				{
					RpcMethod:      "Activity",
					Use:            "show-activity [id]",
					Short:          "Shows a activity by id",
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
					RpcMethod:      "CreateTask",
					Use:            "create-task [taskId] [assigner] [state] [beginTask] [deadline] [finishTask] [wager]",
					Short:          "Create task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "taskId"}, {ProtoField: "assigner"}, {ProtoField: "state"}, {ProtoField: "beginTask"}, {ProtoField: "deadline"}, {ProtoField: "finishTask"}, {ProtoField: "wager"}},
				},
				{
					RpcMethod:      "UpdateTask",
					Use:            "update-task [id] [taskId] [assigner] [state] [beginTask] [deadline] [finishTask] [wager]",
					Short:          "Update task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "taskId"}, {ProtoField: "assigner"}, {ProtoField: "state"}, {ProtoField: "beginTask"}, {ProtoField: "deadline"}, {ProtoField: "finishTask"}, {ProtoField: "wager"}},
				},
				{
					RpcMethod:      "DeleteTask",
					Use:            "delete-task [id]",
					Short:          "Delete task",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "CreateActivity",
					Use:            "create-activity [worker] [beginWork] [finishWork] [workingTime] [taskId]",
					Short:          "Create activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "worker"}, {ProtoField: "beginWork"}, {ProtoField: "finishWork"}, {ProtoField: "workingTime"}, {ProtoField: "taskId"}},
				},
				{
					RpcMethod:      "UpdateActivity",
					Use:            "update-activity [id] [worker] [beginWork] [finishWork] [workingTime] [taskId]",
					Short:          "Update activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "worker"}, {ProtoField: "beginWork"}, {ProtoField: "finishWork"}, {ProtoField: "workingTime"}, {ProtoField: "taskId"}},
				},
				{
					RpcMethod:      "DeleteActivity",
					Use:            "delete-activity [id]",
					Short:          "Delete activity",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "BeginTask",
					Use:            "begin-task [taskid] [assigner] [begin-task] [deadline] [wager]",
					Short:          "Send a begin-task tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "taskid"}, {ProtoField: "assigner"}, {ProtoField: "beginTask"}, {ProtoField: "deadline"}, {ProtoField: "wager"}},
				},
				{
					RpcMethod:      "Work",
					Use:            "work [worker] [taskid] [begin-work] [finish-work] [working-time]",
					Short:          "Send a work tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "worker"}, {ProtoField: "taskid"}, {ProtoField: "beginWork"}, {ProtoField: "finishWork"}, {ProtoField: "workingTime"}},
				},
				{
					RpcMethod:      "Work",
					Use:            "work [worker] [taskid] [begin-work] [finish-work] [working-time]",
					Short:          "Send a work tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "worker"}, {ProtoField: "taskid"}, {ProtoField: "beginWork"}, {ProtoField: "finishWork"}, {ProtoField: "workingTime"}},
				},
				{
					RpcMethod:      "FinishTask",
					Use:            "finish-task [taskid] [finish-task]",
					Short:          "Send a finish-task tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "taskid"}, {ProtoField: "finishTask"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
