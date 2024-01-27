package main // Use the appropriate package name

import (
	"context"
	"math"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"testing"
	"time"
	"github.com/ethereum/go-ethereum/ethclient"
)
func TestConnectToTestnet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Assuming rpcURL is defined as a constant or variable that contains your Ethereum testnet RPC URL
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		t.Fatalf("Failed to connect to the testnet: %v", err)
	}
	defer client.Close()

	// Fetching the network ID
	networkID, err := client.NetworkID(ctx)
	if err != nil {
		t.Fatalf("Failed to get network ID: %v", err)
	}

	// Fetching the latest block number
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		t.Fatalf("Failed to get the latest block number: %v", err)
	}

	t.Logf("Network ID: %v", networkID)
	t.Logf("Latest block number: %d", blockNumber)
}
