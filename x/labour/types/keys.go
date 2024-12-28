package types

const (
	// ModuleName defines the module name
	ModuleName = "labour"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_labour"
)

var (
	ParamsKey = []byte("p_labour")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TaskKey      = "Task/value/"
	TaskCountKey = "Task/count/"
)
const (
	BeginTaskGas  = 15000
	WorkGas       = 1000
	FinishTaskGas = 2000
)
const (
	ActivityKey      = "Activity/value/"
	ActivityCountKey = "Activity/count/"
)

const (
	BeginningTaskEventType = "move-played"
	CreatedTaskId          = "id"
)
