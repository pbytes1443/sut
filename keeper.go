package sentinel

import (
	"encoding/json"

	//	"fmt"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	keys "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/x/auth"
	crypto "github.com/tendermint/tendermint/crypto"
	dbm "github.com/tendermint/tendermint/libs/db"
	senttype "github.com/cosmos/cosmos-sdk/examples/sut/types"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	//	"strconv"
	//	"strings"

	"math/rand"
	"time"
)


type                                                                                                       map[crypto.PublicKey][]MsgSigntoVpn
type vpndb struct {
	db dbm.DB
}
type PubKeyEd25519 [32]byte

type Keeper struct {
	sentStoreKey sdk.StoreKey
	coinKeeper   bank.Keeper
	cdc          *wire.Codec

	// codespace
	codespace sdk.CodespaceType
}
// type PayVpnInit struct {
// 	coins   sdk.Coin
// 	pubkey  *crypto.PubKey
// 	session map[pubkey]lock
// }
// type lock struct {
// 	lock  sdk.Coin
// 	total sdk.Coin
// 	unlock sdk.Coin
// 	coins_to add sdk.Coin
// 	time time.Time()
// }
type Sign struct {
	coin    sdk.Coin
	vpnaddr sdk.Address
	counter int64
	hash    string
	sign    crypto.SignatureEd25519
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, ck bank.Keeper, codespace sdk.CodespaceType) Keeper {
	keeper := Keeper{
		sentStoreKey: key,
		cdc:          cdc,
		coinKeeper:   ck,
		codespace:    codespace,            ////learn WHAT THIS DOES
	}
	return keeper
}

func (keeper Keeper) RegisterVpnService(ctx sdk.Context, msg MsgRegisterVpnService) (string, sdk.Error) {                                      
	sentKey := msg.address 
	store := ctx.KVStore(keeper.sentStoreKey)
	var vpnreg registervpn                    
	bz, _ := keeper.cdc.MarshalBinary(vpnreg)       
	store.Set(sentKey, bz)
	return "", nil
}


func (keeper Keeper) QueryRegisteredVpnService(ctx sdk.Context, msg MsgQueryRegisteredVpnService) (regvpn, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	var vpnreg registervpn           
	bz := store.Get(msg.address)    
	if bz != nil {                                        
		senttype.ErrCommon("Address is not valid").Result()
	}
	keeper.cdc.UnmarshalBinary(bz, &prov)
	return prov, nil
}

func (keeper Keeper) RegisterMasterNode(ctx sdk.Context, msg MsgRegisterMasterNode) (sdk.Address, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	address := msg.pubkey.Address()                      
	bz, _ := keeper.cdc.MarshalBinary(address)
	store.Set(msg.address, bz)
	return msg.address, nil

}

func (keeper Keeper) QueryFromRegisterMasterNode(ctx sdk.Context, msg MsgQueryFromMasterNode) (registervpn, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	var vpnreg registervpn
	bz := store.Get(msg.address)
	if bz != nil {
		keeper.cdc.UnmarshalBinary(bz, &vpnreg)
		return vpnreg, nil
	}
	return registervpn{}, nil
}

func (keeper Keeper) DeleteVpnService(ctx sdk.Context, msg MsgDeleteVpnUser) (sdk.AccAddress, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	store.Delete(msg.address)
	return msg.address, nil
}
func (keeper Keeper) DeleteMasterNode(ctx sdk.Context, msg MsgDeleteMasterNode) (sdk.AccAddress, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	store.Delete(msg.address)
	return msg.address, nil
}
func (keeper Keeper) PayVpnService(ctx sdk.Context, msg MsgPayVpnService) (string, sdk.Error) {
	var sessionMap types.SessionMap
	publicKey := keeper.coinKeeper.am.GetAccount(msg.address).PubKey
	sessionMap := types.GetNewSessionMap(publicKey
	msg.coins,msg.timestamp,msg.vpnAddr)            
	store:=ctx.KVStore(keeper.sentStoreKey)
	sentKey:= types.GetSessionId()
	bz, _ := keeper.cdc.MarshalBinary(sentKey)       //PLEASE USE SMAE VARIABLE NAMES  FOR FUNCTION parameter
	store.Set(sentKey, bz)
	keeper.bankKeeper.SubtractCoins(ctx,msg.address,msg.coins)
	return string(sessionid), nil
}

func (keeper Keeper) sendSigntoVpn(ctx sdk.Context, msg MsgSigntoVpn) ( sdk.Error) {
	// var k keys.Keybase
	// k, err = Keybase.GetKeyBase()
	// sig, pub,err := k.Sign(name,passphrase,msg)
	// msg.sign:sig
	vpnsignstore[pub][msg]
	return nil
}
func (keeper Keeper) GetVpnPayment(ctx sdk.Context, msg MsgGetVpnPayment) ( sdk.Error) {           // TODO Change the function name to GetVpnPayment
	
	var clientSession sessionMap
	signature := msg.clientSig.signature
	pubKey := signature.PubKey
	sig:=signature.Signature
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.sessionid)
	err := keeper.cdc.UnmarshalBinary(x,&clientSession)
	if err != nil {
		return nil, err
	}
	ClientPubkey := reflect.ValueOf(clientSession).MapKeys()[0]
	 if ClientPubkey != pubkey {
		senttype.ErrCommon("Invalid Public key").Result()
	}
    signBytes := ClientStdSignBytes(signature.coins,signature.sessioid,signature.counter)
	if !ClientPubkey.VerifyBytes(signBytes, sig.Signature) {
		return nil, sdk.ErrUnauthorized("signature verification failed").Result()
	}
	clientSessionData := clientSession[ClientPubkey]
	if    clientsig := msg.clientsig;  clientSessionData.currentLockedCoins > 0 &&  clientSessionData.counter <= clientSig.counter {
			CoinsToAdd:=msg.clientSig.coins.Minus(UnlockedCoins)
			clientSessionData.currentLockedCoins=clientSessionData.currentLockedCoins.Minus(CoinsToAdd)
			clientSessionData.UnlockedCoins=clientSessionData.UnlockedCoins.Plus(clientSessionData.totalLockedCoins.Minus(clientSessionData.currentLockedCoins))
			clientSessionData.counter = msg.clientSig.counter
			keeper.bankKeeper.AddCoins(ctx,clientSessionData.vpnaddr,coins)

	}
	sentkey := clientSig.sessionid
	bz,_:= keeper.cdc.MarshalBinary(clientSessionData)
	store.Set(sentKey,bz)
	return sentKey
}
func (keeper Keeper) RefundBal(ctx sdk.Context, msg MsgRefund) (sdk.Error){
	
	var t time.Time
	var clientSession sessionMap
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.sessionid)
	err := keeper.cdc.UnmarshalBinary(x,&clientSession)
	if err != nil {
		return nil, err
	}
	pubkey := keeper.coinKeeper.am.GetPubkey(ctx,msg.from)
    clientSessionData := clientSession[pubkey]
	tm :=clientSessionData.timestamp
	diff:=time.Now().Sub(tm)
	//
	if diff.Hours()>=24{
	coins=clientSessionData.currentLockedCoins
	keeper.bankKeeper.AddCoins(ctx,msg.address,coins)
	store.Delete(msg.sessionid)
	}
	
	return  nil

}