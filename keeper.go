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
	"github.com/cosmos/cosmos-sdk/examples/sut/types"
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
type PayVpnInit struct {
	coins   sdk.Coin
	pubkey  *crypto.PubKey
	session map[pubkey]lock
}
type lock struct {
	lock  sdk.Coin
	total sdk.Coin
	time time.Time()
}
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
	var bal sdk.Coin                         /// TODO : PLEASE CHANGE BAL to BALANCE WHEREVER IT's applicable
	bal=0                                            
	sentKey := msg.address      /*
	
	TODO: check the type of sentkey as sdk.ACCaddress
			   stdTx, ok := tx.(StdTx)
		if !ok {
			return ctx, sdk.ErrInternal("tx must be StdTx").Result(), true
		}
*/
	store := ctx.KVStore(keeper.sentStoreKey)
	var p regvpn                      /// This regvpn is of type registervpn located under types package   use it as types.registervpn


	/*
	p.ip = msg.ip
	p.coins=bal //initially zero coins
	p.netspeed = msg.netspeed
	p.ppgb = msg.ppgb


	TODO the above is poor way of writing code instead write it as
	 p{ip:ip,
	 }

	*/
   
	bz, _ := keeper.cdc.MarshalBinary(p)       //PLEASE USE SMAE VARIABLE NAMES  FOR FUNCTION parameter
	store.Set(sentKey, bz)
	fmt.Println("Service provider register with this address", sentKey)           //PLEASE REMOVE THIS OR ELSE PRINT IN A STANDARD FORMAT which includes module and other related info
	return "", nil
}


func (keeper Keeper) QueryRegisteredVpnService(ctx sdk.Context, msg MsgQueryRegisteredVpnService) (regvpn, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	bz := store.Get(msg.address)      /// DONOT PASS DIRECTLY CHECK TYPE
	var prov regvpn                               // change variable name
	if bz != nil {                                        /// IT IS bz==nil then return error of type keynotfoundinkvstore

	}
	keeper.cdc.UnmarshalBinary(bz, &prov)
	return prov, nil
}

func (keeper Keeper) RegisterMasterNode(ctx sdk.Context, msg MsgRegisterMasterNode) (sdk.Address, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	pubKey := msg.pubkey.Address()                      //change the variable name to address
	//fmt.Println("public key " + )
	bz, _ := keeper.cdc.MarshalBinary(pubKey)"github.com/cosmos/cosmos-sdk/x/bank"
	store.Set(msg.address, bz)
	return msg.address, nil

}

func (keeper Keeper) QueryFromRegisterMasterNode(ctx sdk.Context, msg MsgQueryFromMasterNode) (regvpn, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	bz := store.Get(msg.address)"github.com/cosmos/cosmos-sdk/x/bank"
	var a regvpn
	if bz != nil {
		keeper.cdc.UnmarshalBinary(bz, &a)
		return a, nil
	}
	return regvpn{}, nil
}

func (keeper Keeper) DeleteVpnService(ctx sdk.Context, msg MsgDeleteVpnUser) (sdk.Address, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	store.Delete(msg.address)
	return msg.address, nil
}
func (keeper Keeper) DeleteMasterNode(ctx sdk.Context, msg MsgDeleteMasterNode) (sdk.Address, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	store.Delete(msg.address)
	return msg.address, nil
}
func (keeper Keeper) PayVpnService(ctx sdk.Context, msg MsgPayVpnService) (string, sdk.Error) {
	            /// TODO : do this session id more random
	var sessionMap types.SessionMap
	publicKey := keeper.coinKeeper.am.GetAccount(msg.address).PubKey
	sessionMap := types.GetNewSessionMap(publicKey
		msg.coins,msg.timestamp,msg.vpnAddr)             //to do define this structure in types along with lock struct
	// sessionidata[msg.pubkey].lock = msg.coins               
	// sessionidata[msg.pubkey].total = msg.coins 
	// sessiondata[msg.pubkey].time= ctx.BlockHeader().Time 
	// val,err:=json.Marshal(sessiondata)
	// vpndb.db.Set(sessionid, val)
	store:=ctx.KVStore(keeper.sentStoreKey)
	sentKey:= types.GetSessionId()
	bz, _ := keeper.cdc.MarshalBinary(p)       //PLEASE USE SMAE VARIABLE NAMES  FOR FUNCTION parameter
	store.Set(sentKey, bz)
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
	// var k keys.Keybase
	// k, err = Keybase.GetKeyBase()
	// sig, pub,err := k.Sign(name,passphrase,msg)
	// msg.sign: sig
	// var session session
	// data:=vpndb.db.Get([]byte(msg.clientSig.sessionid))
	// json.Unmarshal(data ,&session)
	
	// msg.clientSig.session[msg.sing.pubkey].lock-=msg.sing.coins  
	// bytsession,_:= keeper.cdc.MarshalBinary(session)
	// vpndb.db.Set([]byte(msg.sing.sessioid),bytsession)
	
	// sentKey := msg.address
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
		// return Sdk.errors 
	}
    signBytes := ClientStdSignBytes(signature.coins,signature.sessioid,signature.counter)
	if !ClientPubkey.VerifyBytes(signBytes, sig.Signature) {
		return nil, sdk.ErrUnauthorized("signature verification failed").Result()
	}
	clientSessionData := clientSession[ClientPubkey]
	if    clientsig := msg.clientsig;  clientSessionData.currentLockedCoins > 0 &&  clientSessionData.counter <= clientSig.counter {
			coins := clientSessionData.initialLockedCoins - clientSessionData.currentLockedCoins
			clientSessionData.currentLockedCoins := clientSessionData.currentLockedCoins - coins
			clientSessionData.counter = clientSig.counter
			keeper.bankKeeper.AddCoins(ctx,clientSessionData.vpnaddr,coins)
	}
	sentkey := clientSig.sessionid
	bz,_:= keeper.cdc.MarshalBinary(clientSessionData)
	store.Set(sentKey,bz)
	// vpndata:=store.Get(sentKey)
	// var vpn regvpn
	// json.Unmarshal(vpndata,&vpn)
	// vpn.coins=vpn.coins+ msg.sing.coins
	// bz, _ := keeper.cdc.MarshalBinary(vpn)
	// store.Set(sentKey, bz)
	return sentKey
}
func (keeper Keeper) RefundBal(ctx sdk.Context, msg MsgRefund) (sdk.Error){
	
	var t time.Time
	// var sessiondata session
	var clientSession sessionMap
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.sessionid)
	err := keeper.cdc.UnmarshalBinary(x,&clientSession)
	if err != nil {
		return nil, err
	}
	// data:=vpndb.db.Get(msg.sessioid)
	// json.Unmarshal(data,&sessiondata)
	pubkey := keeper.coinKeeper.am.GetAccount(msg.address).PubKey
    clientSessionData := clientSession[pubkey]
	tm :=clientSessionData.timestamp
	diff:=time.Now().Sub(tm)
	//
	if diff.Hours()>=24{
	coins=clientSessionData.currentLockedCoins
	keeper.bankKeeper.AddCoins(ctx,msg.address,coins)
	store.Delete(msg.sessionid)
	// coin:= 	sessionidata[msg.pubkey].lock 
	// sessionidata[msg.pubkey].lock = bal
	// sessionidata[msg.pubkey].total = msg.coins 
	// sessiondata[msg.pubkey].time= ctx.BlockHeader().Time
	// val,err:=json.Marshal(sessiondata)
	// vpndb.db.Set(msg.sessionid, val)
	}
	
	return  nil

}