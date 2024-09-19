package migrations_test

import (
	"testing"
	"time"

	testkeeper "github.com/KiiChain/kiichainV3/testutil/keeper"
	"github.com/KiiChain/kiichainV3/x/evm/migrations"
	"github.com/KiiChain/kiichainV3/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestMigrateCastAddressBalances(t *testing.T) {
	k := testkeeper.EVMTestApp.EvmKeeper
	ctx := testkeeper.EVMTestApp.GetContextForDeliverTx([]byte{}).WithBlockTime(time.Now())
	require.Nil(t, k.BankKeeper().MintCoins(ctx, types.ModuleName, testkeeper.UseiCoins(100)))
	// unassociated account with funds
	seiAddr1, evmAddr1 := testkeeper.MockAddressPair()
	require.Nil(t, k.BankKeeper().SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(evmAddr1[:]), testkeeper.UseiCoins(10)))
	// associated account without funds
	seiAddr2, evmAddr2 := testkeeper.MockAddressPair()
	k.SetAddressMapping(ctx, seiAddr2, evmAddr2)
	// associated account with funds
	seiAddr3, evmAddr3 := testkeeper.MockAddressPair()
	require.Nil(t, k.BankKeeper().SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(evmAddr3[:]), testkeeper.UseiCoins(10)))
	k.SetAddressMapping(ctx, seiAddr3, evmAddr3)

	require.Nil(t, migrations.MigrateCastAddressBalances(ctx, &k))

	require.Equal(t, sdk.NewInt(10), k.BankKeeper().GetBalance(ctx, sdk.AccAddress(evmAddr1[:]), "ukii").Amount)
	require.Equal(t, sdk.ZeroInt(), k.BankKeeper().GetBalance(ctx, seiAddr1, "ukii").Amount)
	require.Equal(t, sdk.ZeroInt(), k.BankKeeper().GetBalance(ctx, sdk.AccAddress(evmAddr2[:]), "ukii").Amount)
	require.Equal(t, sdk.ZeroInt(), k.BankKeeper().GetBalance(ctx, seiAddr2, "ukii").Amount)
	require.Equal(t, sdk.ZeroInt(), k.BankKeeper().GetBalance(ctx, sdk.AccAddress(evmAddr3[:]), "ukii").Amount)
	require.Equal(t, sdk.NewInt(10), k.BankKeeper().GetBalance(ctx, seiAddr3, "ukii").Amount)
}
