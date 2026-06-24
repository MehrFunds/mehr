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

// ── Feed ──────────────────────────────────────────────────────────────────────

func feedKey(id string) []byte {
	return append([]byte(types.FeedKeyPrefix), []byte(id)...)
}

func (k Keeper) RegisterFeed(ctx sdk.Context, f *types.Feed) (*types.Feed, error) {
	bz, err := proto.Marshal(f)
	if err != nil {
		return nil, err
	}
	k.kvStore(ctx).Set(feedKey(f.Id), bz)
	return f, nil
}

func (k Keeper) GetFeed(ctx sdk.Context, id string) (*types.Feed, bool) {
	bz := k.kvStore(ctx).Get(feedKey(id))
	if bz == nil {
		return nil, false
	}
	var f types.Feed
	if err := proto.Unmarshal(bz, &f); err != nil {
		return nil, false
	}
	return &f, true
}

func (k Keeper) FeedExists(ctx sdk.Context, id string) bool {
	return k.kvStore(ctx).Has(feedKey(id))
}

func (k Keeper) GetAllFeeds(ctx sdk.Context) []*types.Feed {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.FeedKeyPrefix))
	defer iter.Close()

	var out []*types.Feed
	for ; iter.Valid(); iter.Next() {
		var f types.Feed
		if err := proto.Unmarshal(iter.Value(), &f); err == nil {
			out = append(out, &f)
		}
	}
	return out
}

// ── Fact ──────────────────────────────────────────────────────────────────────

func (k Keeper) nextFactID(ctx sdk.Context) uint64 {
	bz := k.kvStore(ctx).Get([]byte(types.FactCountKey))
	if bz == nil {
		return 1
	}
	return binary.BigEndian.Uint64(bz) + 1
}

func (k Keeper) setFactCount(ctx sdk.Context, n uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, n)
	k.kvStore(ctx).Set([]byte(types.FactCountKey), bz)
}

func factKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.FactKeyPrefix), bz...)
}

func factDedupKey(feedID, dedupKey string) []byte {
	return append([]byte(types.FactDedupPrefix+feedID+"/"), []byte(dedupKey)...)
}

func factFeedIndexKey(feedID string, id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.FactFeedIndexPrefix+feedID+"/"), bz...)
}

func (k Keeper) IsDuplicateFact(ctx sdk.Context, feedID, dedupKeyStr string) bool {
	return k.kvStore(ctx).Has(factDedupKey(feedID, dedupKeyStr))
}

func (k Keeper) CreateFact(ctx sdk.Context, f *types.Fact) (*types.Fact, error) {
	id := k.nextFactID(ctx)
	f.Id = id
	bz, err := proto.Marshal(f)
	if err != nil {
		return nil, err
	}
	st := k.kvStore(ctx)
	st.Set(factKey(id), bz)
	st.Set(factDedupKey(f.FeedId, f.DeduplicationKey), []byte{1})
	st.Set(factFeedIndexKey(f.FeedId, id), []byte{1})
	k.setFactCount(ctx, id)
	return f, nil
}

func (k Keeper) GetFact(ctx sdk.Context, id uint64) (*types.Fact, bool) {
	bz := k.kvStore(ctx).Get(factKey(id))
	if bz == nil {
		return nil, false
	}
	var f types.Fact
	if err := proto.Unmarshal(bz, &f); err != nil {
		return nil, false
	}
	return &f, true
}

func (k Keeper) GetFactsByFeed(ctx sdk.Context, feedID string) []*types.Fact {
	st := k.kvStore(ctx)
	prefix := []byte(types.FactFeedIndexPrefix + feedID + "/")
	iter := storetypes.KVStorePrefixIterator(st, prefix)
	defer iter.Close()

	var out []*types.Fact
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		if len(key) < 8 {
			continue
		}
		id := binary.BigEndian.Uint64(key[len(key)-8:])
		if f, ok := k.GetFact(ctx, id); ok {
			out = append(out, f)
		}
	}
	return out
}

