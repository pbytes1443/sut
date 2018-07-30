package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

func GetNewSessionId() []byte {
	Sessionid := crypto.CRandBytes(20)
	return Sessionid
}
func GetNewSessionMap(publickey crypto.PubKey, coins sdk.Coin, timestamp time.Time, vpnAddr sdk.AccAddress) SessionMap {
	var sessionmap SessionMap
	sessionmap[publickey] = session{ // may be changed the intiallization using make
		totalLockedCoins:   coins,
		currentLockedCoins: coins,
		UnlockedCoins:      sdk.NewCoin("", 0),
		counter:            0,
		timestamp:          time.Time{},
		vpnAddr:            vpnAddr,
	}
	return sessionmap
}

type session struct {
	totalLockedCoins   sdk.Coin
	currentLockedCoins sdk.Coin
	UnlockedCoins      sdk.Coin
	counter            int32
	timestamp          time.Time
	vpnAddr            sdk.AccAddress
}

type SessionMap map[crypto.PubKey]session
