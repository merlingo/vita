package keeper

import (
	"context"
	"encoding/binary"

	"vita/x/gainsharing/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetMechanismCount get the total number of mechanism
func (k Keeper) GetMechanismCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MechanismCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMechanismCount set the total number of mechanism
func (k Keeper) SetMechanismCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MechanismCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMechanism appends a mechanism in the store with a new id and update the count
func (k Keeper) AppendMechanism(
	ctx context.Context,
	mechanism types.Mechanism,
) uint64 {
	// Create the mechanism
	count := k.GetMechanismCount(ctx)

	// Set the ID of the appended value
	mechanism.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MechanismKey))
	appendedValue := k.cdc.MustMarshal(&mechanism)
	store.Set(GetMechanismIDBytes(mechanism.Id), appendedValue)

	// Update mechanism count
	k.SetMechanismCount(ctx, count+1)

	return count
}

// SetMechanism set a specific mechanism in the store
func (k Keeper) SetMechanism(ctx context.Context, mechanism types.Mechanism) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MechanismKey))
	b := k.cdc.MustMarshal(&mechanism)
	store.Set(GetMechanismIDBytes(mechanism.Id), b)
}

// GetMechanism returns a mechanism from its id
func (k Keeper) GetMechanism(ctx context.Context, id uint64) (val types.Mechanism, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MechanismKey))
	b := store.Get(GetMechanismIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMechanism removes a mechanism from the store
func (k Keeper) RemoveMechanism(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MechanismKey))
	store.Delete(GetMechanismIDBytes(id))
}

// GetAllMechanism returns all mechanism
func (k Keeper) GetAllMechanism(ctx context.Context) (list []types.Mechanism) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MechanismKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mechanism
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMechanismIDBytes returns the byte representation of the ID
func GetMechanismIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.MechanismKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
