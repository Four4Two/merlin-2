package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/four4two/merlin/v16/app/keepers"
	"github.com/four4two/merlin/v16/app/upgrades"
)

// CreateUpgradeHandler returns an x/upgrade handler for the Merlin v4 on-chain
// upgrade. Namely, it executes:
//
// 1. Setting x/gamm parameters for pool creation
// 2. Executing prop 12
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	_ upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// Kept as comments for recordkeeping. SetParams is now private:
		// 		keepers.GAMMKeeper.SetParams(ctx, gammtypes.NewParams(sdk.Coins{sdk.NewInt64Coin("ufury", 1)})) // 1 uFURY

		Prop12(ctx, keepers.BankKeeper, keepers.DistrKeeper)

		return vm, nil
	}
}
