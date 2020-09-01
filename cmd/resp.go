package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(respCmd)

	respCmd.PersistentFlags().StringVarP(&documentName, "document", "d", "", "Document name")

}

var respCmd = &cobra.Command{
	Use:   "resp",
	Short: "echo last response of a document",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
