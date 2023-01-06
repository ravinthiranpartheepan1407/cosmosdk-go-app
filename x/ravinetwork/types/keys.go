package types

const (
	// ModuleName defines the module name
	ModuleName = "ravinetwork"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ravinetwork"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
