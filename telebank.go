package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	globalFlags struct {
		debug   bool
		verbose bool
	}

	telebankCmd = &cobra.Command{
		Use:   "telebank",
		Short: "telebank is a simple aqbanking-cli replacement written in go",
		Long:  `telebank manages your bank accounts and lists your transactions.`,
		Run:   telebankRun,
	}

	projectVersion string
)

func init() {
	initGlobalFlags(telebankCmd.PersistentFlags())
}

func initGlobalFlags(flags *pflag.FlagSet) {
	flags.BoolVarP(&globalFlags.debug, "debug", "d", false, "Print debug output")
	flags.BoolVarP(&globalFlags.verbose, "verbose", "v", false, "Print verbose output")
}

func telebankRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	telebankCmd.AddCommand(versionCmd)
	telebankCmd.AddCommand(transactionsCmd)
	telebankCmd.AddCommand(accountsCmd)
	telebankCmd.AddCommand(usersCmd)

	telebankCmd.Execute()
}
