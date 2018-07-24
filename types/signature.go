package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
	


)

type ClientSignature struct {
	coins     sdk.Coin
	sessionid [20]byte
	counter   int64
	signature  Signature
}
type Signature struct {
	crypto.PubKey    `json:"pub_key"` // optional
	crypto.Signature `json:"signature"`
}

type StdSig struct {
	coins sdk.Coin
	sessionid [20]byte
	counter   int64
}

func ClientStdSignBytes(coins sdk.Coin,sessionid [20]byte,counter int64) []byte {
	bz, err := msgCdc.MarshalJSON(StdSig{
		coins :     coins,
		sessionid : sessionid,
		counter : counter,
	})
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}