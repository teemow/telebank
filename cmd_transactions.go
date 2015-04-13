package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/teemow/telebank/banking"
	"github.com/teemow/telebank/export"
)

var (
	transactionsCmd = &cobra.Command{
		Use:   "transactions",
		Short: "Show transactions",
		Long:  "Show transactions",
		Run:   transactionsRun,
	}

	transactionsFlags struct {
		name      string
		purpose   string
		user      string
		export    string
		valueType string
		month     string
		monthly   bool
		full      bool
	}
)

const (
	defaultExport    = "stdout"
	defaultValueType = "both"
	defaultUser      = ""
	defaultName      = ""
	defaultPurpose   = ""
	defaultMonth     = ""
	defaultMonthly   = false
	defaultFull      = false
)

func init() {
	initTransactionsFlags(transactionsCmd.Flags())
}

func initTransactionsFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&transactionsFlags.user, "user", "u", defaultUser, "List transations of a specific user id")
	flags.StringVarP(&transactionsFlags.export, "export", "e", defaultExport, "Export of the transations eg. csv (Default: stdout)")
	flags.StringVarP(&transactionsFlags.valueType, "type", "t", defaultValueType, "Value type of the transations eg. income, expenses (Default: both)")
	flags.StringVarP(&transactionsFlags.name, "name", "r", defaultName, "Filter transactions by name of the counterpart")
	flags.StringVarP(&transactionsFlags.purpose, "purpose", "p", defaultPurpose, "Filter transactions by purpose")
	flags.StringVarP(&transactionsFlags.month, "month", "m", defaultMonth, "Filter transactions by month")
	flags.BoolVar(&transactionsFlags.monthly, "monthly", defaultMonthly, "Summarize transactions by month")
	flags.BoolVarP(&transactionsFlags.full, "full", "l", defaultFull, "Do not ellipsize fields on output")
}

func transactionsRun(cmd *cobra.Command, args []string) {
	aq, err := banking.Init(globalFlags.verbose, globalFlags.debug)
	if err != nil {
		log.Fatal("unable to init banking: %v", err)
	}
	defer aq.Free()

	if globalFlags.verbose {
		fmt.Printf("using aqbanking %d.%d.%d\n",
			aq.Version.Major,
			aq.Version.Minor,
			aq.Version.Patchlevel,
		)
	}

	for _, pin := range banking.LoadPins("pins.json") {
		aq.RegisterPin(pin)
	}

	transactions := banking.Transactions(aq, transactionsFlags.user)

	transactions = banking.FilterTransactions(transactions, banking.Filter{
		RemoteName: transactionsFlags.name,
		Purpose:    transactionsFlags.purpose,
		ValueType:  transactionsFlags.valueType,
		Month:      transactionsFlags.month,
	})

	if transactionsFlags.monthly {
		export.Monthly(transactions)
	} else {
		switch transactionsFlags.export {
		case "csv":
			export.WriteCSV("transactions.csv", transactions)
		case "stdout":
			export.Out(transactions, transactionsFlags.full)
		}
	}
}
