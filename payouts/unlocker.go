// File: unlocker.go

package payouts

import (
	"fmt"
	"log"
)

// Define network constants
const (
	ZetherMainnet = "ZetherMainnet"
	ZetherTestnet = "ZetherTestnet"
)

// NetworkConfig holds network-specific configurations
type NetworkConfig struct {
	Name          string
	BlockReward   map[int64]float64 // Block height -> Reward
	AdjustmentFreq int64            // Frequency of reward adjustment in blocks
	HasUncles     bool              // Indicates if the network considers uncle blocks
}

// Reward schedule for Zether mainnet
var ZetherRewards = map[int64]float64{
	0:       50.0,  // Initial reward
	100000:  25.0,  // Reduced reward
	200000:  12.5,  // Further reduced
	300000:   6.25, // Final reduction
}

// Reward schedule for ZTH-Test network
var TestnetRewards = map[int64]float64{
	0:    50.0,
	1000: 25.0,
	2000: 12.5,
	3000:  6.25,
}

// Define network configurations
var ZetherNetwork = NetworkConfig{
	Name:          ZetherMainnet,
	BlockReward:   ZetherRewards,
	AdjustmentFreq: 100000,
	HasUncles:     false,
}

var ZetherTestnet = NetworkConfig{
	Name:          ZetherTestnet,
	BlockReward:   TestnetRewards,
	AdjustmentFreq: 1000,
	HasUncles:     false,
}

// GetReward calculates the block reward based on network configuration and block height
func GetReward(config NetworkConfig, blockHeight int64) float64 {
	reward := 0.0
	for height, r := range config.BlockReward {
		if blockHeight >= height {
			reward = r
		} else {
			break
		}
	}
	return reward
}

// SimulateUnlock simulates the unlock process, calculating rewards for a given block height
func SimulateUnlock(networkConfig NetworkConfig, blockHeight int64) {
	if blockHeight < 0 {
		log.Fatalf("Invalid block height: %d", blockHeight)
	}

	reward := GetReward(networkConfig, blockHeight)
	fmt.Printf("Network: %s\n", networkConfig.Name)
	fmt.Printf("Block Height: %d\n", blockHeight)
	fmt.Printf("Block Reward: %.2f ZTH\n", reward)

	if !networkConfig.HasUncles {
		fmt.Println("Note: This network does not include uncle blocks.")
	}
}

func main() {
	// Test mainnet reward schedule
	fmt.Println("---- Zether Mainnet ----")
	SimulateUnlock(ZetherNetwork, 50000)  // Check reward before adjustment
	SimulateUnlock(ZetherNetwork, 150000) // Check reward after first adjustment

	// Test testnet reward schedule
	fmt.Println("---- Zether Testnet ----")
	SimulateUnlock(ZetherTestnet, 500)    // Check reward before adjustment
	SimulateUnlock(ZetherTestnet, 1500)   // Check reward after first adjustment
}
