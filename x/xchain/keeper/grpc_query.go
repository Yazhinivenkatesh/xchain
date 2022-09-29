package keeper

import (
	"github.com/yazhini/xchain/x/xchain/types"
)

var _ types.QueryServer = Keeper{}
