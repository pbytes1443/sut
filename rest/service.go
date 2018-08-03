package rest

import (
	"encoding/json"

	// "fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	ioutill "io/ioutil"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/examples/sentinel"
	senttype "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/gorilla/mux"
	log "github.com/logger"
)

type MsgRegisterVpnService struct {
	//Address      string `json:"address"`
	Ip           string `json:"ip"`
	Netspeed     int64  `json:"netspeed"`
	Ppgb         int64  `json:"ppgb"`
	Location     string `json:"location"`
	Localaccount string `json:"account"`
	//Password     string `json:"password"`
	ChainID string `json:"chain-id"`
	Gas     int64  `json:"gas"`
	//Sequence     int64  `json:"sequence"`
}
type MsgRegisterMasterNode struct {
	//Address string `json:"address",omitempty`
	Name    string `json:"name"`
	ChainID string `json:"chain-id"`
	Gas     int64  `json:"gas"`
	//pubkey  crypto.PubKey `json:"pubkey",omitempty`
}

type MsgDeleteVpnUser struct {
	address sdk.AccAddress `json:"address", omitempty`
}
type MsgDeleteMasterNode struct {
	address sdk.AccAddress `json:"address", omitempty`
}
type MsgPayVpnService struct {
	Coins        string `json:"coins", omitempty`
	From         string `json:"address", omitempty`
	Vpnaddr      string `json:"vaddress", omitempty`
	Localaccount string `json:"account"`
	//	Password     string `json:"password"`
	ChainID  string `json:"chain-id"`
	Gas      int64  `json:"gas"`
	Sequence int64  `json:"sequence"`
}

type MsgSigntoVpn struct {
	coins     sdk.Coin          `json:"coins", omitempty`
	address   sdk.AccAddress    `json:"address", omitempty`
	sessionid int64             `json:"session_id", omitempty`
	signature auth.StdSignature `json:"signature", omitempty`
	from      sdk.AccAddress    `json:"from", omitempty`
}

type MsgGetVpnPayment struct {
	Coins        string `json:"coin"`
	Sessionid    string `json:"session-id"`
	Counter      int64  `json:"counter"`
	ChainID      string `json:"chain-id"`
	Localaccount string `json:"account"`
	Gas          int64  `json:"gas"`
	IsFinal      bool   `json:"isfinal"`

	//	Signature    string `json:"signature"`
}

type MsgRefund struct {
	Name      string `json:"name"`
	Sessionid string `json:"session_id", omitempty`
	ChainID   string `json:"chain-id"`
	Gas       int64  `json:"gas`
}

// client Signature :

type ClientSignature struct {
	Coins        string `json:"coin"`
	Sessionid    string `json:"session-id"`
	Counter      int64  `json:"counter"`
	isFinal      bool   `json:"isfinal"`
	Localaccount string `json:"account"`
	Password     string `json:"password"`
}

func ServiceRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec) {

	r.HandleFunc(
		"/registervpn",
		registervpnHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/register_master_node",
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
	r.HandleFunc(
		"/payvpn",
		PayVpnServiceHandlerFn(ctx, cdc),
	).Methods("POST")
	r.HandleFunc(
		"/sendsign",
		SendSignHandlerFn(ctx, cdc),
	).Methods("POST")
	r.HandleFunc(
		"/getvpnpayment",
		GetVpnPaymentHandlerFn(ctx, cdc),
	).Methods("POST")

}
func registervpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var a int64
		msg := MsgRegisterVpnService{}
		body, err := ioutill.ReadAll(r.Body)
		w.Write(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid  Msg Unmarshal function Request"))
			return
		} else {
			w.Write([]byte(" Request"))
		}

		if !validateIp(msg.Ip) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("  invalid Ip address."))
			return

		}
		if reflect.TypeOf(msg.Ppgb) != reflect.TypeOf(a) || msg.Ppgb < 0 || msg.Ppgb > 100000 {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid amount of price per Gb"))
			return
		}
		if reflect.TypeOf(msg.Netspeed) != reflect.TypeOf(a) || msg.Netspeed < 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid details"))
			return
		}
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithGas(msg.Gas)
		ctx = ctx.WithFromAddressName(msg.Localaccount)
		addr, err := ctx.GetFromAddress()
		if err != nil {
			panic(err)
		}
		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		//ctx=ctx.WithAccountNumber(msg.AccountNumber)
		msg1 := sentinel.NewMsgRegisterVpnService(addr, msg.Ip, msg.Ppgb, msg.Netspeed, msg.Location)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {

			panic(err)
			return
		}

	}
	return nil
}
func registermasterdHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := MsgRegisterMasterNode{}
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}

		json.Unmarshal(body, &msg)
		ctx = ctx.WithFromAddressName(msg.Name)
		ctx = ctx.WithGas(msg.Gas)
		addr, err := ctx.GetFromAddress()
		if err != nil {
			panic(err)
		}
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithGas(msg.Gas)
		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		//ctx=ctx.WithAccountNumber(msg.AccountNumber)
		msg1 := sentinel.NewMsgRegisterMasterNode(addr)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {

			panic(err)
			return
		}

	}
	return nil
}

func deleteMasterHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var msg MsgDeleteMasterNode
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid address."))
			return
		}
		msg1 := sentinel.NewMsgDeleteMasterNode(msg.address)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {

			panic(err)
			return
		}
	}
	return nil
}

func RefundHandleFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := MsgRefund{}
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		json.Unmarshal(body, &msg)
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithFromAddressName(msg.Name)
		ctx = ctx.WithGas(msg.Gas)
		addr, err := ctx.GetFromAddress()
		if err != nil {
			panic(err)
		}
		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		log.WriteLog("session id from client" + msg.Sessionid)
		msg1 := sentinel.NewMsgRefund(addr, []byte(msg.Sessionid))
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {
			panic(err)
		}
	}
}

func deleteVpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var msg MsgDeleteVpnUser
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid address."))
			return
		}

		// msg:= sentinel.MsgDeleteVpnUser{addres}
		msg1 := sentinel.NewMsgDeleteVpnUser(msg.address)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
	}
	return nil
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

func PayVpnServiceHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var msg MsgRegisterVpnService
		msg := MsgPayVpnService{}
		body, err := ioutill.ReadAll(r.Body)
		w.Write(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Invalid  Msg Unmarshal function Request"))
			return
		}
		if msg.Coins == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid address."))
			return
		}
		if msg.Vpnaddr == "" {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid vpn address"))
			return
		}
		vaddr, err := sdk.AccAddressFromBech32(msg.Vpnaddr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		if msg.From == "" {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("  invalid address"))
			return
		}
		addr, err := sdk.AccAddressFromBech32(msg.From)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		coins, err := sdk.ParseCoin(msg.Coins)
		if err != nil {
			panic(err)
		}

		ctx = ctx.WithFromAddressName(msg.Localaccount)
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithGas(msg.Gas)
		ctx = ctx.WithSequence(msg.Sequence)
		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		//Time := time.Now()
		msg1 := sentinel.NewMsgPayVpnService(coins, vaddr, addr)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {

			panic(err)
			return
		}

	}
	return nil
}

//To create client signature....... This is not a transaction......

func SendSignHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := ClientSignature{}
		var err error
		body, err := ioutill.ReadAll(r.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Invalid  Msg Unmarshal function Request"))
			return
		}
		coins, err := sdk.ParseCoin(msg.Coins)
		if err != nil {
			panic(err)
		}
		bz := senttype.ClientStdSignBytes(coins, []byte(msg.Sessionid), msg.Counter, msg.isFinal)

		keybase, err := keys.GetKeyBase()
		if err != nil {
			panic(err)
		}

		sig, pubkey, err := keybase.Sign(msg.Localaccount, msg.Password, bz)
		if err != nil {
			panic(err)
		}
		Signature.Signature = sig
		Signature.Pubkey = pubkey
		//signature := types.NewSignature(pubkey, sig)
		val := senttype.NewClientSignature(coins, []byte(msg.Sessionid), msg.Counter, pubkey, sig, msg.isFinal)

		address := val.Signature.Pubkey.Address().String()
		log.WriteLog("address of signed " + address)
		data, err := json.Marshal(val)
		if err != nil {
			panic(err)
		}
		log.WriteLog(string(data))
	}
	return nil
}

var Signature struct {
	Pubkey    crypto.PubKey
	Signature crypto.Signature
}

func GetVpnPaymentHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		msg := MsgGetVpnPayment{}
		body, err := ioutill.ReadAll(r.Body)
		w.Write(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Invalid  Msg Unmarshal function Request"))
			return
		}
		log.WriteLog("coins" + msg.Coins + "sessionid" + msg.Sessionid)
		if msg.Coins == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid address."))
			return
		}
		if msg.Sessionid == "" {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" Session Id is wrong"))
			return
		}
		log.WriteLog("coins" + msg.Coins + "sessionid" + msg.Sessionid)
		if msg.Counter < 0 {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Counter"))
			return
		}
		coins, err := sdk.ParseCoin(msg.Coins)
		if err != nil {
			panic(err)
		}

		ctx = ctx.WithFromAddressName(msg.Localaccount)
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithGas(msg.Gas)
		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		addr, err := ctx.GetFromAddress()
		if err != nil {
			panic(err)
		}
		//Time := time.Now()
		//	msg1 := sentinel.NewMsgGetVpnPayment(coins, []byte(msg.Sessionid), msg.Counter, addr,Signature, senttype.ClientSignature.IsFinal, caddr)
		msg1 := sentinel.NewMsgGetVpnPayment(coins, []byte(msg.Sessionid), msg.Counter, addr, Signature.Signature, Signature.Pubkey, msg.IsFinal)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {
			panic(err)
			return
		}
	}
	return nil
}