func (k Keeper) GetAllFacts(ctx sdk.Context) []*types.Fact {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.FactKeyPrefix))
	defer iter.Close()

	var out []*types.Fact
	for ; iter.Valid(); iter.Next() {
		var f types.Fact
		if err := proto.Unmarshal(iter.Value(), &f); err == nil {
			out = append(out, &f)
		}
	}
	return out
}

// ── Subscription ──────────────────────────────────────────────────────────────

func (k Keeper) nextSubscriptionID(ctx sdk.Context) uint64 {
	bz := k.kvStore(ctx).Get([]byte(types.SubscriptionCountKey))
	if bz == nil {
		return 1
	}
	return binary.BigEndian.Uint64(bz) + 1
}

func (k Keeper) setSubscriptionCount(ctx sdk.Context, n uint64) {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, n)
	k.kvStore(ctx).Set([]byte(types.SubscriptionCountKey), bz)
}

func subscriptionKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.SubscriptionKeyPrefix), bz...)
}

func subscriptionOwnerKey(owner string, id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append([]byte(types.SubscriptionOwnerIndexPrefix+owner+"/"), bz...)
}

func (k Keeper) CreateSubscription(ctx sdk.Context, s *types.Subscription) (*types.Subscription, error) {
	id := k.nextSubscriptionID(ctx)
	s.Id = id
	bz, err := proto.Marshal(s)
	if err != nil {
		return nil, err
	}
	st := k.kvStore(ctx)
	st.Set(subscriptionKey(id), bz)
	st.Set(subscriptionOwnerKey(s.Owner, id), []byte{1})
	k.setSubscriptionCount(ctx, id)
	return s, nil
}

func (k Keeper) GetSubscription(ctx sdk.Context, id uint64) (*types.Subscription, bool) {
	bz := k.kvStore(ctx).Get(subscriptionKey(id))
	if bz == nil {
		return nil, false
	}
	var s types.Subscription
	if err := proto.Unmarshal(bz, &s); err != nil {
		return nil, false
	}
	return &s, true
}

func (k Keeper) DeleteSubscription(ctx sdk.Context, id uint64, owner string) {
	st := k.kvStore(ctx)
	st.Delete(subscriptionKey(id))
	st.Delete(subscriptionOwnerKey(owner, id))
}

func (k Keeper) GetSubscriptionsByOwner(ctx sdk.Context, owner string) []*types.Subscription {
	st := k.kvStore(ctx)
	prefix := []byte(types.SubscriptionOwnerIndexPrefix + owner + "/")
	iter := storetypes.KVStorePrefixIterator(st, prefix)
	defer iter.Close()

	var out []*types.Subscription
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		if len(key) < 8 {
			continue
		}
		id := binary.BigEndian.Uint64(key[len(key)-8:])
		if s, ok := k.GetSubscription(ctx, id); ok {
			out = append(out, s)
		}
	}
	return out
}

