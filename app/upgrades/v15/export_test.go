package v15

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	icqkeeper "github.com/cosmos/ibc-apps/modules/async-icq/v4/keeper"

	ibcratelimit "github.com/four4two/merlin/v16/x/ibc-rate-limit"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	gammkeeper "github.com/four4two/merlin/v16/x/gamm/keeper"
	"github.com/four4two/merlin/v16/x/poolmanager"
)

func MigrateNextPoolId(ctx sdk.Context, gammKeeper *gammkeeper.Keeper, poolmanager *poolmanager.Keeper) {
	migrateNextPoolId(ctx, gammKeeper, poolmanager)
}

func RegisterFuryIonMetadata(ctx sdk.Context, bankKeeper bankkeeper.Keeper) {
	registerFuryIonMetadata(ctx, bankKeeper)
}

func SetICQParams(ctx sdk.Context, icqKeeper *icqkeeper.Keeper) {
	setICQParams(ctx, icqKeeper)
}

func MigrateBalancerPoolToSolidlyStable(ctx sdk.Context, gammKeeper *gammkeeper.Keeper, poolmanager *poolmanager.Keeper, bankKeeper bankkeeper.Keeper, poolId uint64) {
	migrateBalancerPoolToSolidlyStable(ctx, gammKeeper, poolmanager, bankKeeper, poolId)
}

func SetRateLimits(ctx sdk.Context, accountKeeper *authkeeper.AccountKeeper, rateLimitingICS4Wrapper *ibcratelimit.ICS4Wrapper, wasmKeeper *wasmkeeper.Keeper) {
	setRateLimits(ctx, accountKeeper, rateLimitingICS4Wrapper, wasmKeeper)
}
