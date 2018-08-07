package sentinel

import (
	//"encoding/json"
	"math"

	//	"fmt"

	senttype "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/logger"
	"github.com/tendermint/tendermint/crypto"
	//	"strconv"
	//	"strings"
	"time"
)

type PubKeyEd25519 [32]byte

type Keeper struct {
	sentStoreKey sdk.StoreKey
	coinKeeper   bank.Keeper
	cdc          *wire.Codec

	codespace sdk.CodespaceType
	account   auth.AccountMapper
}

type Sign struct {
	coin    sdk.Coin
	vpnaddr sdk.AccAddress
	counter int64
	hash    string
	sign    crypto.PubKeySecp256k1
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, ck bank.Keeper, am auth.AccountMapper, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		sentStoreKey: key,
		cdc:          cdc,
		coinKeeper:   ck,
		codespace:    codespace,
		account:      am,
	}

}

func (keeper Keeper) RegisterVpnService(ctx sdk.Context, msg MsgRegisterVpnService) (sdk.AccAddress, sdk.Error) {
	sentKey := msg.From.String()
	store := ctx.KVStore(keeper.sentStoreKey)
	address := store.Get([]byte(sentKey))
	if address == nil {
		vpnreg := senttype.NewVpnRegister(msg.Ip, msg.Location, msg.Ppgb, msg.Netspeed)
		bz, _ := keeper.cdc.MarshalBinary(vpnreg)
		store.Set([]byte(sentKey), bz)
		return msg.From, nil

	}
	return nil, sdk.ErrCommon("Address already Registered as VPN node")
}

func (keeper Keeper) RegisterMasterNode(ctx sdk.Context, msg MsgRegisterMasterNode) (sdk.AccAddress, sdk.Error) {
	sentkey := msg.Address.String()
	store := ctx.KVStore(keeper.sentStoreKey)
	address := store.Get([]byte(sentkey))
	if address == nil {
		address := msg.Address
		bz, _ := keeper.cdc.MarshalBinaryBare(address)
		store.Set([]byte(msg.Address.String()), bz)
		return msg.Address, nil

	}
	return nil, sdk.ErrCommon("Address already registered as MasterNode")
}

func (keeper Keeper) StoreKey() sdk.StoreKey {
	return keeper.sentStoreKey
}

//func (keeper Keeper) QueryStore(addr sdk.AccAddress,Name sdk.StoreKey) []byte {
////	//var msg MsgRegisterVpnService
////	store := sdk.Context.KVStore(Name)
////	data := store.Get([]byte(addr))
////	//err := keeper.cdc.UnmarshalBinary(data,&msg)
////	return data
//}
func (keeper Keeper) DeleteVpnService(ctx sdk.Context, msg MsgDeleteVpnUser) (sdk.AccAddress, sdk.Error) {

	store := ctx.KVStore(keeper.sentStoreKey)
	db := store.Get([]byte(msg.Vaddr.String()))
	if db == nil {
		return nil, sdk.ErrCommon("Account is not exist")
	}
	store.Delete([]byte(msg.Vaddr.String()))
	return msg.Vaddr, nil
}
func (keeper Keeper) DeleteMasterNode(ctx sdk.Context, msg MsgDeleteMasterNode) (sdk.AccAddress, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	db := store.Get([]byte(msg.Maddr.String()))
	if db == nil {
		return nil, sdk.ErrCommon("Account is not exist")
	}
	store.Delete([]byte(msg.Maddr.String()))
	return msg.Maddr, nil
}

