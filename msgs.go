package sentinel

import (
	"crypto"
	"encoding/json"

	types "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	crypto "github.com/tendermint/go-crypto"
)
//

//
//
//
//
type MsgRegisterVpnService struct {
	address  sdk.Address
	ip       string
	netspeed string
	ppgb     string
	//signature auth.Signature
}

func (msc MsgRegisterVpnService) Type() string {
	return "sentinel"
}

func (msc MsgRegisterVpnService) GetSignBytes() []byte {
    crypto.
	return b
}

func (msc MsgRegisterVpnService) ValidateBasic() sdk.Error { 
	if msc.address == nil {
		return nil
	}
	if msc.ppgb == "" {
		return nil
	}
	if msc.ip == "" {
		return nil
	}
	if msc.netspeed == "" {
		return nil
	}
	return nil
}

func (msc MsgRegisterVpnService) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
func (msc MsgRegisterVpnService) NewMsgRegisterVpnService(address sdk.Address, ip, ppgb, netspeed string) *MsgRegisterVpnService {
	return &MsgRegisterVpnService{
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
	pubkey  crypto.PubKey
}

//Newmsgcreate : TODO

func (msc MsgRegisterMasterNode) Type() string {
	return ""
}

func (msc MsgRegisterMasterNode) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgRegisterMasterNode) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return nil
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

func (msc MsgQueryRegisteredVpnService) Type() string {
	return "sentinel"
}

func (msc MsgQueryRegisteredVpnService) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgQueryRegisteredVpnService) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return nil
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
	coins   sdk.Coin
	pubkey  *crypto.PubKey
	vpnaddr sdk.Address
}

func (msc MsgPayVpnService) Type() string {
	return ""
}

func (msc MsgPayVpnService) GetSignBytes() []byte {
	b, _ := json.Marshal(msc)
	return b
}

func (msc MsgPayVpnService) ValidateBasic() sdk.Error {
	if msc.coins == nil {
		return nil
	}
	if msc.pubkey == nil {
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
	b, _ := json.Marshal(msc)
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
