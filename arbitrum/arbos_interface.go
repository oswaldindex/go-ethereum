package arbitrum

import (
	"context"

	"github.com/oswaldindex/arbi-geth/arbitrum_types"
	"github.com/oswaldindex/arbi-geth/core"
	"github.com/oswaldindex/arbi-geth/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
