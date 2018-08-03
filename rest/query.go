package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/examples/sentinel"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/gorilla/mux"
	log "github.com/logger"
)

const (
	storeName = "sentinel"
)

// type MsgQueryRegisteredVpnService struct {
// 	address sdk.AccAddress `json:"address",omitempty`
// }
// type MsgQueryFromMasterNode struct {
// 	address sdk.AccAddress `json:"address",omitempty`
// }
var ctx1 sdk.Context

func QueryRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec, keeper sentinel.Keeper) {
	r.HandleFunc(
		"/query_vpn/{address}",
		queryvpnHandlerFn(ctx, cdc, authcmd.GetAccountDecoder(cdc), keeper),
	).Methods("GET")

	r.HandleFunc(
		"/query_master_node/{address}",
		querymasterHandlerFn(ctx, authcmd.GetAccountDecoder(cdc), cdc),
	).Methods("GET")

}

func queryvpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec, decoder auth.AccountDecoder, keeper sentinel.Keeper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		address := vars["address"]
		if address == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid address."))
			return
		}
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		log.WriteLog("Account Store " + ctx.AccountStore)
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), ctx.AccountStore)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query account. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(fmt.Sprint("asdfghjf")))
			return

		}
		// decode the value
		account, err := decoder(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		// print out whole account
		output, err := cdc.MarshalJSON(account)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)

	}
	return nil
}

func querymasterHandlerFn(ctx context.CoreContext, decoder auth.AccountDecoder, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read parametersvar msg MsgQueryRegisteredVpnService
		vars := mux.Vars(r)
		address := vars["address"]
		if address == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(" entered invalid address."))
			return
		}
		addr, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		res, err := ctx.QueryStore(auth.AddressStoreKey(addr), storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't query account. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this account
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// decode the value
		account, err := decoder(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		output, err := cdc.MarshalJSON(account)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("couldn't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)

	}

	return nil
}
