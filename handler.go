package sentinel

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgRegisterVpnService:
			return handleRegisterVpnService(ctx, k, msg)
		case MsgQueryRegisteredVpnService:
			return handleQueryRegisteredVpnService(ctx, k, msg)
		case MsgDeleteVpnUser:
			return handleDeleteVpnUser(ctx, k, msg)
		case MsgRegisterMasterNode:
			return handleMsgRegisterMasterNode(ctx, k, msg)
		case MsgQueryFromMasterNode:
			return handleMsgQueryFromMasterNode(ctx, k, msg)
		case MsgDeleteMasterNode:
			return handleMsgDeleteMasterNode(ctx, k, msg)
		case MsgPayVpnService:
			return handleMsgPayVpnService(ctx, k, msg)
		case MsgSigntoVpn:
			return handleMsgSend(ctx, k, msg)
		case MsgSigntoChain:
			return handleMsgSigntoChain(ctx, k, msg)
		case MsgRefund:
			return handleMsgRefund(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("unrecognized message").Result()
		}
	}
}

func handleMsgRegisterMasterNode(ctx sdk.Context, keeper Keeper, msg MsgRegisterMasterNode) sdk.Result {
	id, err := keeper.RegisterMasterNode(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalBinary(id)
	return sdk.Result{
		Data: d,
	}
}

func handleMsgQueryFromMasterNode(ctx sdk.Context, keeper Keeper, msg MsgQueryFromMasterNode) sdk.Result {

	id, err := keeper.QueryFromRegisterMasterNode(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalBinary(id)
	return sdk.Result{
		Data: d,
	}

}

func handleRegisterVpnService(ctx sdk.Context, keeper Keeper, msg MsgRegisterVpnService) sdk.Result {
	id, err := keeper.RegisterVpnService(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalBinary(id)
	return sdk.Result{
		Data: d,
	}
}
func handleQueryRegisteredVpnService(ctx sdk.Context, keeper Keeper, msg MsgQueryRegisteredVpnService) sdk.Result {
	id, err := keeper.QueryRegisteredVpnService(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalJSON(id)
	return sdk.Result{
		Data: d,
	}
}
func handleDeleteVpnUser(ctx sdk.Context, keeper Keeper, msg MsgDeleteVpnUser) sdk.Result {
	id, err := keeper.DeleteVpnService(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalBinary(id)
	return sdk.Result{
		Data: d,
	}
}
func handleMsgDeleteMasterNode(ctx sdk.Context, keeper Keeper, msg MsgDeleteMasterNode) sdk.Result {
	id, err := keeper.DeleteMasterNode(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalJSON(id)
	return sdk.Result{
		Data: d,
	}
}
func handleMsgPayVpnService(ctx sdk.Context, keeper Keeper, msg MsgPayVpnService) sdk.Result {
	id, err := keeper.MsgPayVpnService(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalJSON(id)
	return sdk.Result{
		Data: d,
	}
}
func handleMsgSend(ctx sdk.Context, k Keeper, msg MsgSigntoVpn) sdk.Result {
	// NOTE: totalIn == totalOut should already have been checked
	tags, err := k.sendSigntoVpn(ctx, msg)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: tags,
	}
}

func handleMsgSigntoChain(ctx sdk.Context, k Keeper, msg MsgGetVpnPayment) sdk.Result {

	tags, err := k.sendSigntoChain(ctx, msg)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: tags,
	}
}
func handleMsgRefund(ctx sdk.Context, k Keeper, msg MsgRefund) sdk.Result {
	tags, err := k.RefundBal(ctx, msg)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Tags: tags
	}
}
