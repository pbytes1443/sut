package types

type Registervpn struct {
	Ip       string
	Netspeed int64
	Ppgb     int64
	// Coins    sdk.Coin
	Location string
	// PLEASE CHECK THIS PROPERLY, remove this
	// TODO :/// MUST ADD LOCATION PARAMETER of type string
}

func NewVpnRegister(ip, location string, ppgb, netspeed int64) Registervpn {
	return Registervpn{
		Ip:       ip,
		Netspeed: netspeed,
		Ppgb:     netspeed,
		Location: location,
	}
}
