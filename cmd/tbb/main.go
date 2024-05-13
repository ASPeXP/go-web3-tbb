package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main(){
	var tbbCmd = &cobra.Command{
		Use: "tbb",
		Short: "The Blockchain Bar CLI",
		Run: func(cmd *cobra.Command, args []string ){

		},
	}

	tbbCmd.AddCommand(versionCmd)
	
	err := tbbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err )
		os.Exit(1)
	}

	
	
}