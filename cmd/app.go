package cmd

import (
	"github.com/mungkiice/goNutri/route"
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run app",
	Long: `You can migrate, refresh, rollback, or even seed your database using commands`,
	Run: func(cmd *cobra.Command, args []string) {
		route.NewRouter().Run(":8000")
	},
}

func init(){
	rootCmd.AddCommand(appCmd)
}