package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
)

type MsgRegisterVpnService struct {
	address  sdk.Address `json:"address",omitempty`
	ip       string      `json:"ip",omitempty`
	netspeed int64      `json:"netspeed",omitempty`
	ppgb     int64      `json:"ppgb",omitempty`
	location string     `json:"location",omitempty`
}
type MsgRegisterMasterNode struct {
	address sdk.Address   `json:"address",omitempty`
	//pubkey  crypto.PubKey `json:"pubkey",omitempty`
}

type MsgDeleteVpnUser struct {
	addressService sdk.Address `json:"address", omitempty`
}
type MsgDeleteMasterNode struct {
	address sdk.Address `json:"address", omitempty`
}
type MsgPayVpnService struct {
	coins   sdk.Coin       `json:"coins", omitempty`
	pubkey  *crypto.PubKey `json:"pubkey", omitempty`
	vpnaddr sdk.Address    `json:"address", omitempty`
}

type MsgSigntoVpn struct {
	coins     sdk.Coin       `json:"coins", omitempty`
	address   sdk.Adress     `json:"address", omitempty`
	sessionid int64          `json:"session_id", omitempty`
	signature auth.Signature `json:"signature", omitempty`
	from      sdk.Address    `json:"from", omitempty`
}

type MsgGetVpnPayment struct {
	clientSig types.ClientSignature `json:"signature", omitempty`
	from      sdk.Address           `json:"address", omitempty`
}

type MsgRefund struct {
	pubkey    crypto.PublicKey `json:"pubkey", omitempty`
	sessionid int64            `json:"session_id", omitempty`
}


func ServiceRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec) {

	r.HandleFunc(
		"/registervpn/",
		registervpnHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/register_master_node/",
		registermasterdHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/refund/",
		RefundHandleFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/deletemaster/{address}",
		deleteMasterHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/deletevpnnode/{address}",
		deleteVpnHandlerFn(ctx, cdc),
	).Methods("POST")

}
func registervpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var msg MsgRegisterVpnService
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid address."))
			return
		}
		if (!validateIp(msg.ip)){
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid Ip address."))
			return

		}
		if msg.ppgb == nil || reflect.TypeOf(msg.ppgb) != int64 || msg.ppgb > 0 || msg.ppgb < 1000 {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid amount of price per Gb"))
			return
		}
		if msg.netspeed == nil || reflect.TypeOf(msg.netspeed) != int64 || msg.netspeed > 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid details"))
			return
		}
		
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			return nil
}

func registermasterdHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	
		var msg MsgRegisterMasterNode
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &msg)

		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid address."))
			return
		}
		res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)

	}
}


func deleteMasterHandlerFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{

	return func(w http.ResponseWriter, r http.Request){
		var msg MsgDeleteMasterNode
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid address."))
			return
		}
		res,err:= ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
	}
}/*
func RefundHandleFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{
	return func(w http.ResponseWriter, r http.Request){

		var := mux.Vars(r)
		addres:= var["address"]
		pubkey:= var["public_key"]
		msg:= sender.MsgRefund{addres,pubkey}
		res, err :=ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
	}
}
	*/
	
	
func deleteVpnHandlerFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{

	return func(w http.ResponseWriter, r *http.Request){

		var msg MsgDeleteVpnUser
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid address."))
			return
		}
	
		// msg:= sentinel.MsgDeleteVpnUser{addres}
		res,err:= ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
	}
}

func validateIp(host string) bool {
	parts := strings.Split(host, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}

	}
	return true
}
