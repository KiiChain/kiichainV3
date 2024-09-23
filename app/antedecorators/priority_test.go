package antedecorators_test

import (
	"testing"

	"github.com/KiiChain/kiichainV3/app/antedecorators"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestPriorityAnteDecorator(t *testing.T) {
	output = ""
	anteDecorators := []sdk.AnteFullDecorator{
		sdk.DefaultWrappedAnteDecorator(antedecorators.NewPriorityDecorator()),
	}
	ctx := sdk.NewContext(nil, tmproto.Header{}, false, nil)
	chainedHandler, _ := sdk.ChainAnteDecorators(anteDecorators...)
	// test with normal priority
	newCtx, err := chainedHandler(
		ctx.WithPriority(125),
		FakeTx{},
		false,
	)
	require.NoError(t, err)
	require.Equal(t, int64(125), newCtx.Priority())
}
