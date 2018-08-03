package types

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	log "github.com/logger"
)

var pool = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Session struct {
	TotalLockedCoins   sdk.Coin
	CurrentLockedCoins sdk.Coin
	Counter            int64
	Timestamp          int64
	VpnAddr            sdk.AccAddress
}
type SessionMap map[string]Session

func GetNewSessionId() []byte {

	bytes := make([]byte, 20)
	for i := 0; i < 20; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	// fmt.Println(bytes)
	log.WriteLog("bytes in seesion id " + string(bytes[:]))
	return bytes

}
func GetNewSessionMap(publickey string, coins sdk.Coin, vpnAddr sdk.AccAddress) SessionMap {
	Sess := make(SessionMap)
	ti := time.Now().UnixNano()
	Sess[publickey] = Session{
		TotalLockedCoins:   coins,
		CurrentLockedCoins: coins,
		VpnAddr:            vpnAddr,
		Timestamp:          ti,
	}

	return Sess
}
