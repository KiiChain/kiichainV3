package keeper_test

import (
	"testing"

	testkeeper "github.com/KiiChain/kiichainV3/testutil/keeper"
	"github.com/KiiChain/kiichainV3/x/evm/types"
	"github.com/stretchr/testify/require"
)

func TestParams(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	require.Equal(t, "ukii", k.GetBaseDenom(ctx))
	require.Equal(t, types.DefaultPriorityNormalizer, k.GetPriorityNormalizer(ctx))
	require.Equal(t, types.DefaultBaseFeePerGas, k.GetBaseFeePerGas(ctx))
	require.Equal(t, types.DefaultMinFeePerGas, k.GetMinimumFeePerGas(ctx))

	require.Nil(t, k.GetParams(ctx).Validate())
}
