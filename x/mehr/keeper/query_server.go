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
