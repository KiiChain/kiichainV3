package keeper

import (
	"math/big"

	"github.com/KiiChain/kiichainV3/x/evm/state"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k *Keeper) GetBalance(ctx sdk.Context, addr sdk.AccAddress) *big.Int {
	denom := k.GetBaseDenom(ctx)
	allUsei := k.BankKeeper().GetBalance(ctx, addr, denom).Amount
	lockedUsei := k.BankKeeper().LockedCoins(ctx, addr).AmountOf(denom) // LockedCoins doesn't use iterators
	usei := allUsei.Sub(lockedUsei)
	wei := k.BankKeeper().GetWeiBalance(ctx, addr)
	return usei.Mul(state.SdkUkiiToSweiMultiplier).Add(wei).BigInt()
}
