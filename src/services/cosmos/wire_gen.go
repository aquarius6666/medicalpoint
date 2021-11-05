// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package cosmos

import (
	"context"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/common"
)

// Injectors from cosmos.wire.go:

func InitCosmosServiceClient(ctx context.Context, logger *logrus.Logger, ChainId ChainID, KeyRing keyring.UnsafeKeyring, CosmosEp CosmosEndpoint, Mne Mnemonic) (*CosmosServiceClient, error) {
	cosmosOpts := CosmosOpts{
		ChainId:  ChainId,
		KeyRing:  KeyRing,
		CosmosEp: CosmosEp,
		Mne:      Mne,
	}
	cosmosServiceClient := NewCosmosServiceClient(ctx, logger, cosmosOpts)
	return cosmosServiceClient, nil
}

func InitTestCosmosServiceClient(ctx context.Context, ChainId ChainID, KeyRing keyring.UnsafeKeyring, CosmosEp CosmosEndpoint, Mne Mnemonic) (*CosmosServiceClient, error) {
	logger := common.InitLoggerWithoutCLIContext()
	cosmosOpts := CosmosOpts{
		ChainId:  ChainId,
		KeyRing:  KeyRing,
		CosmosEp: CosmosEp,
		Mne:      Mne,
	}
	cosmosServiceClient := NewCosmosServiceClient(ctx, logger, cosmosOpts)
	return cosmosServiceClient, nil
}