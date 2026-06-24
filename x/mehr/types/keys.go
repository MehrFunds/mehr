package types

const (
	ModuleName = "mehr"
	StoreKey   = ModuleName

	// Feed: string-keyed
	FeedKeyPrefix = "f/"

	// Fact: uint64-keyed
	FactCountKey        = "fac"
	FactKeyPrefix       = "fa/"
	FactFeedIndexPrefix = "ff/"
	FactDedupPrefix     = "fdd/"

	// Subscription: uint64-keyed
	SubscriptionCountKey        = "sc"
	SubscriptionKeyPrefix       = "s/"
	SubscriptionOwnerIndexPrefix = "so/"

	// Webhook: uint64-keyed (unchanged)
	WebhookCountKey         = "whc"
	WebhookKeyPrefix        = "wh/"
	OwnerWebhookIndexPrefix = "owh/"

	// FeederDelegation: delegator → FeederDelegation
	FeederDelegationKeyPrefix = "fd/"
	// reverse index: feeder → delegator address
	FeederReverseIndexPrefix = "fr/"
)
