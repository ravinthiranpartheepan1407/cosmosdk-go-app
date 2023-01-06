package ravinetwork_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ravi-network/testutil/keeper"
	"ravi-network/testutil/nullify"
	"ravi-network/x/ravinetwork"
	"ravi-network/x/ravinetwork/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RavinetworkKeeper(t)
	ravinetwork.InitGenesis(ctx, *k, genesisState)
	got := ravinetwork.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
