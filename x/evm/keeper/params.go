package keeper

import (
	"math/big"

	"github.com/KiiChain/kiichainV3/utils"
	"github.com/KiiChain/kiichainV3/x/evm/config"
	"github.com/KiiChain/kiichainV3/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const BaseDenom = "ukii"

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.Paramstore.SetParamSet(ctx, &params)
}

func (k *Keeper) GetParams(ctx sdk.Context) types.Params {
	params := types.Params{}
	k.Paramstore.GetParamSet(ctx, &params)
	return params
}

func (k *Keeper) GetBaseDenom(ctx sdk.Context) string {
	return BaseDenom
}

func (k *Keeper) GetPriorityNormalizer(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).PriorityNormalizer
}

func (k *Keeper) GetBaseFeePerGas(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).BaseFeePerGas
}

func (k *Keeper) GetMinimumFeePerGas(ctx sdk.Context) sdk.Dec {
	return k.GetParams(ctx).MinimumFeePerGas
}

func (k *Keeper) ChainID(ctx sdk.Context) *big.Int {
	if k.EthReplayConfig.Enabled || k.EthBlockTestConfig.Enabled {
		// replay is for eth mainnet so always return 1
		return utils.Big1
	}
	// return mapped chain ID
	return config.GetEVMChainID(ctx.ChainID())

}
