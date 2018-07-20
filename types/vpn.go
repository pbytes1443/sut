package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type registervpn struct {
	ip       string
	netspeed string
	ppgb     string
	coins    sdk.Coin
}
