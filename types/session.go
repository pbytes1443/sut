package types

import (
	"crypto/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type Session struct {
	TotalLockedCoins   sdk.Coin
	CurrentLockedCoins sdk.Coin
	//UnlockedCoins      sdk.Coin
	//	Counter int32
	//timestamp          time.Time
	VpnAddr sdk.AccAddress
}
type SessionMap map[string]Session

func GetNewSessionId() []byte {
	Sessionid := make([]byte, 20)
	rand.Read(Sessionid)
	//return []byte("rgukt123")
	return Sessionid
}
func GetNewSessionMap(publickey crypto.PubKey, coins sdk.Coin, vpnAddr sdk.AccAddress) SessionMap {
	Sess := make(SessionMap)
	Sess[publickey.Address().String()] = Session{
		TotalLockedCoins:   coins,
		CurrentLockedCoins: coins,
		VpnAddr:            vpnAddr,
	}

	return Sess
}
