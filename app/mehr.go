package app

import (
	"cosmossdk.io/core/appmodule"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"

	mehr "github.com/mehrfunds/mehr/x/mehr"
	"github.com/mehrfunds/mehr/x/mehr/keeper"
	mehrtypes "github.com/mehrfunds/mehr/x/mehr/types"
)

// registerMehrModule wires up the mehr custom module on the server side.
// Follows the same pattern as registerIBCModules.
func (app *App) registerMehrModule(_ servertypes.AppOptions) error {
	storeKey := storetypes.NewKVStoreKey(mehrtypes.StoreKey)
	if err := app.RegisterStores(storeKey); err != nil {
		return err
	}

	app.MehrKeeper = keeper.NewKeeper(storeKey)

	return app.RegisterModules(mehr.NewAppModule(app.MehrKeeper))
}

// RegisterMehr returns the mehr module for client-side (CLI/genesis) registration.
// Called from cmd/root.go so the moduleBasicManager includes it in genesis init output.
func RegisterMehr(_ codec.Codec) map[string]appmodule.AppModule {
	return map[string]appmodule.AppModule{
		mehrtypes.ModuleName: mehr.NewAppModule(keeper.Keeper{}),
	}
}
