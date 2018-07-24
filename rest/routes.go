package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	//"flags"
)

func RegisterRoutes(ctx context.CoreContext, r *mux.Router, cdc *wire.Codec) {

	ServiceRoutes(ctx, r, cdc)
	QueryRoutes(ctx, r, cdc)

}

//cosmos-sdk/stake/query......
