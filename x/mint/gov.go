package mint

import (
	"github.com/KiiChain/kiichainV3/x/mint/keeper"
	"github.com/KiiChain/kiichainV3/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func HandleUpdateMinterProposal(ctx sdk.Context, k *keeper.Keeper, p *types.UpdateMinterProposal) error {
	err := types.ValidateMinter(*p.Minter)
	if err != nil {
		return err
	}
	k.SetMinter(ctx, *p.Minter)
	return nil
}
