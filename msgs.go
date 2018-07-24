package sentinel

import (
	"crypto"
	"encoding/json"

	types "github.com/cosmos/cosmos-sdk/examples/sut/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/tendermint/crypto"
)
//

//
//
//
//


 /// USE gofmt command for styling/structing the go code



type MsgRegisterVpnService struct {
	address  sdk.Address
	ip       string
	netspeed string                                          //TODO: CHANGE THE NETSPEED TYPE TO INT
	ppgb     string                                              // TODO:   CHANGE THE PRICE PER GB to INT
	//signature auth.Signature                    // TODO :  sign all the above params using client signature and verify it by using public key from account mapper
}

func (msc MsgRegisterVpnService) Type() string {
	return "sentinel"
}

func (msc MsgRegisterVpnService) GetSignBytes() []byte {
	return b                                                       //TODO : JUST RETURN WHAT COSMOS HAS RETURNED FOR THIS SPECIFIC
	                                                                       
}

func (msc MsgRegisterVpnService) ValidateBasic() sdk.Error { 
	if msc.address == nil {
		return nil                                    //TODO: SHOULD RETURN SDK.ERROR MSG
	}
	if msc.ppgb == "" {                  //VALIDATE           int for lower bound=0 and upper bound

		return nil                                   
	}
	if msc.ip == "" {                       /// CHeck If the type is string   and also validate for a proper ip
		return nil
	}
	if msc.netspeed == "" {           ////check for lower bound.... you will get data in bytes/sec
		return nil
	}
	return nil
}

func (msc MsgRegisterVpnService) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
func (msc MsgRegisterVpnService) NewMsgRegisterVpnService(address sdk.Address, ip, ppgb, netspeed string) MsgRegisterVpnService {
	return MsgRegisterVpnService{
		address:  address,
		ip:       ip,
		ppgb:     ppgb,
		netspeed: netspeed,
	}
}
//
//
//
//
//
type MsgRegisterMasterNode struct {
	address sdk.Address
	pubkey  crypto.PubKey                         //TODO REMOVE PUBLIC KEY PARAMETER , AND GET PUBLIC KEY PARAMETER FROM THE BANK KEEPER AS
	                                                                      // bank.keeper.account_mapper.GETAccount()  ---> ACCOUNT WILL HAVE PUBLIC KEY
}

//Newmsgcreate : TODO

func (msc MsgRegisterMasterNode) Type() string {
	return ""
}

func (msc MsgRegisterMasterNode) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)                                      //TODO : PLEASE SET THIS ACCORDINGLY
	return b
}

func (msc MsgRegisterMasterNode) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return nil                                   // ERROR MESSAGE
	}
	if msc.pubkey == nil {
		return nil
	}
	return nil
}
func (msc MsgRegisterMasterNode) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}//
//
//
//
//
//
type MsgQueryRegisteredVpnService struct {
	address sdk.Address
}
/// SHould  restrict QUERYABLE -----> MYTAKS_ALLAGOG
func (msc MsgQueryRegisteredVpnService) Type() string {
	return "sentinel"
}

func (msc MsgQueryRegisteredVpnService) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgQueryRegisteredVpnService) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return nil             //error msg and perform similar address validation as above
	}
	return nil
}

func (msc MsgQueryRegisteredVpnService) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
//
//
//
//
//
type MsgQueryFromMasterNode struct {
	address sdk.Address
}

func (msc MsgQueryFromMasterNode) Type() string {
	return ""
}

func (msc MsgQueryFromMasterNode) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgQueryFromMasterNode) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return nil
	}

	return nil
}
func (msc MsgQueryFromMasterNode) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
//
//
//
//
//
type MsgDeleteVpnUser struct {
	addressService sdk.Address
}

func (msc MsgDeleteVpnUser) Type() string {
	return ""
}

func (msc MsgDeleteVpnUser) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgDeleteVpnUser) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return nil
	}

	return nil
}
func (msc MsgDeleteVpnUser) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}

//
//
//
//
//
type MsgDeleteMasterNode struct {
	address sdk.Address
}

func (msc MsgDeleteMasterNode) Type() string {
	return ""
}

func (msc MsgDeleteMasterNode) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgDeleteMasterNode) ValidateBasic() sdk.Error {

//TODO:CHECK THE SIZE OF MSG.ADDRESS at each and every ValidateBasic() METHOD. 

	if msc.address == nil {
		return nil
	}

	return nil
}
func (msc MsgDeleteMasterNode) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
//
//
//
//
//
type MsgPayVpnService struct {
	coins   sdk.Coin                             //TODO ADD ADDRESS TO THIS TRANSACTION
	pubkey  *crypto.PubKey           // TODO :Should take it from account bankkeeper.am.GetAccount().Pubkey
	vpnaddr sdk.Address
}

func (msc MsgPayVpnService) Type() string {
	return ""
}

func (msc MsgPayVpnService) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)            //todo check for everysuch function
	return b
}

func (msc MsgPayVpnService) ValidateBasic() sdk.Error {
	if msc.coins == nil {                          /// TODO : parse coin and check for valid denom and minimum amount
		return nil                                        
	}
	if msc.pubkey == nil {                /// remove this thing
		return nil
	}
	return nil
}
func (msc MsgPayVpnService) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
//
//
//
//
//
type MsgSigntoVpn struct {
	coins sdk.Coin
	//	pubkey   tmcrypto.PubKey
	address   sdk.Adress
	sessionid int64
	// counter   int64
	// timestamp
	signature auth.Signature
	from      sdk.Address
}

func (msc MsgSigntoVpn) Type() string {
	return ""
}

func (msc MsgSigntoVpn) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgSigntoVpn) ValidateBasic() sdk.Error {
	if msc.coins == nil {
		return nil
	}
	if msc.sessionid==""{
		return error
	}
	if msc.pubkey == nil {
		return nil
	}
	return nil
}
func (msc MsgSigntoVpn) GetSigners() []sdk.Address {
	return []sdk.Address{msg.from}
}
//
//
//
//

type MsgGetVpnPayment struct {
	clientSig types.ClientSignature
	from sdk.Address
}

func (msc MsgGetVpnPayment) Type() string {
	return ""
}

func (msc MsgGetVpnPayment) GetSignBytes() []byte {
	b, err := msgCdc.MarshalJSON(msc)
	//b, _ := json.Marshal(msc)
	return b
}

func (msc MsgGetVpnPayment) ValidateBasic() sdk.Error {
	if msc.coins == nil {
		return nil
	}
	if msc.pubkey == nil {
		return nil
	}
	return nil
}
func (msc MsgGetVpnPayment) GetSigners() []sdk.Address {
	return []sdk.Address{msc.from}
}
//
//
//
//
//
type MsgRefund struct{
	pubkey crypto.PublicKey
	//coins sdk.Coin
	sessionid int64
}


func (msc MsgRefund) Type() string {
	return ""
}

func (msc MsgRefund) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgRefund) ValidateBasic() sdk.Error {
	if msc.sessionid == "" {
		return error
	}
	if msc.pubkey == nil {
		return nil
	}
	return nil
}
func (msc MsgRefund) GetSigners() []sdk.Address {
	return []sdk.Address{msc.from}
}
