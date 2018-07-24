package client

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	sentinel "github.com/cosmos/cosmos-sdk/examples/sut"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

const (
	flagvpnaddr         = "vpn_addr"
	flaguseraddr        = "user_addr"
	flagamount          = "amount"
	flagreceivedbytes   = "receivedBytes"
	flagsessionduration = "sessionduration"
	flagtimestamp       = "timestamp"
	flagsessionid       = "sessionid"
	flagip              = "vpn_ip"
	flagnetspeed        = "netspeed"
	flagppgb            = "price_per_gb"
	flagqueryaddress    = "address"
	flagmspubkey        = "pubkey"
) /*

func AddVpnUsageTxCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add_vpn_usage",
		Short: "Add vpn usage",
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			vpn := viper.GetString(flagvpnaddr)
			vpn_addr := strings.TrimSpace(vpn)
			if len(vpn_addr) == 0 {
				return fmt.Errorf("1")
			}

			vpn_addr_hex, err := hex.DecodeString(vpn_addr)
			if err != nil {
				return err
			}

			vpnaddr := sdk.Address(vpn_addr_hex)

			receivedbytes := viper.GetInt64(flagreceivedbytes)
			if receivedbytes == 0 {
				return fmt.Errorf("3")

			}

			session_duration := viper.GetInt64(flagsessionduration)
			if session_duration == 0 {
				return fmt.Errorf("2")

			}

			amount := viper.GetInt64(flagamount)
			if amount == 0 {
				return fmt.Errorf("2")

			}

			timestamp := viper.GetInt64(flagtimestamp)
			if timestamp == 0 {
				return fmt.Errorf("2")

			}

			session_id := viper.GetString(flagsessionid)
			session_id = strings.TrimSpace(session_id)
			fmt.Printf("%v", session_id)

			if len(session_id) == 0 {
				return fmt.Errorf("2")

			}

			msg := sentinel.MsgAddVpnUsage{vpnaddr, sender, amount, receivedbytes, session_duration, timestamp, session_id}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			newSessionId := new(string)
			if err != nil {
				return err
			}
			err = cdc.UnmarshalBinary(res.DeliverTx.Data, &newSessionId)
			if err != nil {
				return err
			}
			fmt.Printf("sentinel created with id: %s\n", *newSessionId)
			return nil
		},
	}

	cmd.Flags().String(flagvpnaddr, "", "vpn provider Addresse")
	cmd.Flags().String(flaguseraddr, "", "Sender Address")
	cmd.Flags().String(flagamount, "", "Amount to put into sentinel")
	cmd.Flags().String(flagreceivedbytes, "", "ReceivedBytes")
	cmd.Flags().String(flagsessionduration, "", "Session Duration")
	cmd.Flags().String(flagtimestamp, "", "timestamp")
	cmd.Flags().String(flagsessionid, "", "session id")

	return cmd
}

func PayVpnUsageTxCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay_vpn_usage",
		Short: "pay for sentinel service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			amount := viper.GetInt64(flagamount)
			if amount == 0 {
				return fmt.Errorf("2")

			}
			session_id := viper.GetString(flagsessionid)
			session_id = strings.TrimSpace(session_id)
			if len(session_id) == 0 {
				return fmt.Errorf("2")
			}
			msg := sentinel.MsgPayVpnUsage{sender, amount, session_id}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			newSessionId := new(string)
			if err != nil {
				return err
			}
			err = cdc.UnmarshalBinary(res.DeliverTx.Data, *newSessionId)
			if err != nil {
				return err
			}
			fmt.Printf("sentinel created with id: %s\n", *newSessionId)
			return nil
		},
	}
	cmd.Flags().String(flagamount, "", "Amount to put into sentinel")
	cmd.Flags().String(flagsessionid, "", "session id")
	return cmd
}
*/
func RegisterVpnServiceCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "regvpn",
		Short: "Register for sentinel vpn service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			ip := viper.GetString(flagip)
			if ip == "" {
				return fmt.Errorf("Ip flag is not m entioned")

			}
			ppgb := viper.GetString(flagppgb)
			if len(ppgb) == 0 {
				return fmt.Errorf("price per gb not mentioned")
			}
			netspeed := viper.GetString(flagnetspeed)
			if len(netspeed) == 0 {
				return fmt.Errorf("net speed not mentioned")
			}
			address, err := sdk.GetAccAddressBech32(ctx.FromAddressName)
			if err != nil {
				return err
			}

			msg := sentinel.MsgRegisterVpnService.Assign(address, ip, ppgb, netspeed)
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			fmt.Printf("Vpn serivicer registered with address: %s\n", sender)
			return nil
		},
	}
	cmd.Flags().String(flagip, "", "Ip address")
	cmd.Flags().String(flagppgb, "", "price per gb")
	cmd.Flags().String(flagnetspeed, "", "net speed")
	return cmd
}

