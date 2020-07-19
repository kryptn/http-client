package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile      string
	userLicense  string
	documentName string

	debugDumpAst    bool
	debugDumpClient bool

	rootCmd = &cobra.Command{
		Use:   "http-client",
		Short: "A request tool",
		Long:  `lots of words,`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&debugDumpAst, "dump-ast", false, "Dump AST")
	rootCmd.PersistentFlags().BoolVar(&debugDumpClient, "dump-client", false, "Dump Client")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	viper.AutomaticEnv()

}
