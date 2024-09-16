package keeper

import (
	"github.com/KiiChain/kiichainV3/x/epoch/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AfterEpochEnd(ctx sdk.Context, epoch types.Epoch) {
	k.hooks.AfterEpochEnd(ctx, epoch)
}

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epoch types.Epoch) {
	k.hooks.BeforeEpochStart(ctx, epoch)
}
