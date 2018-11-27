package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "goNutri",
	Short: "~~~ Welcome to the goNutri CLI Interface ~~~",
	//Long: `You can migrate, refresh, rollback, or even seed your database using commands`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
