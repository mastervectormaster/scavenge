package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mastervectormaster/scavenge/x/scavenge/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRevealSolution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal-solution [solution]",
		Short: "Broadcast message reveal-solution",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSolution := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevealSolution(
				clientCtx.GetFromAddress().String(),
				argSolution,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
