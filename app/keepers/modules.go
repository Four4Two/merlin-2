package keepers

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	transfer "github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	"github.com/strangelove-ventures/packet-forward-middleware/v4/router"

	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	icq "github.com/cosmos/ibc-apps/modules/async-icq/v4"
	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"

	_ "github.com/four4two/merlin/v17/client/docs/statik"
	clclient "github.com/four4two/merlin/v17/x/concentrated-liquidity/client"
	concentratedliquidity "github.com/four4two/merlin/v17/x/concentrated-liquidity/clmodule"
	cwpoolclient "github.com/four4two/merlin/v17/x/cosmwasmpool/client"
	cosmwasmpoolmodule "github.com/four4two/merlin/v17/x/cosmwasmpool/module"
	downtimemodule "github.com/four4two/merlin/v17/x/downtime-detector/module"
	"github.com/four4two/merlin/v17/x/gamm"
	gammclient "github.com/four4two/merlin/v17/x/gamm/client"
	"github.com/four4two/merlin/v17/x/ibc-rate-limit/ibcratelimitmodule"
	"github.com/four4two/merlin/v17/x/incentives"
	"github.com/four4two/merlin/v17/x/lockup"
	"github.com/four4two/merlin/v17/x/mint"
	poolincentives "github.com/four4two/merlin/v17/x/pool-incentives"
	poolincentivesclient "github.com/four4two/merlin/v17/x/pool-incentives/client"
	poolmanager "github.com/four4two/merlin/v17/x/poolmanager/module"
	"github.com/four4two/merlin/v17/x/protorev"
	superfluid "github.com/four4two/merlin/v17/x/superfluid"
	superfluidclient "github.com/four4two/merlin/v17/x/superfluid/client"
	"github.com/four4two/merlin/v17/x/tokenfactory"
	"github.com/four4two/merlin/v17/x/twap/twapmodule"
	"github.com/four4two/merlin/v17/x/txfees"
	txfeesclient "github.com/four4two/merlin/v17/x/txfees/client"
	valsetprefmodule "github.com/four4two/merlin/v17/x/valset-pref/valpref-module"
	"github.com/osmosis-labs/osmosis/x/epochs"
	ibc_hooks "github.com/osmosis-labs/osmosis/x/ibc-hooks"
)

// AppModuleBasics returns ModuleBasics for the module BasicManager.
var AppModuleBasics = []module.AppModuleBasic{
	auth.AppModuleBasic{},
	genutil.AppModuleBasic{},
	bank.AppModuleBasic{},
	capability.AppModuleBasic{},
	staking.AppModuleBasic{},
	mint.AppModuleBasic{},
	downtimemodule.AppModuleBasic{},
	distr.AppModuleBasic{},
	gov.NewAppModuleBasic(
		append(
			wasmclient.ProposalHandlers,
			paramsclient.ProposalHandler,
			distrclient.ProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
			poolincentivesclient.UpdatePoolIncentivesHandler,
			poolincentivesclient.ReplacePoolIncentivesHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
			superfluidclient.SetSuperfluidAssetsProposalHandler,
			superfluidclient.RemoveSuperfluidAssetsProposalHandler,
			superfluidclient.UpdateUnpoolWhitelistProposalHandler,
			gammclient.ReplaceMigrationRecordsProposalHandler,
			gammclient.UpdateMigrationRecordsProposalHandler,
			clclient.CreateConcentratedLiquidityPoolProposalHandler,
			clclient.TickSpacingDecreaseProposalHandler,
			cwpoolclient.UploadCodeIdAndWhitelistProposalHandler,
			cwpoolclient.MigratePoolContractsProposalHandler,
			txfeesclient.SubmitUpdateFeeTokenProposalHandler,
		)...,
	),
	params.AppModuleBasic{},
	crisis.AppModuleBasic{},
	slashing.AppModuleBasic{},
	authzmodule.AppModuleBasic{},
	ibc.AppModuleBasic{},
	upgrade.AppModuleBasic{},
	evidence.AppModuleBasic{},
	transfer.AppModuleBasic{},
	vesting.AppModuleBasic{},
	gamm.AppModuleBasic{},
	poolmanager.AppModuleBasic{},
	twapmodule.AppModuleBasic{},
	concentratedliquidity.AppModuleBasic{},
	protorev.AppModuleBasic{},
	txfees.AppModuleBasic{},
	incentives.AppModuleBasic{},
	lockup.AppModuleBasic{},
	poolincentives.AppModuleBasic{},
	epochs.AppModuleBasic{},
	superfluid.AppModuleBasic{},
	tokenfactory.AppModuleBasic{},
	valsetprefmodule.AppModuleBasic{},
	wasm.AppModuleBasic{},
	icq.AppModuleBasic{},
	ica.AppModuleBasic{},
	ibc_hooks.AppModuleBasic{},
	ibcratelimitmodule.AppModuleBasic{},
	router.AppModuleBasic{},
	cosmwasmpoolmodule.AppModuleBasic{},
}
