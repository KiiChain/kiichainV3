package aclmapping

import (
	aclbankmapping "github.com/KiiChain/kiichainV3/aclmapping/bank"
	aclevmmapping "github.com/KiiChain/kiichainV3/aclmapping/evm"
	acloraclemapping "github.com/KiiChain/kiichainV3/aclmapping/oracle"
	acltokenfactorymapping "github.com/KiiChain/kiichainV3/aclmapping/tokenfactory"
	aclwasmmapping "github.com/KiiChain/kiichainV3/aclmapping/wasm"
	evmkeeper "github.com/KiiChain/kiichainV3/x/evm/keeper"
	aclkeeper "github.com/cosmos/cosmos-sdk/x/accesscontrol/keeper"
)

type CustomDependencyGenerator struct{}

func NewCustomDependencyGenerator() CustomDependencyGenerator {
	return CustomDependencyGenerator{}
}

func (customDepGen CustomDependencyGenerator) GetCustomDependencyGenerators(evmKeeper evmkeeper.Keeper) aclkeeper.DependencyGeneratorMap {
	dependencyGeneratorMap := make(aclkeeper.DependencyGeneratorMap)
	wasmDependencyGenerators := aclwasmmapping.NewWasmDependencyGenerator()

	dependencyGeneratorMap = dependencyGeneratorMap.Merge(aclbankmapping.GetBankDepedencyGenerator())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acltokenfactorymapping.GetTokenFactoryDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(wasmDependencyGenerators.GetWasmDependencyGenerators())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(acloraclemapping.GetOracleDependencyGenerator())
	dependencyGeneratorMap = dependencyGeneratorMap.Merge(aclevmmapping.GetEVMDependencyGenerators(evmKeeper))

	return dependencyGeneratorMap
}
