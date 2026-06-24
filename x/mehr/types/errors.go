package types

import sdkerrors "cosmossdk.io/errors"

var (
	ErrInvalidAddress     = sdkerrors.Register(ModuleName, 1, "invalid address")
	ErrWebhookNotFound    = sdkerrors.Register(ModuleName, 2, "webhook not found")
	ErrUnauthorized       = sdkerrors.Register(ModuleName, 3, "unauthorized")
	ErrInvalidURL         = sdkerrors.Register(ModuleName, 4, "invalid webhook URL")
	ErrFeedNotFound       = sdkerrors.Register(ModuleName, 5, "feed not found")
	ErrFeedAlreadyExists  = sdkerrors.Register(ModuleName, 6, "feed already exists")
	ErrFactNotFound       = sdkerrors.Register(ModuleName, 7, "fact not found")
	ErrDuplicateFact      = sdkerrors.Register(ModuleName, 8, "fact already recorded")
	ErrSubscriptionNotFound = sdkerrors.Register(ModuleName, 9, "subscription not found")
	ErrUnauthorizedFeeder = sdkerrors.Register(ModuleName, 10, "feeder is not authorized — run: mehrd tx mehr delegate-feeder <feeder-address>")
	ErrDelegationNotFound = sdkerrors.Register(ModuleName, 11, "feeder delegation not found")
)
