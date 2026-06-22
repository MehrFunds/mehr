package keeper

import (
	"encoding/binary"

	storetypes "cosmossdk.io/store/types"
	proto "github.com/cosmos/gogoproto/proto"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mehrfunds/mehr/x/mehr/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
}

func NewKeeper(storeKey storetypes.StoreKey) Keeper {
	return Keeper{storeKey: storeKey}
}

func (k Keeper) kvStore(ctx sdk.Context) storetypes.KVStore {
	return ctx.KVStore(k.storeKey)
}

// ── Watch ────────────────────────────────────────────────────────────────────

func (k Keeper) nextWatchID(ctx sdk.Context) uint64 {
	bz := k.kvStore(ctx).Get([]byte(types.WatchCountKey))
	if bz == nil {
		return 1
	}
	return binary.BigEndian.Uint64(bz) + 1
}

func (k Keeper) setWatchCount(ctx sdk.Context, n uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, n)
	k.kvStore(ctx).Set([]byte(types.WatchCountKey), bz)
}

func watchKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.WatchKeyPrefix), bz...)
}

func ownerWatchKey(owner string, id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.OwnerWatchIndexPrefix+owner+"/"), bz...)
}

func (k Keeper) CreateWatch(ctx sdk.Context, owner, network, address, label string) (*types.Watch, error) {
	id := k.nextWatchID(ctx)
	w := &types.Watch{Id: id, Owner: owner, Network: network, Address: address, Label: label}
	bz, err := proto.Marshal(w)
	if err != nil {
		return nil, err
	}
	st := k.kvStore(ctx)
	st.Set(watchKey(id), bz)
	st.Set(ownerWatchKey(owner, id), []byte{1})
	k.setWatchCount(ctx, id)
	return w, nil
}

func (k Keeper) GetWatch(ctx sdk.Context, id uint64) (*types.Watch, bool) {
	bz := k.kvStore(ctx).Get(watchKey(id))
	if bz == nil {
		return nil, false
	}
	var w types.Watch
	if err := proto.Unmarshal(bz, &w); err != nil {
		return nil, false
	}
	return &w, true
}

func (k Keeper) DeleteWatch(ctx sdk.Context, id uint64, owner string) {
	st := k.kvStore(ctx)
	st.Delete(watchKey(id))
	st.Delete(ownerWatchKey(owner, id))
}

func (k Keeper) GetWatchesByOwner(ctx sdk.Context, owner string) []*types.Watch {
	st := k.kvStore(ctx)
	prefix := []byte(types.OwnerWatchIndexPrefix + owner + "/")
	iter := storetypes.KVStorePrefixIterator(st, prefix)
	defer iter.Close()

	var out []*types.Watch
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		if len(key) < 8 {
			continue
		}
		id := binary.BigEndian.Uint64(key[len(key)-8:])
		if w, ok := k.GetWatch(ctx, id); ok {
			out = append(out, w)
		}
	}
	return out
}

func (k Keeper) GetAllWatches(ctx sdk.Context) []*types.Watch {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.WatchKeyPrefix))
	defer iter.Close()

	var out []*types.Watch
	for ; iter.Valid(); iter.Next() {
		var w types.Watch
		if err := proto.Unmarshal(iter.Value(), &w); err == nil {
			out = append(out, &w)
		}
	}
	return out
}

// ── Webhook ──────────────────────────────────────────────────────────────────

func (k Keeper) nextWebhookID(ctx sdk.Context) uint64 {
	bz := k.kvStore(ctx).Get([]byte(types.WebhookCountKey))
	if bz == nil {
		return 1
	}
	return binary.BigEndian.Uint64(bz) + 1
}

func (k Keeper) setWebhookCount(ctx sdk.Context, n uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, n)
	k.kvStore(ctx).Set([]byte(types.WebhookCountKey), bz)
}

func webhookKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.WebhookKeyPrefix), bz...)
}

func ownerWebhookKey(owner string, id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.OwnerWebhookIndexPrefix+owner+"/"), bz...)
}

func (k Keeper) CreateWebhook(ctx sdk.Context, owner, url, secretHash string) (*types.Webhook, error) {
	id := k.nextWebhookID(ctx)
	wh := &types.Webhook{Id: id, Owner: owner, Url: url, SecretHash: secretHash}
	bz, err := proto.Marshal(wh)
	if err != nil {
		return nil, err
	}
	st := k.kvStore(ctx)
	st.Set(webhookKey(id), bz)
	st.Set(ownerWebhookKey(owner, id), []byte{1})
	k.setWebhookCount(ctx, id)
	return wh, nil
}

func (k Keeper) GetWebhook(ctx sdk.Context, id uint64) (*types.Webhook, bool) {
	bz := k.kvStore(ctx).Get(webhookKey(id))
	if bz == nil {
		return nil, false
	}
	var wh types.Webhook
	if err := proto.Unmarshal(bz, &wh); err != nil {
		return nil, false
	}
	return &wh, true
}

func (k Keeper) DeleteWebhook(ctx sdk.Context, id uint64, owner string) {
	st := k.kvStore(ctx)
	st.Delete(webhookKey(id))
	st.Delete(ownerWebhookKey(owner, id))
}

func (k Keeper) GetWebhooksByOwner(ctx sdk.Context, owner string) []*types.Webhook {
	st := k.kvStore(ctx)
	prefix := []byte(types.OwnerWebhookIndexPrefix + owner + "/")
	iter := storetypes.KVStorePrefixIterator(st, prefix)
	defer iter.Close()

	var out []*types.Webhook
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		if len(key) < 8 {
			continue
		}
		id := binary.BigEndian.Uint64(key[len(key)-8:])
		if wh, ok := k.GetWebhook(ctx, id); ok {
			out = append(out, wh)
		}
	}
	return out
}

