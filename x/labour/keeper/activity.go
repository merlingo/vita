package keeper

import (
	"context"
	"encoding/binary"

	"vita/x/labour/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetActivityCount get the total number of activity
func (k Keeper) GetActivityCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ActivityCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetActivityCount set the total number of activity
func (k Keeper) SetActivityCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.ActivityCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendActivity appends a activity in the store with a new id and update the count
func (k Keeper) AppendActivity(
	ctx context.Context,
	activity types.Activity,
) uint64 {
	// Create the activity
	count := k.GetActivityCount(ctx)

	// Set the ID of the appended value
	activity.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActivityKey))
	appendedValue := k.cdc.MustMarshal(&activity)
	store.Set(GetActivityIDBytes(activity.Id), appendedValue)

	// Update activity count
	k.SetActivityCount(ctx, count+1)

	return count
}

// SetActivity set a specific activity in the store
func (k Keeper) SetActivity(ctx context.Context, activity types.Activity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActivityKey))
	b := k.cdc.MustMarshal(&activity)
	store.Set(GetActivityIDBytes(activity.Id), b)
}

// GetActivity returns a activity from its id
func (k Keeper) GetActivity(ctx context.Context, id uint64) (val types.Activity, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActivityKey))
	b := store.Get(GetActivityIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveActivity removes a activity from the store
func (k Keeper) RemoveActivity(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActivityKey))
	store.Delete(GetActivityIDBytes(id))
}

// GetAllActivity returns all activity
func (k Keeper) GetAllActivity(ctx context.Context) (list []types.Activity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActivityKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Activity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetActivityIDBytes returns the byte representation of the ID
func GetActivityIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.ActivityKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

func (k Keeper) GetTaskActivities(ctx context.Context, worker string, taskid uint64) (list []types.Activity) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.ActivityKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Activity
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.TaskId == taskid && val.Worker == worker {
			list = append(list, val)
		}
	}

	return
}
