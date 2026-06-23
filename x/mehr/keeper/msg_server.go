package keeper

import (
	"context"
	"net/url"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mehrfunds/mehr/x/mehr/types"
)

// MsgServer handles state-mutating messages for the mehr module.
type MsgServer struct {
	Keeper
}

func NewMsgServer(k Keeper) MsgServer {
	return MsgServer{Keeper: k}
}

func (m MsgServer) CreateWatch(goCtx context.Context, msg *types.MsgCreateWatch) (*types.MsgCreateWatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Address == "" {
		return nil, types.ErrInvalidAddress
	}
	w, err := m.Keeper.CreateWatch(ctx, msg.Creator, msg.Network, msg.Address, msg.Label)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateWatchResponse{Watch: w}, nil
}

func (m MsgServer) DeleteWatch(goCtx context.Context, msg *types.MsgDeleteWatch) (*types.MsgDeleteWatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	w, ok := m.Keeper.GetWatch(ctx, msg.WatchId)
	if !ok {
		return nil, types.ErrWatchNotFound
	}
	if w.Owner != msg.Creator {
		return nil, types.ErrUnauthorized
	}
	m.Keeper.DeleteWatch(ctx, msg.WatchId, msg.Creator)
	return &types.MsgDeleteWatchResponse{}, nil
}

func (m MsgServer) CreateWebhook(goCtx context.Context, msg *types.MsgCreateWebhook) (*types.MsgCreateWebhookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, err := url.ParseRequestURI(msg.Url); err != nil {
		return nil, types.ErrInvalidURL
	}
	wh, err := m.Keeper.CreateWebhook(ctx, msg.Creator, msg.Url, msg.SecretHash)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateWebhookResponse{Webhook: wh}, nil
}

func (m MsgServer) DeleteWebhook(goCtx context.Context, msg *types.MsgDeleteWebhook) (*types.MsgDeleteWebhookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	wh, ok := m.Keeper.GetWebhook(ctx, msg.WebhookId)
	if !ok {
		return nil, types.ErrWebhookNotFound
	}
	if wh.Owner != msg.Creator {
		return nil, types.ErrUnauthorized
	}
	m.Keeper.DeleteWebhook(ctx, msg.WebhookId, msg.Creator)
	return &types.MsgDeleteWebhookResponse{}, nil
}

func (m MsgServer) DelegateFeeder(goCtx context.Context, msg *types.MsgDelegateFeeder) (*types.MsgDelegateFeederResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if err := m.Keeper.SetDelegation(ctx, msg.Delegator, msg.Feeder); err != nil {
		return nil, err
	}
	return &types.MsgDelegateFeederResponse{}, nil
}

func (m MsgServer) RevokeDelegation(goCtx context.Context, msg *types.MsgRevokeDelegation) (*types.MsgRevokeDelegationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if _, ok := m.Keeper.GetDelegation(ctx, msg.Delegator); !ok {
		return nil, types.ErrDelegationNotFound
	}
	m.Keeper.RevokeDelegation(ctx, msg.Delegator)
	return &types.MsgRevokeDelegationResponse{}, nil
}

func (m MsgServer) SubmitEvent(goCtx context.Context, msg *types.MsgSubmitEvent) (*types.MsgSubmitEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Feeder must be delegated by someone (or be self-delegated).
	// Self-delegation: delegator == feeder. External: reverse index maps feeder → delegator.
	delegator := m.Keeper.GetDelegatorForFeeder(ctx, msg.Feeder)
	if delegator == "" {
		// Check if feeder is a self-delegator (delegated themselves).
		if _, ok := m.Keeper.GetDelegation(ctx, msg.Feeder); !ok {
			return nil, types.ErrUnauthorizedFeeder
		}
	}

	watch := m.Keeper.FindWatchForAddress(ctx, msg.Network, msg.Address)
	if watch == nil {
		return nil, types.ErrWatchNotActive
	}

	if m.Keeper.IsDuplicateEvent(ctx, msg.Network, msg.TxHash, msg.LogIndex) {
		return nil, types.ErrDuplicateEvent
	}

	e := &types.Event{
		Feeder:      msg.Feeder,
		WatchId:     watch.Id,
		Network:     msg.Network,
		Address:     msg.Address,
		TxHash:      msg.TxHash,
		BlockNumber: msg.BlockNumber,
		LogIndex:    msg.LogIndex,
		Asset:       msg.Asset,
		Contract:    msg.Contract,
		FromAddress: msg.FromAddress,
		AmountRaw:   msg.AmountRaw,
		Decimals:    msg.Decimals,
		Metadata:    msg.Metadata,
		SubmittedAt: ctx.BlockTime().Format(time.RFC3339),
	}

	saved, err := m.Keeper.CreateEvent(ctx, e)
	if err != nil {
		return nil, err
	}

	return &types.MsgSubmitEventResponse{Event: saved}, nil
}
