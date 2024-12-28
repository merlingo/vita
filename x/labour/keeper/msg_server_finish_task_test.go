package keeper_test

import (
	"testing"

	"vita/x/labour/types"

	"github.com/stretchr/testify/require"
)

func TestFinishTask(t *testing.T) {
	msgServer, _, context := setupMsgServerBeginTask(t)
	createResponse, err := msgServer.BeginTask(context, &types.MsgBeginTask{
		Creator:   alice,
		Taskid:    taskId,
		Assigner:  alice,
		BeginTask: 1729687287000,
		Deadline:  1732365735000,
	})

	finishResponse, err := msgServer.FinishTask(context, &types.MsgFinishTask{
		Creator:    alice,
		Taskid:     createResponse.Id,
		FinishTask: 1732365735000,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgFinishTaskResponse{
		// TODO: update with a proper value when updated
	}, *finishResponse)

}
