package main

import (
	"fmt"
	"log"

	"github.com/ryanuber/columnize"
	"github.com/spf13/cobra"
	"github.com/teemow/telebank/banking"
)

var (
	usersCmd = &cobra.Command{
		Use:   "users",
		Short: "Show users",
		Long:  "Show users",
		Run:   usersRun,
	}
)

const (
	usersHeader = "Name | UserID | CustomerID | BankCode | ServerURI | HbciVersion"
	usersScheme = "%s | %s | %s | %s | %s | %d"
)

func usersRun(cmd *cobra.Command, args []string) {
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

	users, err := banking.Users(aq)
	if err != nil {
		log.Fatal("Unable to list bank users: %v", err)
	}

	lines := []string{usersHeader}
	for _, u := range users.Users {
		lines = append(lines, fmt.Sprintf(usersScheme, u.Name, u.UserID, u.CustomerID, u.BankCode, u.ServerURI, u.HbciVersion))
	}
	fmt.Println(columnize.SimpleFormat(lines))
}