//
//
//

func RegisterMasterNodeCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "regmsnode",
		Short: "Register Master node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			mspubkey := viper.GetString(flagmspubkey)
			if mspubkey == "" {
				return fmt.Errorf("public key is not mentioned")

			}
			msg := sentinel.MsgRegisterMasterNode{sender, mspubkey}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			fmt.Printf("Masternode registered with address: %s\n", sender)
			return nil
		},
	}
	cmd.Flags().String(flagmspubkey, "", "register master node")
	return cmd
}

func QueryVpnServiceCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vpnservice [address]",
		Short: "query for sentinel vpn service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper()
			// sender, err := ctx.GetFromAddress()
			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				 return err
			}
			// queryaddress := viper.GetString(flagqueryaddress)
			// if queryaddress == "" {
				// return fmt.Errorf("address is not mentioned")
			// }
			// address, err := sdk.GetAccAddressBech32(queryaddress)
			// if err != nil {
				// return err
			// }
			res,err := ctx.QueryStore(key, storeName)
			if err != nil {
				return err
			} else if len(res) == 0 {
				return fmt.Errorf("No vpn service provider found with address %s", args[0])
			}
			vpn_service := types.MustUnmarshalValidator(cdc, addr, res)
			switch viper.Get(cli.OutputFlag) {
			case "text":
				human, err := validator.HumanReadableString()
				if err != nil {
					return err
				}
				fmt.Println(human)

			case "json":
				// parse out the validator
				output, err := wire.MarshalJSONIndent(cdc, validator)
				if err != nil {
					return err
				}
				fmt.Println(string(output))
			}
			// TODO output with proofs / machine parseable etc.
			return cmd
		},
	}

	return cmd
}

			// msg := sentinel.MsgQueryRegisteredVpnService{address}
			// res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			//fmt.Printf("Vpn serivicer registered with address: %s\n", sender)
			return nil
		},
	}
	cmd.Flags().String(flagqueryaddress, "", "to query vpn service")
	return cmd
}
func QueryMasterNodeCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "query to  master node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			queryaddress := viper.GetString(flagqueryaddress)
			if queryaddress == "" {
				return fmt.Errorf("address is not mentioned")
			}
			key, err := sdk.GetAccAddressBech32(queryaddress)
			if err != nil {
				return err
			}
			msg := sentinel.MsgQueryFromMasterNode{key}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			//fmt.Printf("Vpn serivicer registered with address: %s\n", sender)
			return nil
		},
	}
	cmd.Flags().String(flagqueryaddress, "", "to query masternode")
	return cmd
}

func UnRegisterMasterNodeCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unreg_msnode",
		Short: "Unregister Master node",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			addr := viper.GetString(flagqueryaddress)
			if addr == "" {
				return fmt.Errorf("address is not mentioned")

			}
			key, err := sdk.GetAccAddressBech32(addr)
			if err != nil {
				return err
			}
			msg := sentinel.MsgDeleteMasterNode{key}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			//fmt.Printf("Masternode registered with address: %s\n", sender)
			return nil
		},
	}
	cmd.Flags().String(flagmspubkey, "", "register master node")
	return cmd
}
func UnRegisterVpnServiceCmd(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unreg_vpn",
		Short: "Unregister vpn service",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.NewCoreContextFromViper().WithDecoder(authcmd.GetAccountDecoder(cdc))
			sender, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}
			addr := viper.GetString(flagqueryaddress)
			if addr == "" {
				return fmt.Errorf("address is not mentioned")

			}
			key, err := sdk.GetAccAddressBech32(addr)
			if err != nil {
				return err
			}

			msg := sentinel.MsgDeleteVpnUser{key}
			res, err := ctx.EnsureSignBuildBroadcast(ctx.FromAddressName, msg, cdc)
			//fmt.Printf("Masternode registered with address: %s\n", sender)
			return nil
		},
	}
	cmd.Flags().String(flagqueryaddress, "", "Unregister vpn node")
	return cmd
}
