package sentinel

import (

	//	"fmt"

	senttype "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	bank "github.com/cosmos/cosmos-sdk/x/bank"
	crypto "github.com/tendermint/tendermint/crypto"
	//	"strconv"
	//	"strings"
)

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
	vpnaddr sdk.AccAddress
	counter int64
	hash    string
	sign    crypto.PubKeySecp256k1
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, ck bank.Keeper, codespace sdk.CodespaceType) Keeper {
	keeper := Keeper{
		sentStoreKey: key,
		cdc:          cdc,
		coinKeeper:   ck,
		codespace:    codespace, ////learn WHAT THIS DOES
	}
	return keeper
}

func (keeper Keeper) RegisterVpnService(ctx sdk.Context, msg MsgRegisterVpnService) (string, sdk.Error) {
	sentKey := msg.From.String()
	store := ctx.KVStore(keeper.sentStoreKey)
	var vpnreg senttype.Registervpn
	bz, _ := keeper.cdc.MarshalBinary(vpnreg)
	store.Set([]byte(sentKey), bz)
	return "", nil
}

func (keeper Keeper) QueryRegisteredVpnService(ctx sdk.Context, msg MsgQueryRegisteredVpnService) (senttype.Registervpn, sdk.Error) {
	store := ctx.KVStore(keeper.sentStoreKey)
	var vpnreg senttype.Registervpn
	bz := store.Get(msg.address)
	if bz != nil {
		sdk.ErrCommon("Address is not valid").Result()
	}
	keeper.cdc.UnmarshalBinary(bz, &vpnreg)
	return vpnreg, nil
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

func (keeper Keeper) PayVpnService(ctx sdk.Context, msg MsgPayVpnService, acc auth.AccountMapper) (string, sdk.Error) {
	//var sessionMap senttype.SessionMap
	publicKey := acc.GetAccount(ctx, msg.From).GetPubKey()
	// fmt.Println(publicKey.String())
	sessionMap := senttype.GetNewSessionMap(publicKey, msg.Coins, msg.Timestamp, msg.Vpnaddr)
	store := ctx.KVStore(keeper.sentStoreKey)
	sentKey := senttype.GetNewSessionId()
	bz, _ := keeper.cdc.MarshalBinary(sessionMap) //PLEASE USE SMAE VARIABLE NAMES  FOR FUNCTION parameter
	store.Set(sentKey, bz)
	keeper.coinKeeper.SubtractCoins(ctx, msg.From, sdk.Coins{msg.Coins}) //coins type sdk.Coins
	return string(sentKey), nil
}

// func (keeper Keeper) sendSigntoVpn(ctx sdk.Context, msg MsgSigntoVpn) sdk.Error {
// 	// var k keys.Keybasei
// 	// k, err = Keybase.GetKeyBase()
// 	// sig, pub,err := k.Sign(name,passphrase,msg)
// 	// msg.sign:sig
// 	vpnsignstore[pub][msg]
// 	return nil
// }/*
/*
func (keeper Keeper) GetVpnPayment(ctx sdk.Context, msg MsgGetVpnPayment) ([]byte, sdk.Error) { // TODO Change the function name to GetVpnPayment

	var clientSession senttype.SessionMap
	signature := msg.ClientSig.Signature
	pubKey := signature.Pubkey
	sig := signature.Signature
	store := ctx.KVStore(keeper.sentStoreKey)
	key := msg.ClientSig.Sessionid
	x := store.Get(key) //changes to be done
	err := keeper.cdc.UnmarshalBinary(x, &clientSession)

	//ClientPubkey := reflect.ValueOf(clientSession).MapKeys()[0]
	//ClientPubkey.

		ctx.
	if ClientPubkey.String() != pubKey.String() {
		sdk.ErrCommon("Invalid Public key").Result()
	}
	signBytes := senttype.ClientStdSignBytes(msg.ClientSig.Coins, msg.ClientSig.Sessionid, msg.ClientSig.Counter)
	//need to be evaluated
	pubkey := sdk.GetValPubKeyBech32(ClientPubkey.String())
	if !(ClientPubkey.VerifyBytes(msg.GetSignBytes(), signature.Signature)) { //type cast
		return nil, sdk.ErrUnauthorized("signature verification failed").Result()
	}
	clientSessionData := clientSession[ClientPubkey]
	if clientsig := msg.clientsig; clientSessionData.currentLockedCoins > 0 && clientSessionData.counter <= clientSig.counter {
		CoinsToAdd := msg.clientSig.coins.Minus(UnlockedCoins)
		clientSessionData.currentLockedCoins = clientSessionData.currentLockedCoins.Minus(CoinsToAdd)
		clientSessionData.UnlockedCoins = clientSessionData.UnlockedCoins.Plus(clientSessionData.totalLockedCoins.Minus(clientSessionData.currentLockedCoins))
		clientSessionData.counter = msg.clientSig.counter
		keeper.bankKeeper.AddCoins(ctx, clientSessionData.vpnaddr, coins)

	}
	sentkey := clientSig.sessionid
	bz, _ := keeper.cdc.MarshalBinary(clientSessionData)
	store.Set(sentKey, bz)
	return sentKey, nil
}
func (keeper Keeper) RefundBal(ctx sdk.Context, msg MsgRefund) (sdk.AccAddress, sdk.Error) {

	var t time.Time
	var clientSession senttype.SessionMap
	store := ctx.KVStore(keeper.sentStoreKey)
	x := store.Get(msg.sessionid)
	err := keeper.cdc.UnmarshalBinary(x, &clientSession)
	// if err != nil {
	// 	return nil, err
	// }
	a := reflect.ValueOf(keeper.coinKeeper)
	c := a.FieldByName("am").GetAccount(ctx, msg.from).PubKey()
	pubkey := keeper.coinKeeper.am.GetPubkey(ctx, msg.from)
	clientSessionData := clientSession[pubkey]
	tm := clientSessionData.timestamp
	diff := time.Now().Sub(tm)
	//
	if diff.Hours() >= 24 {
		coins = clientSessionData.currentLockedCoins
		keeper.bankKeeper.AddCoins(ctx, msg.address, coins)
		store.Delete(msg.sessionid)
	}

	return msg.from, nil

}
*/
