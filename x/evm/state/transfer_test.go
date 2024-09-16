package state_test

import (
	"math/big"
	"testing"

	testkeeper "github.com/KiiChain/kiichainV3/testutil/keeper"
	"github.com/KiiChain/kiichainV3/x/evm/state"
	"github.com/stretchr/testify/require"
)

func TestEventlessTransfer(t *testing.T) {
	k, ctx := testkeeper.MockEVMKeeper()
	db := state.NewDBImpl(ctx, k, false)
	_, fromAddr := testkeeper.MockAddressPair()
	_, toAddr := testkeeper.MockAddressPair()

	beforeLen := len(ctx.EventManager().ABCIEvents())

	state.TransferWithoutEvents(db, fromAddr, toAddr, big.NewInt(100))

	// should be unchanged
	require.Len(t, ctx.EventManager().ABCIEvents(), beforeLen)
}
