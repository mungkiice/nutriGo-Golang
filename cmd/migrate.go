package cmd

import (
	"github.com/mungkiice/goNutri/database/migration"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Create or update database",
	Long: `You can migrate, refresh, rollback, or even seed your database using commands`,
	Run: func(cmd *cobra.Command, args []string) {
		migration.Up()
	},
}

func init(){
	rootCmd.AddCommand(migrateCmd)
}