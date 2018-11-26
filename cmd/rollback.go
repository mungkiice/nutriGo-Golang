package cmd

import (
	"github.com/mungkiice/goNutri/database/migration"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Drop all tables on database",
	//Long: `You can migrate, refresh, rollback, or even seed your database using commands`,
	Run: func(cmd *cobra.Command, args []string) {
		migration.Down()
	},
}

func init(){
	rootCmd.AddCommand(rollbackCmd)
}