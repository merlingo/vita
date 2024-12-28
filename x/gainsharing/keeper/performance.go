package keeper

import (
	"context"
	"encoding/binary"

	"vita/x/gainsharing/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetPerformanceCount get the total number of performance
func (k Keeper) GetPerformanceCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PerformanceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPerformanceCount set the total number of performance
func (k Keeper) SetPerformanceCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PerformanceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPerformance appends a performance in the store with a new id and update the count
func (k Keeper) AppendPerformance(
	ctx context.Context,
	performance types.Performance,
) uint64 {
	// Create the performance
	count := k.GetPerformanceCount(ctx)

	// Set the ID of the appended value
	performance.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PerformanceKey))
	appendedValue := k.cdc.MustMarshal(&performance)
	store.Set(GetPerformanceIDBytes(performance.Id), appendedValue)

	// Update performance count
	k.SetPerformanceCount(ctx, count+1)

	return count
}

// SetPerformance set a specific performance in the store
func (k Keeper) SetPerformance(ctx context.Context, performance types.Performance) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PerformanceKey))
	b := k.cdc.MustMarshal(&performance)
	store.Set(GetPerformanceIDBytes(performance.Id), b)
}

// GetPerformance returns a performance from its id
func (k Keeper) GetPerformance(ctx context.Context, id uint64) (val types.Performance, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PerformanceKey))
	b := store.Get(GetPerformanceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePerformance removes a performance from the store
func (k Keeper) RemovePerformance(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PerformanceKey))
	store.Delete(GetPerformanceIDBytes(id))
}

// GetAllPerformance returns all performance
func (k Keeper) GetAllPerformance(ctx context.Context) (list []types.Performance) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PerformanceKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Performance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPerformanceIDBytes returns the byte representation of the ID
func GetPerformanceIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PerformanceKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
