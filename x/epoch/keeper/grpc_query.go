package keeper

import (
	"github.com/KiiChain/kiichainV3/x/epoch/types"
)

var _ types.QueryServer = Keeper{}
