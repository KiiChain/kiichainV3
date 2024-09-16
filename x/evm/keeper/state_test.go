package keeper_test

import (
	"testing"

	testkeeper "github.com/KiiChain/kiichainV3/testutil/keeper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestState(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	_, addr := testkeeper.MockAddressPair()

	initialState := k.GetState(ctx, addr, common.HexToHash("0xabc"))
	require.Equal(t, common.Hash{}, initialState)
	k.SetState(ctx, addr, common.HexToHash("0xabc"), common.HexToHash("0xdef"))

	got := k.GetState(ctx, addr, common.HexToHash("0xabc"))
	require.Equal(t, common.HexToHash("0xdef"), got)
}
