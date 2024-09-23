package keeper

import (
	"context"
	"errors"

	"github.com/KiiChain/kiichainV3/x/evm/artifacts/cw20"
	"github.com/KiiChain/kiichainV3/x/evm/artifacts/cw721"
	"github.com/KiiChain/kiichainV3/x/evm/artifacts/erc20"
	"github.com/KiiChain/kiichainV3/x/evm/artifacts/erc721"
	"github.com/KiiChain/kiichainV3/x/evm/artifacts/native"
	"github.com/KiiChain/kiichainV3/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

var _ types.QueryServer = Querier{}

// Querier defines a wrapper around the x/mint keeper providing gRPC method
// handlers.
type Querier struct {
	*Keeper
}

func NewQuerier(k *Keeper) Querier {
	return Querier{Keeper: k}
}

func (q Querier) KiiAddressByEVMAddress(c context.Context, req *types.QueryKiiAddressByEVMAddressRequest) (*types.QueryKiiAddressByEVMAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req.EvmAddress == "" {
		return nil, sdkerrors.ErrInvalidRequest
	}
	evmAddr := common.HexToAddress(req.EvmAddress)
	addr, found := q.Keeper.GetKiiAddress(ctx, evmAddr)
	if !found {
		return &types.QueryKiiAddressByEVMAddressResponse{Associated: false}, nil
	}

	return &types.QueryKiiAddressByEVMAddressResponse{KiiAddress: addr.String(), Associated: true}, nil
}

func (q Querier) EVMAddressByKiiAddress(c context.Context, req *types.QueryEVMAddressByKiiAddressRequest) (*types.QueryEVMAddressByKiiAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req.KiiAddress == "" {
		return nil, sdkerrors.ErrInvalidRequest
	}
	seiAddr, err := sdk.AccAddressFromBech32(req.KiiAddress)
	if err != nil {
		return nil, err
	}
	addr, found := q.Keeper.GetEVMAddress(ctx, seiAddr)
	if !found {
		return &types.QueryEVMAddressByKiiAddressResponse{Associated: false}, nil
	}

	return &types.QueryEVMAddressByKiiAddressResponse{EvmAddress: addr.Hex(), Associated: true}, nil
}

func (q Querier) StaticCall(c context.Context, req *types.QueryStaticCallRequest) (*types.QueryStaticCallResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if req.To == "" {
		return nil, errors.New("cannot use static call to create contracts")
	}
	if ctx.GasMeter().Limit() == 0 {
		ctx = ctx.WithGasMeter(sdk.NewGasMeterWithMultiplier(ctx, q.QueryConfig.GasLimit))
	}
	to := common.HexToAddress(req.To)
	res, err := q.Keeper.StaticCallEVM(ctx, q.Keeper.AccountKeeper().GetModuleAddress(types.ModuleName), &to, req.Data)
	if err != nil {
		return nil, err
	}
	return &types.QueryStaticCallResponse{Data: res}, nil
}

func (q Querier) Pointer(c context.Context, req *types.QueryPointerRequest) (*types.QueryPointerResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	switch req.PointerType {
	case types.PointerType_NATIVE:
		p, v, e := q.Keeper.GetERC20NativePointer(ctx, req.Pointee)
		return &types.QueryPointerResponse{
			Pointer: p.Hex(),
			Version: uint32(v),
			Exists:  e,
		}, nil
	case types.PointerType_CW20:
		p, v, e := q.Keeper.GetERC20CW20Pointer(ctx, req.Pointee)
		return &types.QueryPointerResponse{
			Pointer: p.Hex(),
			Version: uint32(v),
			Exists:  e,
		}, nil
	case types.PointerType_CW721:
		p, v, e := q.Keeper.GetERC721CW721Pointer(ctx, req.Pointee)
		return &types.QueryPointerResponse{
			Pointer: p.Hex(),
			Version: uint32(v),
			Exists:  e,
		}, nil
	case types.PointerType_ERC20:
		p, v, e := q.Keeper.GetCW20ERC20Pointer(ctx, common.HexToAddress(req.Pointee))
		return &types.QueryPointerResponse{
			Pointer: p.String(),
			Version: uint32(v),
			Exists:  e,
		}, nil
	case types.PointerType_ERC721:
		p, v, e := q.Keeper.GetCW721ERC721Pointer(ctx, common.HexToAddress(req.Pointee))
		return &types.QueryPointerResponse{
			Pointer: p.String(),
			Version: uint32(v),
			Exists:  e,
		}, nil
	default:
		return nil, errors.ErrUnsupported
	}
}

func (q Querier) PointerVersion(c context.Context, req *types.QueryPointerVersionRequest) (*types.QueryPointerVersionResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	switch req.PointerType {
	case types.PointerType_NATIVE:
		return &types.QueryPointerVersionResponse{
			Version: uint32(native.CurrentVersion),
		}, nil
	case types.PointerType_CW20:
		return &types.QueryPointerVersionResponse{
			Version: uint32(cw20.CurrentVersion(ctx)),
		}, nil
	case types.PointerType_CW721:
		return &types.QueryPointerVersionResponse{
			Version: uint32(cw721.CurrentVersion),
		}, nil
	case types.PointerType_ERC20:
		return &types.QueryPointerVersionResponse{
			Version:  uint32(erc20.CurrentVersion),
			CwCodeId: q.GetStoredPointerCodeID(ctx, types.PointerType_ERC20),
		}, nil
	case types.PointerType_ERC721:
		return &types.QueryPointerVersionResponse{
			Version:  uint32(erc721.CurrentVersion),
			CwCodeId: q.GetStoredPointerCodeID(ctx, types.PointerType_ERC721),
		}, nil
	default:
		return nil, errors.ErrUnsupported
	}
}
