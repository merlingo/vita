package types

const (
	// ModuleName defines the module name
	ModuleName = "vita"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_vita"
)

var (
	ParamsKey = []byte("p_vita")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
