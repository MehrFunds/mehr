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
)
