package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
)

type MsgQueryRegisteredVpnService struct {
	address sdk.Address `json:"address",omitempty`
}
type MsgQueryFromMasterNode struct {
	address sdk.Address `json:"address",omitempty`
}

func QueryRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec) {
	r.HandleFunc(
		"/query_vpn/{address}",
		queryvpnHandlerFn(ctx, cdc),
	).Methods("POST")

	r.HandleFunc(
		"/query_master_node/{address}",
		querymasterHandlerFn(ctx, cdc),
	).Methods("POST")

}

func queryvpnHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var msg MsgQueryRegisteredVpnService
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid address."))
			return
		}
		// read parameters
		// vars := mux.Vars(r)
		// address := vars["address"]
		// var ctx1 sdk.context
		// var k sent.Keeper

		// msg := sentinel.MsgQueryRegisteredVpnService{address}
		res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)

	}
}

func querymasterHandlerFn(ctx context.CoreContext, cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read parametersvar msg MsgQueryRegisteredVpnService
		var msg MsgQueryFromMasterNode
		var err error
		body, err := ioutill.ReadAll(r.Body)
		json.Unmarshal(body, &msg)
		if msg.address == nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You have to entered invalid address."))
			return
		}

		// vars := mux.Vars(r)
		// address := vars["address"]
		// msg := sentinel.MsgQueryFromMasterNode{address}
		res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
		if err != nil {
			return err
		}
	}
}
