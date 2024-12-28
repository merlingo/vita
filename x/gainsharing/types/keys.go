package types

const (
	// ModuleName defines the module name
	ModuleName = "gainsharing"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gainsharing"
)

var (
	ParamsKey = []byte("p_gainsharing")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	MechanismKey      = "Mechanism/value/"
	MechanismCountKey = "Mechanism/count/"
)

const (
	PerformanceKey      = "Performance/value/"
	PerformanceCountKey = "Performance/count/"
)
const (
	SetNewMechanism = 25000
	PerformanceGas  = 1000
	FinishTaskGas   = 2000
)
