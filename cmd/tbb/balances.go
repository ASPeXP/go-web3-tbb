package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/web3coach/the-blockchain-bar/database"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command {
		Use: "balances",
		Short: "Interact with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string ) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var balancesListCmd = &cobra.Command {
		Use: "list",
		Short: "Lists all balances",
		Run: func(cmd *cobra.Command, args []string ){
			state, err := database.NewStateFromDisk()
			if err != nil {
				fmt.Fprintln(os.Stderr, err )
				os.Exit(1)
			}
			defer state.Close()

			fmt.Println("Accounts balances:")
			fmt.Println("------------------")
			fmt.Println("")
			for account, balance := range state.Balances {
				fmt.Println(fmt.Sprintf("%s: %d", account, balance))
			}
		},
	}
	balancesCmd.AddCommand(balancesListCmd)
	return balancesCmd
}
