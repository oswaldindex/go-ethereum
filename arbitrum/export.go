package arbitrum

import (
	"context"

	"github.com/oswaldindex/arbi-geth/common/hexutil"
	"github.com/oswaldindex/arbi-geth/core"
	"github.com/oswaldindex/arbi-geth/internal/ethapi"
	"github.com/oswaldindex/arbi-geth/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *ethapi.StateOverride, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}
