package cmd

import (
	"github.com/mungkiice/nutriGo-Golang/database/seed"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed fake data into database",
	//Long: `You can migrate, refresh, rollback, or even seed your database using commands`,
	Run: func(cmd *cobra.Command, args []string) {
		seed.Run()
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
