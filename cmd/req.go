package cmd

import (

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reqCmd)
	reqCmd.PersistentFlags().StringVarP(&documentName, "document", "d", "", "Document name")

}

var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "send a request defined by document",
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