func (k Keeper) GetAllSubscriptions(ctx sdk.Context) []*types.Subscription {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.SubscriptionKeyPrefix))
	defer iter.Close()

	var out []*types.Subscription
	for ; iter.Valid(); iter.Next() {
		var s types.Subscription
		if err := proto.Unmarshal(iter.Value(), &s); err == nil {
			out = append(out, &s)
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

// ── FeederDelegation ─────────────────────────────────────────────────────────

func delegationKey(delegator string) []byte {
	return append([]byte(types.FeederDelegationKeyPrefix), []byte(delegator)...)
}

func feederReverseKey(feeder string) []byte {
	return append([]byte(types.FeederReverseIndexPrefix), []byte(feeder)...)
}

func (k Keeper) SetDelegation(ctx sdk.Context, delegator, feeder string) error {
	d := &types.FeederDelegation{Delegator: delegator, Feeder: feeder}
	bz, err := proto.Marshal(d)
	if err != nil {
		return err
	}
	st := k.kvStore(ctx)
	// Remove old reverse index if there was a previous delegation.
	if old := st.Get(delegationKey(delegator)); old != nil {
		var prev types.FeederDelegation
		if err2 := proto.Unmarshal(old, &prev); err2 == nil && prev.Feeder != "" {
			st.Delete(feederReverseKey(prev.Feeder))
		}
	}
	st.Set(delegationKey(delegator), bz)
	st.Set(feederReverseKey(feeder), []byte(delegator))
	return nil
}

func (k Keeper) GetDelegation(ctx sdk.Context, delegator string) (*types.FeederDelegation, bool) {
	bz := k.kvStore(ctx).Get(delegationKey(delegator))
	if bz == nil {
		return nil, false
	}
	var d types.FeederDelegation
	if err := proto.Unmarshal(bz, &d); err != nil {
		return nil, false
	}
	return &d, true
}

// GetDelegatorForFeeder returns the delegator address for a given feeder address,
// or empty string if the feeder is not delegated by anyone.
func (k Keeper) GetDelegatorForFeeder(ctx sdk.Context, feeder string) string {
	bz := k.kvStore(ctx).Get(feederReverseKey(feeder))
	if bz == nil {
		return ""
	}
	return string(bz)
}

func (k Keeper) RevokeDelegation(ctx sdk.Context, delegator string) {
	st := k.kvStore(ctx)
	bz := st.Get(delegationKey(delegator))
	if bz != nil {
		var d types.FeederDelegation
		if err := proto.Unmarshal(bz, &d); err == nil && d.Feeder != "" {
			st.Delete(feederReverseKey(d.Feeder))
		}
	}
	st.Delete(delegationKey(delegator))
}

func (k Keeper) GetAllDelegations(ctx sdk.Context) []*types.FeederDelegation {
	st := k.kvStore(ctx)
	iter := storetypes.KVStorePrefixIterator(st, []byte(types.FeederDelegationKeyPrefix))
	defer iter.Close()

	var out []*types.FeederDelegation
	for ; iter.Valid(); iter.Next() {
		var d types.FeederDelegation
		if err := proto.Unmarshal(iter.Value(), &d); err == nil {
			out = append(out, &d)
		}
	}
	return out
}

// ── Genesis helpers ──────────────────────────────────────────────────────────

func (k Keeper) ImportGenesis(ctx sdk.Context, gs *types.GenesisState) {
	for _, f := range gs.Feeds {
		bz, _ := proto.Marshal(f)
		k.kvStore(ctx).Set(feedKey(f.Id), bz)
	}
	for _, s := range gs.Subscriptions {
		bz, _ := proto.Marshal(s)
		st := k.kvStore(ctx)
		st.Set(subscriptionKey(s.Id), bz)
		st.Set(subscriptionOwnerKey(s.Owner, s.Id), []byte{1})
		k.setSubscriptionCount(ctx, s.Id)
	}
	for _, f := range gs.Facts {
		bz, _ := proto.Marshal(f)
		st := k.kvStore(ctx)
		st.Set(factKey(f.Id), bz)
		st.Set(factDedupKey(f.FeedId, f.DeduplicationKey), []byte{1})
		st.Set(factFeedIndexKey(f.FeedId, f.Id), []byte{1})
		k.setFactCount(ctx, f.Id)
	}
	for _, wh := range gs.Webhooks {
		bz, _ := proto.Marshal(wh)
		st := k.kvStore(ctx)
		st.Set(webhookKey(wh.Id), bz)
		st.Set(ownerWebhookKey(wh.Owner, wh.Id), []byte{1})
		k.setWebhookCount(ctx, wh.Id)
	}
	for _, d := range gs.Delegations {
		_ = k.SetDelegation(ctx, d.Delegator, d.Feeder)
	}
}

func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Feeds:         k.GetAllFeeds(ctx),
		Subscriptions: k.GetAllSubscriptions(ctx),
		Facts:         k.GetAllFacts(ctx),
		Webhooks:      k.GetAllWebhooks(ctx),
		Delegations:   k.GetAllDelegations(ctx),
	}
}
