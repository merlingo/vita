package types

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (task Task) Validate() (err error) {
	_, err = task.GetDaysUntilDeadline()
	if err != nil {
		return err
	}
	if task.State == 2 {
		_, err = task.GetAllDays()
		if err != nil {
			return err
		}
	}
	return
}
func (task Task) GetDaysUntilDeadline() (days int, err error) {

	if task.GetBeginTask() >= task.GetDeadline() {
		return -1, ErrDeadlineConfliction
	}
	diff := time.UnixMilli(int64(task.Deadline)).Sub(time.UnixMilli(int64(task.BeginTask)))

	days = int(diff.Hours() / 24)
	return
}
func (task Task) GetAllDays() (days int, err error) {

	if task.GetBeginTask() >= task.GetFinishTask() {
		return -1, ErrDeadlineConfliction
	}
	diff := time.UnixMilli(int64(task.Deadline)).Sub(time.UnixMilli(int64(task.BeginTask)))

	days = int(diff.Hours() / 24)
	return
}

func (task Task) GetAddress() (assigner sdk.AccAddress, err error) {
	assigner, errBlack := sdk.AccAddressFromBech32(task.Assigner)
	return assigner, sdkerrors.Wrapf(errBlack, ErrInvalidAssignerAddress.Error(), task.Assigner)
}
