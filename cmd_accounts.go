package main

import (
	"fmt"
	"log"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"github.com/teemow/telebank/banking"
)

var (
	accountsCmd = &cobra.Command{
		Use:   "accounts",
		Short: "Show accounts",
		Long:  "Show accounts",
		Run:   accountsRun,
	}
)

const (
	accountsHeader = "Name | Number | Owner | Bank | BIC"
	accountsScheme = "%s | %s | %s | %s | %s"
)

func accountsRun(cmd *cobra.Command, args []string) {
	aq, err := banking.Init(globalFlags.verbose, globalFlags.debug)
	if err != nil {
		log.Fatal("Unable to init banking: %v", err)
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

	accounts, err := banking.Accounts(aq)
	if err != nil {
		log.Fatal("Unable to list bank accounts: %v", err)
	}

	lines := []string{accountsHeader}
	for _, a := range accounts.Accounts {
		lines = append(lines, fmt.Sprintf(accountsScheme, a.Name, a.AccountNumber, a.Owner, a.Bank.Name, a.Bank.BankCode))
	}
	fmt.Println(columnize.SimpleFormat(lines))
}
