package services

import (
	"context"

	core "github.com/iden3/go-iden3-core/v2"

	"github.com/polygonid/sh-id-platform/internal/config"
	"github.com/polygonid/sh-id-platform/internal/log"
)

const (
	DIDMethodOptimismID   = "opid"
	DIDMethodOptimismByte = 0b01000011
)

// RegisterCustomDIDMethods registers custom DID methods
func RegisterCustomDIDMethods(ctx context.Context, customsDis []config.CustomDIDMethods) error {
	for _, cdid := range customsDis {
		params := core.DIDMethodNetworkParams{
			Method:      DIDMethodOptimismID,
			Blockchain:  core.Blockchain(cdid.Blockchain),
			Network:     core.NetworkID(cdid.Network),
			NetworkFlag: cdid.NetworkFlag,
		}
		if err := core.RegisterDIDMethodNetwork(params, core.WithChainID(cdid.ChainID)); err != nil {
			log.Error(ctx, "cannot register custom DID method", "err", err, "customDID", cdid)
			return err
		}
	}
	return nil
}

func RegisterOptimismIdMethod(ctx context.Context) error {
	return core.RegisterDIDMethod(DIDMethodOptimismID, DIDMethodOptimismByte)
}
