package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type ClientSignature struct {
	coins     sdk.Coin
	sessionid int64
	counter   int64
	timestamp
	signature auth.Signature
}