func (k Keeper) GetAllWebhooks(ctx sdk.Context) []*types.Webhook {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.WebhookKeyPrefix))
	defer iter.Close()

	var out []*types.Webhook
	for ; iter.Valid(); iter.Next() {
		var wh types.Webhook
		if err := proto.Unmarshal(iter.Value(), &wh); err == nil {
			out = append(out, &wh)
		}
	}
	return out
}

// ── Event ─────────────────────────────────────────────────────────────────────

func (k Keeper) nextEventID(ctx sdk.Context) uint64 {
	bz := k.kvStore(ctx).Get([]byte(types.EventCountKey))
	if bz == nil {
		return 1
	}
	return binary.BigEndian.Uint64(bz) + 1
}

func (k Keeper) setEventCount(ctx sdk.Context, n uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, n)
	k.kvStore(ctx).Set([]byte(types.EventCountKey), bz)
}

func eventKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.EventKeyPrefix), bz...)
}

// dedupKey returns the key used to detect duplicate events.
// Format: ed/<network>/<txHash>/<logIndex_8bytes>
func dedupKey(network, txHash string, logIndex int32) []byte {
	bz := make([]byte, 4)
	binary.BigEndian.PutUint32(bz, uint32(logIndex))
	key := types.EventDedupPrefix + network + "/" + txHash + "/"
	return append([]byte(key), bz...)
}

// addrEventKey returns the address index key for an event.
// Format: ea/<address>/<id_8bytes>
func addrEventKey(address string, id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.EventAddrIndexPrefix+address+"/"), bz...)
}

// FindWatchForAddress returns the first active watch matching network+address, or nil.
func (k Keeper) FindWatchForAddress(ctx sdk.Context, network, address string) *types.Watch {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.WatchKeyPrefix))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var w types.Watch
		if err := proto.Unmarshal(iter.Value(), &w); err == nil {
			if w.Network == network && w.Address == address {
				return &w
			}
		}
	}
	return nil
}

func (k Keeper) IsDuplicateEvent(ctx sdk.Context, network, txHash string, logIndex int32) bool {
	return k.kvStore(ctx).Has(dedupKey(network, txHash, logIndex))
}

func (k Keeper) CreateEvent(ctx sdk.Context, e *types.Event) (*types.Event, error) {
	id := k.nextEventID(ctx)
	e.Id = id
	bz, err := proto.Marshal(e)
	if err != nil {
		return nil, err
	}
	st := k.kvStore(ctx)
	st.Set(eventKey(id), bz)
	st.Set(dedupKey(e.Network, e.TxHash, e.LogIndex), []byte{1})
	st.Set(addrEventKey(e.Address, id), []byte{1})
	k.setEventCount(ctx, id)
	return e, nil
}

func (k Keeper) GetEvent(ctx sdk.Context, id uint64) (*types.Event, bool) {
	bz := k.kvStore(ctx).Get(eventKey(id))
	if bz == nil {
		return nil, false
	}
	var e types.Event
	if err := proto.Unmarshal(bz, &e); err != nil {
		return nil, false
	}
	return &e, true
}

func (k Keeper) GetEventsByAddress(ctx sdk.Context, address string) []*types.Event {
	st := k.kvStore(ctx)
	prefix := []byte(types.EventAddrIndexPrefix + address + "/")
	iter := storetypes.KVStorePrefixIterator(st, prefix)
	defer iter.Close()

	var out []*types.Event
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		if len(key) < 8 {
			continue
		}
		id := binary.BigEndian.Uint64(key[len(key)-8:])
		if e, ok := k.GetEvent(ctx, id); ok {
			out = append(out, e)
		}
	}
	return out
}

func (k Keeper) GetAllEvents(ctx sdk.Context) []*types.Event {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.EventKeyPrefix))
	defer iter.Close()

	var out []*types.Event
	for ; iter.Valid(); iter.Next() {
		var e types.Event
		if err := proto.Unmarshal(iter.Value(), &e); err == nil {
			out = append(out, &e)
		}
	}
	return out
}

// ── Genesis helpers ──────────────────────────────────────────────────────────

func (k Keeper) ImportGenesis(ctx sdk.Context, gs *types.GenesisState) {
	for _, w := range gs.Watches {
		bz, _ := proto.Marshal(w)
		st := k.kvStore(ctx)
		st.Set(watchKey(w.Id), bz)
		st.Set(ownerWatchKey(w.Owner, w.Id), []byte{1})
		k.setWatchCount(ctx, w.Id)
	}
	for _, wh := range gs.Webhooks {
		bz, _ := proto.Marshal(wh)
		st := k.kvStore(ctx)
		st.Set(webhookKey(wh.Id), bz)
		st.Set(ownerWebhookKey(wh.Owner, wh.Id), []byte{1})
		k.setWebhookCount(ctx, wh.Id)
	}
	for _, e := range gs.Events {
		bz, _ := proto.Marshal(e)
		st := k.kvStore(ctx)
		st.Set(eventKey(e.Id), bz)
		st.Set(dedupKey(e.Network, e.TxHash, e.LogIndex), []byte{1})
		st.Set(addrEventKey(e.Address, e.Id), []byte{1})
		k.setEventCount(ctx, e.Id)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Watches:  k.GetAllWatches(ctx),
		Webhooks: k.GetAllWebhooks(ctx),
		Events:   k.GetAllEvents(ctx),
	}
}
