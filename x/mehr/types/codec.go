package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWatch{},
		&MsgDeleteWatch{},
		&MsgCreateWebhook{},
		&MsgDeleteWebhook{},
		&MsgSubmitEvent{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &Msg_serviceDesc)
}