func (keeper Keeper) PayVpnService(ctx sdk.Context, msg MsgPayVpnService) (string, sdk.Error) {

	var err error
	cpublicKey := keeper.account.GetAccount(ctx, msg.From).GetPubKey()
	sentKey := senttype.GetNewSessionId()
	vpnpub, err := keeper.account.GetPubKey(ctx, msg.Vpnaddr)
	if err != nil {
		sdk.ErrCommon("Vpn pubkey failed").Result()
	}
	session := senttype.GetNewSessionMap(msg.Coins, vpnpub, cpublicKey)

	store := ctx.KVStore(keeper.sentStoreKey)

	//log.WriteLog("Seesion Id" + string(sentKey))
	data := store.Get([]byte(msg.Vpnaddr.String()))
	if data == nil {
		return "", sdk.ErrCommon("VPN address is not registered")
	}
	bz, err := keeper.cdc.MarshalBinary(session)
	if err != nil {
		sdk.ErrCommon("Marshal of session struct is failed").Result()
	}

	//	log.WriteLog("bz map bytes SessionMap Obj" + string(bz))
	store.Set(sentKey, bz)

	// keeper.cdc.UnmarshalBinary(bz, &session)
	// a, err := json.Marshal(session)
	// if err != nil {

	// }
	// log.WriteLog(string(a))
	keeper.coinKeeper.SubtractCoins(ctx, msg.From, sdk.Coins{msg.Coins})
	return string(sentKey[:]), nil
}
func (keeper Keeper) RefundBal(ctx sdk.Context, msg MsgRefund) (sdk.AccAddress, sdk.Error) {

	var err error
	var clientSession senttype.Session
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.Sessionid)
	
	log.WriteLog("x object.........." + string(x))
	err = keeper.cdc.UnmarshalBinary(x, &clientSession)
	if err != nil {
		log.WriteLog("unmarshal error of clientSession")
	}
	caddr := sdk.AccAddress(clientSession.CPubKey.Address())
	if msg.From.String() != caddr.String() {
		return nil, sdk.ErrCommon("Address is not associated with this Session")
	}
	ctime := time.Now().UnixNano()
	if int64(math.Abs(float64(ctime))) >= 86400000 && clientSession.CurrentLockedCoins.IsPositive() && !clientSession.CurrentLockedCoins.IsZero() {

		keeper.coinKeeper.AddCoins(ctx, msg.From, sdk.Coins{clientSession.CurrentLockedCoins})
		store.Delete(msg.Sessionid)
		return msg.From, nil
	} else {
		return nil, sdk.ErrCommon("time is less than 24 hours  or the balance is negative or equal to zero")
	}

	return nil, nil

}
func (keeper Keeper) GetVpnPayment(ctx sdk.Context, msg MsgGetVpnPayment) ([]byte, sdk.Error) {

	var clientSession senttype.Session
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.Sessionid)
	log.WriteLog(string(msg.Sessionid[:]))
	log.WriteLog("unmarshal string ......" + string(x))
	err := keeper.cdc.UnmarshalBinary(x, &clientSession)
	if err != nil {
		panic(err)
	}
	ClientPubkey := clientSession.CPubKey

	signBytes := senttype.ClientStdSignBytes(msg.Coins, msg.Sessionid, msg.Counter, msg.IsFinal) //errr

	if !ClientPubkey.VerifyBytes(signBytes, msg.Signature) {
		return nil, sdk.ErrUnauthorized("signature from the keeper.go verification failed")
	}
	clientSessionData := clientSession
	if clientSessionData.CurrentLockedCoins.IsPositive() && clientSessionData.Counter <= msg.Counter {
		CoinsToAdd := msg.Coins
		clientSessionData.CurrentLockedCoins = clientSessionData.CurrentLockedCoins.Minus(CoinsToAdd)
		clientSessionData.Counter = msg.Counter
		VpnAddr := sdk.AccAddress(clientSessionData.VpnPubKey.Address())
		keeper.coinKeeper.AddCoins(ctx, VpnAddr, sdk.Coins{CoinsToAdd})

		sentKey := msg.Sessionid
		bz, err := keeper.cdc.MarshalBinary(clientSessionData)
		if err != nil {
			return nil, sdk.ErrCommon("unmarshalling error")
		}
		store.Set(sentKey, bz)
		return sentKey, nil
	}

	return x, nil
}

func (keeper Keeper) NewMsgDecoder(acc []byte) (senttype.Registervpn, sdk.Error) {

	msg := senttype.Registervpn{}
	// acct := new(auth.BaseAccount)
	//	var err error
	log.WriteLog(string(acc[:]))
	err := keeper.cdc.UnmarshalBinary(acc, &msg)
	if err != nil {
		panic(err)
	}
	return msg, sdk.ErrCommon("account unmarshal failed")

}

func (keeper Keeper) GetsentStore(ctx sdk.Context, msg MsgRegisterMasterNode) (sdk.KVStore, sdk.AccAddress) {
	return ctx.KVStore(keeper.sentStoreKey), addr1
}
