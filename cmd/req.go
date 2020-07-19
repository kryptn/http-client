package cmd

import (
	"fmt"

	"github.com/kryptn/http-client/internal/client"
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

		sr := client.SendRequest(true)

		c := client.NewClient(client.WithDocument(args[0]), sr, client.SetDebugDumpAst(debugDumpAst), client.SetDebugDumpClient(debugDumpClient))

		result, err := c.Execute()
		if err != nil {
			return err
		}
		fmt.Print(result)
		return nil
	},
}
