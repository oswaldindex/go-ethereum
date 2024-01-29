package eth

import (
	"context"

	"github.com/oswaldindex/arbi-geth/core"
	"github.com/oswaldindex/arbi-geth/core/state"
	"github.com/oswaldindex/arbi-geth/core/types"
	"github.com/oswaldindex/arbi-geth/core/vm"
	"github.com/oswaldindex/arbi-geth/eth/tracers"
	"github.com/oswaldindex/arbi-geth/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
