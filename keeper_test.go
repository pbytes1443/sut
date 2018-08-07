package sentinel

import (
	"testing"

	senttype "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	log "github.com/tendermint/tendermint/libs/log"
)

func CreateMultiStore() (sdk.MultiStore, *sdk.KVStoreKey, *sdk.KVStoreKey) {
	db := dbm.NewMemDB()
	authkey := sdk.NewKVStoreKey("authkey")
	sentinelkey := sdk.NewKVStoreKey("sentinel")
	ms := store.NewCommitMultiStore(db)
	ms.MountStoreWithDB(authkey, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(sentinelkey, sdk.StoreTypeIAVL, db)
	ms.LoadLatestVersion()
	return ms, authkey, sentinelkey

}

func TestPayVpnService(t *testing.T) {
	ms, authkey, sentkey := CreateMultiStore()

	cdc := wire.NewCodec()
	auth.RegisterBaseAccount(cdc)
	ac := auth.NewAccountMapper(cdc, authkey, auth.ProtoBaseAccount)
	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	ac.NewAccountWithAddress(ctx, addr1)
	ac.NewAccountWithAddress(ctx, addr2)
	keeper := Keeper{
		sentStoreKey: sentkey,
		coinKeeper:   bank.NewKeeper(ac),
		cdc:          cdc,
		codespace:    DefaultCodeSpace,
		account:      ac,
	}
	t.Log(keeper)
	t.Log(ctx)
	mp1 := NewMsgPayVpnService(coinPos, addr2, addr1)
	a := mp1.Type()
	require.Equal(t, a, "sentinel")
	require.Equal(t, mp1.GetSigners(), []sdk.AccAddress{addr1})
	t.Log(keeper.sentStoreKey)
	b, add := keeper.GetsentStore(ctx, MsgRegisterMasterNode{Address: addr1})
	require.Equal(t, add, addr1)
	t.Log(b)
	t.Log(keeper.PayVpnService(ctx, mp1))

	//require.NotNil(t, sessionid)
	// ////require.Nil(t,err)
	//t.Log(sessionid)
}

func TestGetVpnPayment(t *testing.T) {
	var err error

	ms, authkey, sentkey := CreateMultiStore()
	cdc := wire.NewCodec()
	auth.RegisterBaseAccount(cdc)
	ctx := sdk.NewContext(ms, abci.Header{}, false, log.NewNopLogger())
	accountmapper := auth.NewAccountMapper(cdc, authkey, auth.ProtoBaseAccount)
	keeper := NewKeeper(cdc, sentkey, bank.NewKeeper(accountmapper), accountmapper, DefaultCodeSpace)

	sessionid := []byte("iK7FDcCc35S4IzoOjgm2")

	require.Nil(t, err)

	clientsession := senttype.GetNewSessionMap(coinPos, pk2, pk1)
	require.Equal(t, clientsession.CPubKey, pk1)
	bz := senttype.ClientStdSignBytes(coinPos, sessionid, 1, false)
	t.Log(bz)
	sign1, err = pvk1.Sign(bz)
	t.Log(ctx, keeper)
	t.Log(sign1)
	//mg := MsgGetVpnPayment{
	//	Signature: sign1,
	//	Coins:     coinPos,
	//	Sessionid: sessionid,
	//	Counter:   1,
	//	Pubkey:    pk1,
	//	From:      addr2,
	//	IsFinal:   false,
	//}
	a, err := keeper.GetVpnPayment(ctx, MsgGetVpnPayment{Signature: sign1, Coins: coinPos, Sessionid: sessionid, Pubkey: pk1, From: addr2, IsFinal: false, Counter: 1})
	require.Nil(t, err)
	require.Equal(t, sessionid, a)
}
