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

func (q QueryServer) Watches(goCtx context.Context, req *types.QueryWatchesRequest) (*types.QueryWatchesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryWatchesResponse{Watches: q.Keeper.GetWatchesByOwner(ctx, req.Owner)}, nil
}

func (q QueryServer) Watch(goCtx context.Context, req *types.QueryWatchRequest) (*types.QueryWatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	w, ok := q.Keeper.GetWatch(ctx, req.Id)
	if !ok {
		return nil, types.ErrWatchNotFound
	}
	return &types.QueryWatchResponse{Watch: w}, nil
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

func (q QueryServer) AllWatches(goCtx context.Context, _ *types.QueryAllWatchesRequest) (*types.QueryAllWatchesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllWatchesResponse{Watches: q.Keeper.GetAllWatches(ctx)}, nil
}

func (q QueryServer) AllWebhooks(goCtx context.Context, _ *types.QueryAllWebhooksRequest) (*types.QueryAllWebhooksResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllWebhooksResponse{Webhooks: q.Keeper.GetAllWebhooks(ctx)}, nil
}

func (q QueryServer) Events(goCtx context.Context, req *types.QueryEventsRequest) (*types.QueryEventsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryEventsResponse{Events: q.Keeper.GetEventsByAddress(ctx, req.Address)}, nil
}

func (q QueryServer) Event(goCtx context.Context, req *types.QueryEventRequest) (*types.QueryEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	e, ok := q.Keeper.GetEvent(ctx, req.Id)
	if !ok {
		return nil, types.ErrEventNotFound
	}
	return &types.QueryEventResponse{Event: e}, nil
}
