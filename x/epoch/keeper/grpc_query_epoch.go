package keeper

import (
	"context"

	"github.com/KiiChain/kiichainV3/x/epoch/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Epoch(c context.Context, _ *types.QueryEpochRequest) (*types.QueryEpochResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	epoch := k.GetEpoch(ctx)
	return &types.QueryEpochResponse{Epoch: epoch}, nil
}
