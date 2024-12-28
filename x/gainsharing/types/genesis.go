package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MechanismList:   []Mechanism{},
		PerformanceList: []Performance{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in mechanism
	mechanismIdMap := make(map[uint64]bool)
	mechanismCount := gs.GetMechanismCount()
	for _, elem := range gs.MechanismList {
		if _, ok := mechanismIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for mechanism")
		}
		if elem.Id >= mechanismCount {
			return fmt.Errorf("mechanism id should be lower or equal than the last id")
		}
		mechanismIdMap[elem.Id] = true
	}
	// Check for duplicated ID in performance
	performanceIdMap := make(map[uint64]bool)
	performanceCount := gs.GetPerformanceCount()
	for _, elem := range gs.PerformanceList {
		if _, ok := performanceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for performance")
		}
		if elem.Id >= performanceCount {
			return fmt.Errorf("performance id should be lower or equal than the last id")
		}
		performanceIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
