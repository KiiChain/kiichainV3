package keeper_test

import (
	"bytes"
	"testing"

	keepertest "github.com/KiiChain/kiichainV3/testutil/keeper"
	"github.com/KiiChain/kiichainV3/x/evm/keeper"
	"github.com/stretchr/testify/require"
)

func TestInitGenesis(t *testing.T) {
	k, ctx := keepertest.MockEVMKeeper() // this would call `InitGenesis`
	// coinbase address must be associated
	coinbaseSeiAddr, associated := k.GetSeiAddress(ctx, keeper.GetCoinbaseAddress())
	require.True(t, associated)
	require.True(t, bytes.Equal(coinbaseSeiAddr, k.AccountKeeper().GetModuleAddress("fee_collector")))
}
