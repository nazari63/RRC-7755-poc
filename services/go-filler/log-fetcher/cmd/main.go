package main

import (
	"log"
	"math/big"

	"github.com/base-org/RIP-7755-poc/services/go-filler/log-fetcher/internal/chains"
	"github.com/base-org/RIP-7755-poc/services/go-filler/log-fetcher/internal/config"
	"github.com/base-org/RIP-7755-poc/services/go-filler/log-fetcher/internal/listener"
	"github.com/base-org/RIP-7755-poc/services/go-filler/log-fetcher/internal/store"
)

func main() {
	cfg, err := config.NewConfig() // Load env vars
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := store.NewMongoClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Close()

	srcChain := chains.GetChainConfig(big.NewInt(421614), cfg.RPCs)

	err = listener.Init(srcChain, cfg, mongoClient)
	if err != nil {
		log.Fatal(err)
	}
}
