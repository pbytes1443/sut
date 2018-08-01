package rest

import (
	"encoding/json"

	// "fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	ioutill "io/ioutil"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/examples/sentinel"
	"github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/gorilla/mux"
	crypto "github.com/tendermint/tendermint/crypto"
)

type MsgRegisterVpnService struct {
	Address      string `json:"address"`
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
	Address string `json:"address",omitempty`
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
	clientSig types.ClientSignature `json:"signature", omitempty`
	from      sdk.AccAddress        `json:"address", omitempty`
}

type MsgRefund struct {
	pubkey    crypto.PubKey `json:"pubkey", omitempty`
	sessionid int64         `json:"session_id", omitempty`
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

	// r.HandleFunc(
	// 	"/refund/",
	// 	RefundHandleFn(ctx, cdc),
	// ).Methods("POST")

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

}
func registervpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var msg MsgRegisterVpnService
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
		if msg.Address == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" invalid address."))
			return
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

		addr, err := sdk.AccAddressFromBech32(msg.Address)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithGas(msg.Gas)
		ctx = ctx.WithFromAddressName(msg.Localaccount)
		ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		//ctx=ctx.WithAccountNumber(msg.AccountNumber)
		msg1 := sentinel.NewMsgRegisterVpnService(addr, msg.Ip, msg.Ppgb, msg.Netspeed, msg.Location)
		err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		if err != nil {

			panic(err)
			return
		}
		//ctx = ctx.WithSequence(msg.Sequence)

		// txBytes, err := ctx.SignAndBuild(msg.Localaccount, msg.Password, []sdk.Msg{msg1}, cdc)
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// res, err := ctx.BroadcastTx(txBytes)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// output, err := json.MarshalIndent(res, "", "  ")
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// w.Write(output)

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
		//vars := mux.Vars(r)
		//strProposalID := vars[RestProposalID]
		//bechDepositerAddr := vars[RestDepositer]
		//addr1 := vars["address"]
		json.Unmarshal(body, &msg)
		addr, err := sdk.AccAddressFromBech32(msg.Address)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// if msg.address == "" {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	w.Write([]byte(" entered invalid address."))
		// 	returnrgukt123

		// }
		ctx = ctx.WithChainID(msg.ChainID)
		ctx = ctx.WithGas(msg.Gas)
		ctx = ctx.WithFromAddressName(msg.Name)
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

// func RefundHandleFn(ctx contex.CoreContext,cdc *wire.Codec) http.HandleFunc{
//  /*	return func(w http.ResponseWriter, r http.Request){

// 		var := mux.Vars(r)
// 		addres:= var["address"]
// 		pubkey:= var["public_key"]
// 		msg:= sender.MsgRefund{addres,pubkey}
// 		res, err :=ctx.EnsureSignBuildBroadcast(ctx.FromAddressName,msg,cdc)
// 	}*/
// }

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
		// ctx = ctx.WithChainID(msg.ChainID)
		// ctx = ctx.WithGas(msg.Gas)
		// ctx = ctx.WithFromAddressName(msg.Name)
		// ctx = ctx.WithDecoder(authcmd.GetAccountDecoder(cdc))
		// //ctx=ctx.WithAccountNumber(msg.AccountNumber)
		// msg1 := sentinel.NewMsgRegisterMasterNode(addr)
		// err = ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, []sdk.Msg{msg1}, cdc)
		// if err != nil {

		// 	panic(err)
		// 	return
		// }
		// txBytes, err := ctx.SignAndBuild(msg.Localaccount, msg.Password, []sdk.Msg{msg1}, cdc)
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// res, err := ctx.BroadcastTx(txBytes)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// output, err := json.MarshalIndent(res, "", "  ")
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(err.Error()))
		// 	return
		// }

		// w.Write(output)

	}
	return nil
}
