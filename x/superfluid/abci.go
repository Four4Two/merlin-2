package superfluid

import (
	"github.com/four4two/merlin/v17/x/superfluid/keeper"
	"github.com/four4two/merlin/v17/x/superfluid/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker is called on every block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ek types.EpochKeeper) {
	numBlocksSinceEpochStart, err := ek.NumBlocksSinceEpochStart(ctx, k.GetEpochIdentifier(ctx))
	if err != nil {
		panic(err)
	}
	if numBlocksSinceEpochStart == 0 {
		k.AfterEpochStartBeginBlock(ctx)
	}
}
