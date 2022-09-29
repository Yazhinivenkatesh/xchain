package xchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/yazhini/xchain/testutil/keeper"
	"github.com/yazhini/xchain/testutil/nullify"
	"github.com/yazhini/xchain/x/xchain"
	"github.com/yazhini/xchain/x/xchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.XchainKeeper(t)
	xchain.InitGenesis(ctx, *k, genesisState)
	got := xchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
