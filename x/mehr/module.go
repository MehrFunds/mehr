package mehr

import (
	"context"
	"encoding/json"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	"cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/mehrfunds/mehr/x/mehr/keeper"
	"github.com/mehrfunds/mehr/x/mehr/types"
)

var (
	_ module.AppModuleBasic = AppModule{}
	_ module.AppModule      = AppModule{}
	_ appmodule.AppModule   = AppModule{}
)

// AppModule implements the mehr Cosmos SDK module.
type AppModule struct {
	keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{keeper: k}
}

// ── depinject.OnePerModuleType (required by appmodule.AppModule) ──────────────

func (AppModule) IsOnePerModuleType() {}

// ── appmodule.AppModule (marker) ─────────────────────────────────────────────

func (AppModule) IsAppModule() {}

// ── module.HasName ────────────────────────────────────────────────────────────

func (AppModule) Name() string { return types.ModuleName }

// ── module.AppModuleBasic ─────────────────────────────────────────────────────

func (AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (AppModule) RegisterInterfaces(codectypes.InterfaceRegistry) {}

func (AppModule) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	gs := &types.GenesisState{Watches: []*types.Watch{}, Webhooks: []*types.Webhook{}}
	bz, _ := json.Marshal(gs)
	return bz
}

func (AppModule) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var gs types.GenesisState
	return json.Unmarshal(bz, &gs)
}

func (AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx.GRPCClient)); err != nil {
		panic(err)
	}
}

func (AppModule) GetTxCmd() *cobra.Command   { return nil }
func (AppModule) GetQueryCmd() *cobra.Command { return nil }

// ── module.AppModule ──────────────────────────────────────────────────────────

func (m AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServer(m.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(m.keeper))
}

func (m AppModule) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) {
	var gs types.GenesisState
	if err := json.Unmarshal(data, &gs); err != nil {
		panic(err)
	}
	m.keeper.ImportGenesis(ctx, &gs)
}

func (m AppModule) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) json.RawMessage {
	gs := m.keeper.ExportGenesis(ctx)
	bz, err := json.Marshal(gs)
	if err != nil {
		panic(err)
	}
	return bz
}

func (AppModule) ConsensusVersion() uint64 { return 1 }
