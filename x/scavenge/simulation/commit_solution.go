package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/mastervectormaster/scavenge/x/scavenge/keeper"
	"github.com/mastervectormaster/scavenge/x/scavenge/types"
)

func SimulateMsgCommitSolution(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCommitSolution{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CommitSolution simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CommitSolution simulation not implemented"), nil, nil
	}
}
