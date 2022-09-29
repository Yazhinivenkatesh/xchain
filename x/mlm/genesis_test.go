package mlm_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/yazhini/xchain/testutil/keeper"
	"github.com/yazhini/xchain/testutil/nullify"
	"github.com/yazhini/xchain/x/mlm"
	"github.com/yazhini/xchain/x/mlm/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MlmKeeper(t)
	mlm.InitGenesis(ctx, *k, genesisState)
	got := mlm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
