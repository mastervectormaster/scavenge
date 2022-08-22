package scavenge_test

import (
	"testing"

	keepertest "github.com/mastervectormaster/scavenge/testutil/keeper"
	"github.com/mastervectormaster/scavenge/testutil/nullify"
	"github.com/mastervectormaster/scavenge/x/scavenge"
	"github.com/mastervectormaster/scavenge/x/scavenge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ScavengeKeeper(t)
	scavenge.InitGenesis(ctx, *k, genesisState)
	got := scavenge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
