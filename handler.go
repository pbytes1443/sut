package sentinel

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgRegisterVpnService:
			return handleRegisterVpnService(ctx, k, msg)
		// case MsgQueryRegisteredVpnService:
		// 	return handleQueryRegisteredVpnService(ctx, k, msg)
		case MsgDeleteVpnUser:
			return handleDeleteVpnUser(ctx, k, msg)
		case MsgRegisterMasterNode:
			return handleMsgRegisterMasterNode(ctx, k, msg)
		// case MsgQueryFromMasterNode:
		// 	return handleMsgQueryFromMasterNode(ctx, k, msg)
		case MsgDeleteMasterNode:
			return handleMsgDeleteMasterNode(ctx, k, msg)
		case MsgPayVpnService:
			return handleMsgPayVpnService(ctx, k, msg)
		// case MsgSigntoVpn:
		// 	return handleMsgSend(ctx, k, msg)
		// case MsgGetVpnPayment:
		// 	return handleMsgSigntoChain(ctx, k, msg)
		// case MsgRefund:
		// 	return handleMsgRefund(ctx, k, msg)
		default:
			return sdk.ErrUnknownRequest("unrecognized message").Result()
		}
	}
}

func handleMsgRegisterMasterNode(ctx sdk.Context, keeper Keeper, msg MsgRegisterMasterNode) sdk.Result {
	id, err := keeper.RegisterMasterNode(ctx, msg)
	if err != nil {
		return err.Result() //CHANGE THIS SPECIFIC ERROR TYPE
	}
	//d, _ := keeper.cdc.MarshalBinary(id)
	return sdk.Result{
		Tags: msg.Tags(),
		Data: id,
	}
}

/*
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
*/
func handleRegisterVpnService(ctx sdk.Context, keeper Keeper, msg MsgRegisterVpnService) sdk.Result {

	/// BEFORE CALLING STORE HANDLER FUNCTION , PLEASE VERIFY SIGNATURE present inside the message
	id, err := keeper.RegisterVpnService(ctx, msg)
	if err != nil {
		return err.Result() /// PLEASE CHANGE THIS TO SPECIFIC ERROR TYPE
	}
	d, _ := keeper.cdc.MarshalBinary(id) ///CHECK WHHE DATA OR NOT

	///CHECK FOR ADDING TAGS MECHANISM
	/// WE SHOULD RETURN TAGS
	return sdk.Result{
		Data: d,
	}
}

/// TODO ://  IMPLEMENT ACL FOR THIS TRANSACTION
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

///TODO WHOLE FUNCTION SHOULD BE IMPLEMENTED WITH ACL
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

/// TODO
func handleMsgPayVpnService(ctx sdk.Context, keeper Keeper, msg MsgPayVpnService) sdk.Result {
	id, err := keeper.PayVpnService(ctx, msg)
	if err != nil {
		return err.Result()
	}
	d, _ := keeper.cdc.MarshalJSON(id)
	return sdk.Result{
		Data: d,
	}
}

// func handleMsgSend(ctx sdk.Context, k Keeper, msg MsgSigntoVpn) sdk.Result {
// 	// NOTE: totalIn == totalOut should already have been checked
// 	tags, err := k.sendSigntoVpn(ctx, msg)
// 	if err != nil {
// 		return err.Result()
// 	}

// 	return sdk.Result{
// 		Tags: tags,
// 	}
// }
/*
func handleMsgSigntoChain(ctx sdk.Context, k Keeper, msg MsgGetVpnPayment) sdk.Result {

	sessionid, err := k.GetVpnPayment(ctx, msg)
	if err != nil {
		return err.Result()
	}

	// func (msg MsgGetVpnPayment) Tags() sdk.Tags {
	// 	return sdk.NewTags("", []byte(msg.from.String())).
	// 		AppendTag("receiver", []byte(msg.Receiver.String()))
	// }
	tags := sdk.NewTags("Vpn Provider Address:", []byte(msg.from.String())).AppendTag("seesionId", sessionid)

	return sdk.Result{
		Tags: tags,
	}
}
func handleMsgRefund(ctx sdk.Context, k Keeper, msg MsgRefund) sdk.Result {
	address, err := k.RefundBal(ctx, msg)
	if err != nil {
		return err.Result()
	}
	tags := sdk.NewTags("client Refund Address:", []byte(address.String()))
	return sdk.Result{
		Tags: tags,
	}
}
*/
