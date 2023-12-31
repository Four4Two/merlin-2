package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLockTokens{}, "merlin/lockup/lock-tokens", nil)
	cdc.RegisterConcrete(&MsgBeginUnlockingAll{}, "merlin/lockup/begin-unlock-tokens", nil)
	cdc.RegisterConcrete(&MsgBeginUnlocking{}, "merlin/lockup/begin-unlock-period-lock", nil)
	cdc.RegisterConcrete(&MsgExtendLockup{}, "merlin/lockup/extend-lockup", nil)
	cdc.RegisterConcrete(&MsgForceUnlock{}, "merlin/lockup/force-unlock-tokens", nil)
	cdc.RegisterConcrete(&MsgSetRewardReceiverAddress{}, "merlin/lockup/set-reward-receiver-address", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgLockTokens{},
		&MsgBeginUnlockingAll{},
		&MsgBeginUnlocking{},
		&MsgExtendLockup{},
		&MsgForceUnlock{},
		&MsgSetRewardReceiverAddress{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterCodec(authzcodec.Amino)

	amino.Seal()
}
