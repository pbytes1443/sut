package sentinel

import (
	"crypto"
	"encoding/json"
	"reflect"

	"strconv"
	"strings"

	types "github.com/cosmos/cosmos-sdk/examples/sentinel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	me "github.com/sut/types"
	crypto "github.com/tendermint/go-crypto"
)

//

//
//
//
//

/// USE gofmt command for styling/structing the go code

type MsgRegisterVpnService struct {
	address  sdk.AccAddress
	ip       string
	netspeed int64
	ppgb     int64
	//signature auth.Signature                    // TODO :  sign all the above params using client signature and verify it by using public key from account mapper
}

func isIp(host string) bool {
	parts := strings.Split(host, ".")

	if len(parts) < 4 {
		return false
	}

	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}

	}
	return true
}

func (msc MsgRegisterVpnService) Type() string {
	return "sentinel"
}

func (msc MsgRegisterVpnService) GetSignBytes() []byte {
	var byte_format []byte
	byte_format, err := json.Marshal(msc)
	if err != nil{
		return err
	}
	return byte_format
}
func (msc MsgRegisterVpnService) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return sdk.ErrInvalidAddress("Address type is Invalid").Result()
	}
	if msc.ppgb != nil || reflect.TypeOf(msc.ppgb) != int64 || msg.ppgb > 0 || msg.ppgb < 1000 {

		return me.ErrCommon("Price per GB is not Valid").Result()
	}
	if msc.ip != "" || !isIp(msc.ip) || reflect.TypeOf(msc.ip) != string {

		return me.ErrInvalidIp("Ip is not Valid").Result()
	}
	if msc.netspeed != nil || reflect.TypeOf(msc.netspeed) != int64 || msg.netspeed > 0 {
		return me.ErrCommon("NetSpeed is not Valid").Result()
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
}
func (msc MsgRegisterMasterNode) Type() string {
	return "sentinel"
}

func (msc MsgRegisterMasterNode) GetSignBytes() []byte {
	byte_format, err := json.Marshal(msc) 
	if err != nil {
		return err
	}
	return byte_format
}

func (msc MsgRegisterMasterNode) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return sdk.ErrInvalidAddress("Address type is Invalid").Result()
	}
	return nil
}
func (msc MsgRegisterMasterNode) GetSigners() []sdk.Address {
	return []sdk.Address{msc.address}
}
func (msc MsgRegisterVpnService) NewMsgRegisterMasterNode(address sdk.Address) MsgRegisterMasterNode {
	return MsgRegisterMasterNode{
		address:  address
	}
}
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
	byte_format,err := json.Marshal(msc)
	if err!= nil{
		return err
	}
	return byte_format
}

func (msc MsgQueryRegisteredVpnService) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return sdk.ErrInvalidAddress("Address type is Invalid").Result()
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
	return "sentinel"
}

func (msc MsgQueryFromMasterNode) GetSignBytes() []byte {
	byte_format, err:= json.Marshal(msc)
	if err != nil{
		return err
	}
	return byte_format
}

func (msc MsgQueryFromMasterNode) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return sdk.ErrInvalidAddress("Address type is Invalid").Result()
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
	return "sentinel"
}

func (msc MsgDeleteVpnUser) GetSignBytes() []byte {
	byte_format,err := json.Marshal(msc)
	if err != nil{
		return err
	}
	return byte_format
}

func (msc MsgDeleteVpnUser) ValidateBasic() sdk.Error {
	if msc.address == nil {
		return sdk.ErrInvalidAddress("Address type is Invalid").Result()
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
	return "sentinel"
}

func (msc MsgDeleteMasterNode) GetSignBytes() []byte {
	byte_format,err := json.Marshal(msc)
	if err != nil{
		return err
	}	
	return byte_format
}

func (msc MsgDeleteMasterNode) ValidateBasic() sdk.Error {

	//TODO:CHECK THE SIZE OF MSG.ADDRESS at each and every ValidateBasic() METHOD.

	if msc.address == nil {
		return sdk.ErrInvalidAddress("Address type is Invalid").Result()
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
	return "sentinel"
}

func (msc MsgPayVpnService) GetSignBytes() []byte {
	byte_format,err := json.Marshal(msc)
	if err != nil{
		return err
	}
	return byte_format
}

func (msc MsgPayVpnService) ValidateBasic() sdk.Error {
	if msc.coins == nil {
		return sdk.ErrInsifficientCoins("Error insufficient coins").Result()
	}	
	if msc.pubkey == nil {
		return sdk.ErrInvalidPubKey("Invalid public key").Result()
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
	address   sdk.Adress
	sessionid int64
	signature auth.Signature
	from      sdk.Address
}

func (msc MsgSigntoVpn) Type() string {
	return "sentinel"
}

func (msc MsgSigntoVpn) GetSignBytes() []byte {
	byte_format,err := json.Marshal(msc)
	id err != nil{
		return err
	}
	return byte_format
}

func (msc MsgSigntoVpn) ValidateBasic() sdk.Error {
	if msc.coins == nil {
		return sdk.ErrInsifficientCoins("Error insufficent coins").Result()
	}
	if msc.sessionid != nil ||reflect.TypeOf(msc.sessionid) != int64 {
		return sdk.Error
	}
	if msc.address == nil {
		return sdk.ErrInvalidAddress("Invalid Address").Result()
	}
	if msc.from == nil {
		return sdk.ErrInvalidAddress("Invalid  from Address").Result()
	}

	if msc.signature ==nil{
		return sdk.Err      //TODO validate signature
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
	from      sdk.Address
}

func (msc MsgGetVpnPayment) Type() string {
	return "sentinel"
}

func (msc MsgGetVpnPayment) GetSignBytes() []byte {
	byte_format, err := json.MarshalJSON(msc)
	if err != nil{
		return err
	}
	//b, _ := json.Marshal(msc)
	return byte_format
}

func (msc MsgGetVpnPayment) ValidateBasic() sdk.Error {
	if msc.coins == nil {
		return sdk.ErrInsifficientCoins("Error insufficient coins").Result()
	}
	if msc.pubkey == nil {
		return sdk.ErrInvalidPubKey("Invalid public key").Result()
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
type MsgRefund struct {
	pubkey crypto.PublicKey
	sessionid int64
}

func (msc MsgRefund) Type() string {
	return "sentinel"
}

func (msc MsgRefund) GetSignBytes() []byte {
	byte_format, err := json.Marshal(msc)
	if err != nil{
		return err
	}
	return byte_format
}

func (msc MsgRefund) ValidateBasic() sdk.Error {
	if msc.sessionid != nil ||reflect.TypeOf(msc.sessionid) != int64 {
		return sdk.Error
	}
	if msc.pubkey == nil {
		return sdk.ErrInvalidPubKey("Error Invalid public key.......").Result()
	}
	return nil
}
func (msc MsgRefund) GetSigners() []sdk.Address {
	return []sdk.Address{msc.from}
}
