package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TaskList:     []Task{},
		ActivityList: []Activity{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in task
	taskIdMap := make(map[uint64]bool)
	taskCount := gs.GetTaskCount()
	for _, elem := range gs.TaskList {
		if _, ok := taskIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for task")
		}
		if elem.Id >= taskCount {
			return fmt.Errorf("task id should be lower or equal than the last id")
		}
		taskIdMap[elem.Id] = true
	}
	// Check for duplicated ID in activity
	activityIdMap := make(map[uint64]bool)
	activityCount := gs.GetActivityCount()
	for _, elem := range gs.ActivityList {
		if _, ok := activityIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for activity")
		}
		if elem.Id >= activityCount {
			return fmt.Errorf("activity id should be lower or equal than the last id")
		}
		activityIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
