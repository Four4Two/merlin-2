package client

import (
	"github.com/four4two/merlin/v16/x/superfluid/client/cli"
	"github.com/four4two/merlin/v16/x/superfluid/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	SetSuperfluidAssetsProposalHandler    = govclient.NewProposalHandler(cli.NewCmdSubmitSetSuperfluidAssetsProposal, rest.ProposalSetSuperfluidAssetsRESTHandler)
	RemoveSuperfluidAssetsProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitRemoveSuperfluidAssetsProposal, rest.ProposalRemoveSuperfluidAssetsRESTHandler)
	UpdateUnpoolWhitelistProposalHandler  = govclient.NewProposalHandler(cli.NewCmdUpdateUnpoolWhitelistProposal, rest.ProposalUpdateUnpoolWhitelistProposal)
)
