package ante_test

import (
	"testing"

	testkeeper "github.com/KiiChain/kiichainV3/testutil/keeper"
	"github.com/KiiChain/kiichainV3/x/evm/ante"
	"github.com/KiiChain/kiichainV3/x/evm/types"
	"github.com/KiiChain/kiichainV3/x/evm/types/ethtx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGasLimitDecorator(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	a := ante.NewGasLimitDecorator(k)
	limitMsg, _ := types.NewMsgEVMTransaction(&ethtx.LegacyTx{GasLimit: 100})
	ctx, err := a.AnteHandle(ctx, &mockTx{msgs: []sdk.Msg{limitMsg}}, false, func(ctx sdk.Context, _ sdk.Tx, _ bool) (sdk.Context, error) {
		return ctx, nil
	})
	require.Nil(t, err)
	require.Equal(t, 100, int(ctx.GasMeter().Limit()))
}
