package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidAddress  = sdkerrors.Register(ModuleName, 1, "invalid address")
	ErrWatchNotFound   = sdkerrors.Register(ModuleName, 2, "watch not found")
	ErrWebhookNotFound = sdkerrors.Register(ModuleName, 3, "webhook not found")
	ErrUnauthorized    = sdkerrors.Register(ModuleName, 4, "unauthorized")
	ErrInvalidURL      = sdkerrors.Register(ModuleName, 5, "invalid webhook URL")
	ErrEventNotFound   = sdkerrors.Register(ModuleName, 6, "event not found")
	ErrDuplicateEvent  = sdkerrors.Register(ModuleName, 7, "event already recorded")
	ErrWatchNotActive      = sdkerrors.Register(ModuleName, 8, "no active watch for this network/address")
	ErrUnauthorizedFeeder  = sdkerrors.Register(ModuleName, 9, "feeder is not authorized — run: mehrd tx mehr delegate-feeder <feeder-address>")
	ErrDelegationNotFound  = sdkerrors.Register(ModuleName, 10, "feeder delegation not found")
)
