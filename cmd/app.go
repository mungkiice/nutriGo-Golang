package cmd

import (
	"log"

	"github.com/mungkiice/nutriGo-Golang/route"
	"github.com/spf13/cobra"
)

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Run app",
	Long:  `You can migrate, refresh, rollback, or even seed your database using commands`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := route.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(appCmd)
}
