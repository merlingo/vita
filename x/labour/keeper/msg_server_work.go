package keeper

import (
	"context"

	"vita/x/labour/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Work(goCtx context.Context, msg *types.MsgWork) (*types.MsgWorkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	activity := types.Activity{
		TaskId:      msg.GetTaskid(),
		Worker:      msg.GetWorker(),
		WorkingTime: msg.GetWorkingTime(),
		BeginWork:   msg.GetBeginWork(),
		FinishWork:  msg.GetFinishWork(),
	}
	//err = task.Validate()
	//if err != nil {
	//	return nil, err
	//}
	id := k.Keeper.AppendActivity(ctx, activity)
	ctx.GasMeter().ConsumeGas(types.WorkGas, "Create New Activity")

	return &types.MsgWorkResponse{Id: id}, nil
}
