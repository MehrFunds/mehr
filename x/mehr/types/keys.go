package types

const (
	ModuleName = "mehr"
	StoreKey   = ModuleName

	WatchCountKey    = "wc"
	WatchKeyPrefix   = "w/"
	WebhookCountKey  = "whc"
	WebhookKeyPrefix = "wh/"

	// secondary index: "<owner>/<id_8bytes>" → empty
	OwnerWatchIndexPrefix   = "ow/"
	OwnerWebhookIndexPrefix = "owh/"

	EventCountKey   = "ec"
	EventKeyPrefix  = "e/"
	// dedup index: network/txhash/logindex → event_id
	EventDedupPrefix = "ed/"
	// address index: address/event_id → empty
	EventAddrIndexPrefix = "ea/"
)
