package sentinel

import (
	"encoding/json"
	"math"
	"time"

	//	"fmt"

	senttype "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	log "github.com/logger"
	crypto "github.com/tendermint/tendermint/crypto"
	//	"strconv"
	//	"strings"
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

	publicKey := keeper.account.GetAccount(ctx, msg.From).GetPubKey()
	log.WriteLog("Public Key from Vpn Service " + publicKey.Address().String())
	log.WriteLog("denom" + msg.Coins.Denom)
	sentKey := senttype.GetNewSessionId()
	log.WriteLog("vpnaddr" + msg.Vpnaddr.String())
	// pubj, err := json.Marshal(publicKey)
	// if err != nil {
	// 	panic(err)
	// }
	var uPub crypto.PubKey
	// err = json.Unmarshal(pubj, &uPub)
	// if err != nil {
	// 	log.WriteLog("unmarswhal pub key failed")
	// }

	a, err := keeper.cdc.MarshalBinary(publicKey)
	if err != nil {
		log.WriteLog("unmarswhal pub key  keeper.cdc.unmarshal failed")
	}
	log.WriteLog("mrashal bytee keeper  " + string(a))
	err = keeper.cdc.UnmarshalBinary(a, &uPub)
	if err != nil {
		log.WriteLog("unmarshal failed from keeper.cdc.unmarshal")
	}
	log.WriteLog("Unmarshal pubkey upub " + uPub.Address().String())
	//	log.WriteLog("pubkey in json from json.Marshal :" + string(pubj[:]))

	// pubkey, err := keeper.cdc.MarshalBinary(publicKey)
	// if err != nil {
	// 	panic(err)
	// }
	// log.WriteLog("pubkey in string :" + string(pubkey[:]))
	sessionMap := senttype.GetNewSessionMap(string(a[:]), msg.Coins, msg.Vpnaddr)

	store := ctx.KVStore(keeper.sentStoreKey)

	log.WriteLog("Seesion Id" + string(sentKey))

	bz, err := json.Marshal(sessionMap)
	if err != nil {

	}

	log.WriteLog("bz map bytes SessionMap Obj" + string(bz))
	store.Set(sentKey, bz)
	keeper.coinKeeper.SubtractCoins(ctx, msg.From, sdk.Coins{msg.Coins})
	return string(sentKey[:]), nil
}

func (keeper Keeper) GetVpnPayment(ctx sdk.Context, msg MsgGetVpnPayment) ([]byte, sdk.Error) {

	var clientSession senttype.SessionMap
	var ClientPubkey crypto.PubKey
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.Sessionid)
	log.WriteLog("unmarshal string ......" + string(x))
	err := json.Unmarshal(x, &clientSession)
	if err != nil {
		panic(err)
	}

	keys := make([]string, 0, len(clientSession))
	for k := range clientSession {
		keys = append(keys, k)
	}
	//ke := reflect.ValueOf(clientSession).MapKeys()[0]
	// bz, err := hex.DecodeString(keys[0])
	// if err != nil {
	// 	return nil, err
	// }
	a2, err := json.Marshal(clientSession[keys[0]])
	if err != nil {
		log.WriteLog("unmarshal pubkey keeper.cdc.failesd")
	}
	log.WriteLog("pubkey associated struct" + string(a2))

	err = keeper.cdc.UnmarshalBinary([]byte(keys[0]), &ClientPubkey)
	if err != nil {
		log.WriteLog("unmarshal pubkey keeper.cdc.failesd")
	}
	log.WriteLog("pubkey address" + ClientPubkey.Address().String())
	log.WriteLog("clients session of publickey map bytes   : " + keys[0])
	a, err := json.Marshal(keys[0])
	if err != nil {
		panic(err)
	}
	log.WriteLog("pubksy " + string(a[:]))

	pubkey, err := crypto.PubKeyFromBytes([]byte(keys[0]))
	if err != nil {
		log.WriteLog("unmarshal Failed from crypto.pubkey")
	}
	a1, err := json.Marshal(pubkey)
	if err != nil {
		panic(err)
	}
	log.WriteLog("strinf pubkey " + string(a1))

	err = json.Unmarshal([]byte(keys[0]), &ClientPubkey) //
	if err != nil {
		log.WriteLog("unmarshal Failed")
	}
	log.WriteLog("client adress " + ClientPubkey.Address().String())

	log.WriteLog("msg sessionfi" + string(msg.Sessionid[:]) + "  counter " + string(msg.Counter) + "   the final vale ")
	signBytes := senttype.ClientStdSignBytes(msg.Coins, msg.Sessionid, msg.Counter, msg.IsFinal) //errr

	if !ClientPubkey.VerifyBytes(signBytes, msg.Signature) {
		return nil, sdk.ErrUnauthorized("signature from the keeper.go verification failed")
	}
	clientSessionData := clientSession[keys[0]]
	if clientSessionData.CurrentLockedCoins.IsPositive() && clientSessionData.Counter <= msg.Counter {
		CoinsToAdd := msg.Coins
		clientSessionData.CurrentLockedCoins = clientSessionData.CurrentLockedCoins.Minus(CoinsToAdd)
		clientSessionData.Counter = msg.Counter
		keeper.coinKeeper.AddCoins(ctx, clientSessionData.VpnAddr, sdk.Coins{CoinsToAdd})

		sentKey := msg.Sessionid
		bz, _ := keeper.cdc.MarshalBinary(clientSessionData)
		store.Set(sentKey, bz)
		return sentKey, nil
	}

	return nil, nil
}

func (keeper Keeper) RefundBal(ctx sdk.Context, msg MsgRefund) (sdk.AccAddress, sdk.Error) {

	var clientSession senttype.SessionMap
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.Sessionid)
	err := json.Unmarshal(x, &clientSession)
	if err != nil {
		panic(err)
	}
	pubkey := keeper.account.GetAccount(ctx, msg.From).GetPubKey()
	cpubkey, err := keeper.cdc.MarshalBinary(pubkey)
	if err != nil {

	}
	ctime := time.Now().UnixNano()
	if int64(math.Abs(float64(ctime))) >= 86400000 {

		keeper.coinKeeper.AddCoins(ctx, msg.From, sdk.Coins{clientSession[string(cpubkey[:])].CurrentLockedCoins})
		store.Delete(msg.Sessionid)
		return msg.From, nil
	} else {
		sdk.ErrCommon("time is less than 24 hours ").Result()
	}

	return nil, nil

}
