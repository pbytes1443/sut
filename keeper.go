package sentinel

import (
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
	keeper := Keeper{
		sentStoreKey: key,
		cdc:          cdc,
		coinKeeper:   ck,
		codespace:    codespace,
		account:      am,
	}
	return keeper
}

func (keeper Keeper) RegisterVpnService(ctx sdk.Context, msg MsgRegisterVpnService) (string, sdk.Error) {
	sentKey := msg.From.String()
	store := ctx.KVStore(keeper.sentStoreKey)

	vpnreg := senttype.NewVpnRegister(msg.Ip, msg.Location, msg.Ppgb, msg.Netspeed)
	bz, _ := keeper.cdc.MarshalBinary(vpnreg)
	store.Set([]byte(sentKey), bz)
	return "", nil
}

func (keeper Keeper) RegisterMasterNode(ctx sdk.Context, msg MsgRegisterMasterNode) (sdk.AccAddress, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	address := msg.Address
	bz, _ := keeper.cdc.MarshalBinary(address)
	store.Set([]byte(msg.Address.String()), bz)
	return msg.Address, nil

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

	var err error
	cpublicKey := keeper.account.GetAccount(ctx, msg.From).GetPubKey()
	sentKey := senttype.GetNewSessionId()
	vpnpub, err:=keeper.account.GetPubKey(ctx,msg.Vpnaddr)
	if err!=nil{
		sdk.ErrCommon("Vpn pubkey failed").Result()
	}
	session := senttype.GetNewSessionMap(msg.Coins,vpnpub,cpublicKey)

	store := ctx.KVStore(keeper.sentStoreKey)

	log.WriteLog("Seesion Id" + string(sentKey))

	bz,err:= keeper.cdc.MarshalBinary(session)
	if err != nil {
		sdk.ErrCommon("Marshal of session struct is failed").Result()
	}

	log.WriteLog("bz map bytes SessionMap Obj" + string(bz))
	store.Set(sentKey, bz)
	keeper.coinKeeper.SubtractCoins(ctx, msg.From, sdk.Coins{msg.Coins})
	return string(sentKey[:]), nil
}

func (keeper Keeper) GetVpnPayment(ctx sdk.Context, msg MsgGetVpnPayment) ([]byte, sdk.Error) {

	var clientSession senttype.Session
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.Sessionid)
	log.WriteLog("unmarshal string ......" + string(x))
	err := keeper.cdc.UnmarshalBinary(x, &clientSession)
	if err != nil {
		panic(err)
	}
	ClientPubkey:=clientSession.CPubKey

	signBytes := senttype.ClientStdSignBytes(msg.Coins, msg.Sessionid, msg.Counter, msg.IsFinal) //errr

	if !ClientPubkey.VerifyBytes(signBytes, msg.Signature) {
		return nil, sdk.ErrUnauthorized("signature from the keeper.go verification failed")
	}
	clientSessionData := clientSession
	if clientSessionData.CurrentLockedCoins.IsPositive() && clientSessionData.Counter <= msg.Counter {
		CoinsToAdd := msg.Coins
		clientSessionData.CurrentLockedCoins = clientSessionData.CurrentLockedCoins.Minus(CoinsToAdd)
		clientSessionData.Counter = msg.Counter
		VpnAddr:=sdk.AccAddress(clientSessionData.VpnPubKey.Address())
		keeper.coinKeeper.AddCoins(ctx, VpnAddr, sdk.Coins{CoinsToAdd})

		sentKey := msg.Sessionid
		bz, _ := keeper.cdc.MarshalBinary(clientSessionData)
		store.Set(sentKey, bz)
		return sentKey, nil
	}

	return nil, nil
}

func (keeper Keeper) RefundBal(ctx sdk.Context, msg MsgRefund) (sdk.AccAddress, sdk.Error) {

	var err error
	var clientSession senttype.Session
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.Sessionid)
	err = keeper.cdc.UnmarshalBinary(x, &clientSession)
	if err != nil {
		log.WriteLog("unmarshal error of clientSession")
	}
	ctime := time.Now().UnixNano()
	if int64(math.Abs(float64(ctime))) >= 86400000  && clientSession.CurrentLockedCoins.IsPositive() && !clientSession.CurrentLockedCoins.IsZero(){

		keeper.coinKeeper.AddCoins(ctx, msg.From, sdk.Coins{clientSession.CurrentLockedCoins})
		store.Delete(msg.Sessionid)
		return msg.From, nil
	} else {
		return nil, sdk.ErrCommon("time is less than 24 hours  or the balance is negative or equal to zero")
	}

	return nil, nil

}
