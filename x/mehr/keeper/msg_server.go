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

func (m MsgServer) RegisterFeed(goCtx context.Context, msg *types.MsgRegisterFeed) (*types.MsgRegisterFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if m.Keeper.FeedExists(ctx, msg.Id) {
		return nil, types.ErrFeedAlreadyExists
	}
	f := &types.Feed{
		Id:          msg.Id,
		Name:        msg.Name,
		Description: msg.Description,
		SourceType:  msg.SourceType,
		PayloadType: msg.PayloadType,
		CreatedBy:   msg.Creator,
	}
	saved, err := m.Keeper.RegisterFeed(ctx, f)
	if err != nil {
		return nil, err
	}
	return &types.MsgRegisterFeedResponse{Feed: saved}, nil
}

func (m MsgServer) SubmitFact(goCtx context.Context, msg *types.MsgSubmitFact) (*types.MsgSubmitFactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// Feeder must be delegated by someone (or be self-delegated).
	delegator := m.Keeper.GetDelegatorForFeeder(ctx, msg.Feeder)
	if delegator == "" {
		if _, ok := m.Keeper.GetDelegation(ctx, msg.Feeder); !ok {
			return nil, types.ErrUnauthorizedFeeder
		}
	}

	if !m.Keeper.FeedExists(ctx, msg.FeedId) {
		return nil, types.ErrFeedNotFound
	}

	if m.Keeper.IsDuplicateFact(ctx, msg.FeedId, msg.DeduplicationKey) {
		return nil, types.ErrDuplicateFact
	}

	f := &types.Fact{
		FeedId:           msg.FeedId,
		Feeder:           msg.Feeder,
		Payload:          msg.Payload,
		DeduplicationKey: msg.DeduplicationKey,
		SubmittedAt:      ctx.BlockTime().Format(time.RFC3339),
	}

	saved, err := m.Keeper.CreateFact(ctx, f)
	if err != nil {
		return nil, err
	}
	return &types.MsgSubmitFactResponse{Fact: saved}, nil
}

func (m MsgServer) CreateSubscription(goCtx context.Context, msg *types.MsgCreateSubscription) (*types.MsgCreateSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if !m.Keeper.FeedExists(ctx, msg.FeedId) {
		return nil, types.ErrFeedNotFound
	}
	s := &types.Subscription{
		FeedId: msg.FeedId,
		Owner:  msg.Creator,
		Filter: msg.Filter,
		Label:  msg.Label,
	}
	saved, err := m.Keeper.CreateSubscription(ctx, s)
	if err != nil {
		return nil, err
	}
	return &types.MsgCreateSubscriptionResponse{Subscription: saved}, nil
}

func (m MsgServer) DeleteSubscription(goCtx context.Context, msg *types.MsgDeleteSubscription) (*types.MsgDeleteSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	s, ok := m.Keeper.GetSubscription(ctx, msg.SubscriptionId)
	if !ok {
		return nil, types.ErrSubscriptionNotFound
	}
	if s.Owner != msg.Creator {
		return nil, types.ErrUnauthorized
	}
	m.Keeper.DeleteSubscription(ctx, msg.SubscriptionId, msg.Creator)
	return &types.MsgDeleteSubscriptionResponse{}, nil
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
