package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Registervpn struct {
	ip       string
	netspeed string
	ppgb     string
	coins    sdk.Coin // PLEASE CHECK THIS PROPERLY, remove this
	// TODO :/// MUST ADD LOCATION PARAMETER of type string
}
