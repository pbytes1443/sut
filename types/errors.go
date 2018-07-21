package types

// abci "github.com/suttypes"
err "github.com/cosmos/cosmos-sdk/types"

const (
	CodeUnknownRequest CodeType = 6
)

func ErrInvalidIp(msg string) Error {
	return err.newErrorWithRootCodespace(CodeUnknownRequest, msg)
}
func ErrCommon(msg string) Error {
	return err.newErrorWithRootCodespace(CodeUnknownRequest, msg)
}
