package types

import (
	 sdk "github.com/cosmos/cosmos-sdk/types"
	 "github.com/tendermint/tendermint/crypto"
)


func GetNewSessionId() []byte {
	sessionId := crypto.CRandBytes(20)
	return sessionId
}
func  GetNewSessionMap(publickey crypto.PubKeySecp256k1,
	coins sdk.Coin,timestamp int64,vpnAddr sdk.AccAddress ) SessionMap  {
		//
		Key := string(publickey)
		sessionmap :=  map[publickey]session{
			initialLockedCoins:coins,
			currentLockedCoins:coins,
			timestamp:timestamp,
			vpnAddr:vpnAddr,
		} 
		return sessionmap
}

type session struct{
    initialLockedCoins    sdk.Coin
	currentLockedCoins sdk.Coin
    counter int32
	timestamp int64
	vpnAddr sdk.AccAddress
}

type SessionMap  map[string]session



