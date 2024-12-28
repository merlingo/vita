package keeper

import (
	"context"

	"vita/x/labour/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FinishTask(goCtx context.Context, msg *types.MsgFinishTask) (*types.MsgFinishTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// // 1- find the Task object 2. update the object 3. validate it 4. set the object 5. emit the event

	// 1- build Task object  2. validate it  3. store the object 4. emit the event
	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	task, found := k.Keeper.GetTask(ctx, msg.Taskid)
	if !found {
		return nil, types.ErrTaskNotFound
	}
	task.FinishTask = msg.GetFinishTask()
	task.State = 2

	err = task.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetTask(ctx, task)
	ctx.GasMeter().ConsumeGas(types.FinishTaskGas, "Complete The Task")

	return &types.MsgFinishTaskResponse{}, nil
}
