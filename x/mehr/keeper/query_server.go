package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mehrfunds/mehr/x/mehr/types"
)

// QueryServer handles read-only queries for the mehr module.
type QueryServer struct {
	Keeper
}

func NewQueryServer(k Keeper) QueryServer {
	return QueryServer{Keeper: k}
}

func (q QueryServer) Feed(goCtx context.Context, req *types.QueryFeedRequest) (*types.QueryFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	f, ok := q.Keeper.GetFeed(ctx, req.Id)
	if !ok {
		return nil, types.ErrFeedNotFound
	}
	return &types.QueryFeedResponse{Feed: f}, nil
}

func (q QueryServer) AllFeeds(goCtx context.Context, _ *types.QueryAllFeedsRequest) (*types.QueryAllFeedsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllFeedsResponse{Feeds: q.Keeper.GetAllFeeds(ctx)}, nil
}

func (q QueryServer) Fact(goCtx context.Context, req *types.QueryFactRequest) (*types.QueryFactResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	f, ok := q.Keeper.GetFact(ctx, req.Id)
	if !ok {
		return nil, types.ErrFactNotFound
	}
	return &types.QueryFactResponse{Fact: f}, nil
}

func (q QueryServer) FactsByFeed(goCtx context.Context, req *types.QueryFactsByFeedRequest) (*types.QueryFactsByFeedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryFactsByFeedResponse{Facts: q.Keeper.GetFactsByFeed(ctx, req.FeedId)}, nil
}

func (q QueryServer) Subscription(goCtx context.Context, req *types.QuerySubscriptionRequest) (*types.QuerySubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	s, ok := q.Keeper.GetSubscription(ctx, req.Id)
	if !ok {
		return nil, types.ErrSubscriptionNotFound
	}
	return &types.QuerySubscriptionResponse{Subscription: s}, nil
}

func (q QueryServer) SubscriptionsByOwner(goCtx context.Context, req *types.QuerySubscriptionsByOwnerRequest) (*types.QuerySubscriptionsByOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QuerySubscriptionsByOwnerResponse{Subscriptions: q.Keeper.GetSubscriptionsByOwner(ctx, req.Owner)}, nil
}

func (q QueryServer) AllSubscriptions(goCtx context.Context, _ *types.QueryAllSubscriptionsRequest) (*types.QueryAllSubscriptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllSubscriptionsResponse{Subscriptions: q.Keeper.GetAllSubscriptions(ctx)}, nil
}

func (q QueryServer) AllWebhooks(goCtx context.Context, _ *types.QueryAllWebhooksRequest) (*types.QueryAllWebhooksResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllWebhooksResponse{Webhooks: q.Keeper.GetAllWebhooks(ctx)}, nil
}

func (q QueryServer) Webhooks(goCtx context.Context, req *types.QueryWebhooksRequest) (*types.QueryWebhooksResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryWebhooksResponse{Webhooks: q.Keeper.GetWebhooksByOwner(ctx, req.Owner)}, nil
}

func (q QueryServer) Webhook(goCtx context.Context, req *types.QueryWebhookRequest) (*types.QueryWebhookResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	wh, ok := q.Keeper.GetWebhook(ctx, req.Id)
	if !ok {
		return nil, types.ErrWebhookNotFound
	}
	return &types.QueryWebhookResponse{Webhook: wh}, nil
}

func (q QueryServer) FeederDelegation(goCtx context.Context, req *types.QueryFeederDelegationRequest) (*types.QueryFeederDelegationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	d, ok := q.Keeper.GetDelegation(ctx, req.Delegator)
	if !ok {
		return nil, types.ErrDelegationNotFound
	}
	return &types.QueryFeederDelegationResponse{Delegation: d}, nil
}

func (q QueryServer) AllDelegations(goCtx context.Context, _ *types.QueryAllDelegationsRequest) (*types.QueryAllDelegationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllDelegationsResponse{Delegations: q.Keeper.GetAllDelegations(ctx)}, nil
}
