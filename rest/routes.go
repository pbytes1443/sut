package rest

import (
	"encoding/json"
	"io/ioutill"
	"net/http"

	sentinel "github.com/cosmos/cosmos-sdk/examples/sentinel"
	"github.com/gorilla/mux"
	//"flags"
)

type MsgRegisterVpnService struct {
	address  sdk.Address `json:"address",omitempty`
	ip       string      `json:"ip",omitempty`
	netspeed string      `json:"netspeed",omitempty`
	ppgb     string      `json:"ppgb",omitempty`
}
type MsgRegisterMasterNode struct {
	address sdk.Address		 `json:"address",omitempty`
	pubkey  crypto.PubKey	 `json:"pubkey",omitempty`
}
type MsgQueryRegisteredVpnService struct {
	address sdk.Address  `json:"address",omitempty`
}
type MsgQueryFromMasterNode struct {
	address sdk.Address 	`json:"address",omitempty`
}
type MsgDeleteVpnUser struct {
	addressService sdk.Address `json:"address", omitempty`
}
type MsgDeleteMasterNode struct {
	address sdk.Address	`json:"address", omitempty`
}
type MsgPayVpnService struct {
	coins   sdk.Coin	`json:"coins", omitempty`
	pubkey  *crypto.PubKey	`json:"pubkey", omitempty`
	vpnaddr sdk.Address	`json:"address", omitempty`
}


type MsgSigntoVpn struct {
	coins sdk.Coin	`json:"coins", omitempty`
	address   sdk.Adress	`json:"address", omitempty`
	sessionid int64	`json:"session_id", omitempty`
	signature auth.Signature	`json:"signature", omitempty`
	from      sdk.Address	`json:"from", omitempty`
}

type MsgGetVpnPayment struct {
	clientSig types.ClientSignature `json:"signature", omitempty`
	from sdk.Address	`json:"address", omitempty`
}

type MsgRefund struct{
	pubkey crypto.PublicKey `json:"pubkey", omitempty`
	sessionid int64	`json:"session_id", omitempty`
}


func RegisterRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec) {

	r.HandleFunc(
		"/registervpn/{address}/{ip}/{ppgb}/{netspeed}",
		registervpnHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/register_master_node/{addres}/{publick_key}",
		registermasterdHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/query_vpn/{address}",
		queryvpnHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/query_master_node/{address}",
		querymasterHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/delete_master/{address}",
		deleteMasterHandlerFn(ctx, cdc),
	).Methods("POST")
	r.HandleFunc(
		"/refund/{address}/{public_key}",
		RefundHandleFn(ctx, cdc),
	).Methods("POST")
}
func registervpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var s MsgRegisterVpnService
		// var err error
		// body, err := ioutill.ReadAll(r.Body)
		// json.Unmarshal(body, &s)
		// read parameters
			vars := mux.Vars(r)
			address := vars["address"]
			ip := vars["ip"]
			ppgb := vars["ppgb"]
			netspeed := vars["netspeed"]
			msg := sentinel.MsgRegisterVpnService{address, ip, ppgb, netspeed}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			fmt.Printf("Vpn serivicer registered with address: %s\n", sender)
			return nil
}
func registermasterHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var s MsgRegisterMasterNode
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &s)
		// read parameters
		vars := mux.Vars(r)
		address := vars["address"]
		ip := vars["public_key"]
		msg := sentinel.MsgRegisterMasterNode{address, ip}
		res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)

	}
}

func queryvpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var s MsgQueryRegisteredVpnService
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &s)

		// read parameters
		vars := mux.Vars(r)
		address := vars["address"]
		var ctx1 sdk.context
		var k sent.Keeper
		
		msg := sentinel.MsgQueryRegisteredVpnService{address}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)

	}
}
func querymasterHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read parameters
		vars := mux.Vars(r)
		address := vars["address"]
		msg := sentinel.MsgQueryFromMasterNode{address}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)

	}
}

func deleteVpnHandlerFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{

	return func(w http.ResponseWriter, r *http.Request){

		var:= mux.Vars(r)
		addres:= var["address"]
		msg:= sentinel.MsgDeleteVpnUser{addres}
		res,err:= ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
	}
}

func deleteMasterHandlerFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{

	return func(w http.ResponseWriter, r http.Request){

		var:= mux.Vars(r)
		addres:=var["addess"]
		msg:= sentinel.MsgDeleteMasterNode{addres}
		res,err:= ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
	}
}
func RefundHandleFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{
	return func(w http.ResponseWriter, r http.Request){

		var := mux.Vars(r)
		addres:= var["address"]
		pubkey:= var["public_key"]
		msg:= sender.MsgRefund{addres,pubkey}
		res, err :=ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
	}
}
//cosmos-sdk/stake/query......
