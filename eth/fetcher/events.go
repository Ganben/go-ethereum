package fetcher

import "github.com/ganben/go-ethereum/core/types"

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
