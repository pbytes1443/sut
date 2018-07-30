package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// "github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/tendermint/tendermint/crypto"
)

type ClientSignature struct {
	Coins     sdk.Coin
	Sessionid []byte
	Counter   int64
	Signature Signature
}
type Signature struct {
	Pubkey    crypto.PubKey    `json:"pub_key"` // optional
	Signature crypto.Signature `json:"signature"`
}

type StdSig struct {
	coins     sdk.Coin
	sessionid []byte
	counter   int64
}

func ClientStdSignBytes(coins sdk.Coin, sessionid []byte, counter int64) []byte {
	bz, err := json.Marshal(StdSig{
		coins:     coins,
		sessionid: sessionid,
		counter:   counter,
	})
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(bz)
}
