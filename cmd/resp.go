package cmd

import (
	"fmt"

	"github.com/kryptn/http-client/internal/client"
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

		sr := client.SendRequest(false)

		c := client.NewClient(client.WithDocument(args[0]), sr, client.SetDebugDumpAst(debugDumpAst), client.SetDebugDumpClient(debugDumpClient))

		result, err := c.Execute()
		if err != nil {
			return err
		}
		fmt.Print(result)
		return nil
	},
}
