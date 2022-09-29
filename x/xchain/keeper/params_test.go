package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/yazhini/xchain/testutil/keeper"
	"github.com/yazhini/xchain/x/xchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.XchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
