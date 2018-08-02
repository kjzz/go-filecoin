package iface

import (
	"context"

	"gx/ipfs/QmYVNvtQkeZ6AKSwDrjQTs432QtL6umrrK41EBq3cu7iSP/go-cid"

	"github.com/filecoin-project/go-filecoin/types"
)

type BlockAPI interface {
	Get(ctx context.Context, id *cid.Cid) (*types.Block, error)
}