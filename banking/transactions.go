package banking

import (
	"fmt"
	"log"
	"strings"
	"time"

	aqbanking "github.com/umsatz/go-aqbanking"
)

func TransactionsFor(ab *aqbanking.AQBanking, account *aqbanking.Account) ([]aqbanking.Transaction, error) {
	//fromDate := time.Date(2015, 01, 14, 0, 0, 0, 0, time.UTC)
	//toDate := time.Date(2015, 01, 16, 0, 0, 0, 0, time.UTC)
	//transactions, err := ab.Transactions(account, &fromDate, &toDate)

	transactions, err := ab.AllTransactions(account)
	if err != nil {
		return []aqbanking.Transaction{}, err
	}

	return transactions, nil
}

func Transactions(ab *aqbanking.AQBanking, userId string) []aqbanking.Transaction {
	userCollection, err := ab.Users()
	if err != nil {
		log.Fatal("unable to list users: %v", err)
	}

	allTransactions := []aqbanking.Transaction{}

	for _, user := range userCollection.Users {
		if userId != "" && userId != user.UserID {
			continue
		}

		accountList, err := ab.AccountsFor(&user)
		if err != nil {
			log.Fatal("unable to list accounts: %v", err)
		}

		for _, account := range accountList.Accounts {
			transactions, err := TransactionsFor(ab, &account)
			if err != nil {
				fmt.Printf("Unable to get transactions: %v\n", err)
			} else {
				allTransactions = append(allTransactions, transactions...)
			}
		}
	}

	return allTransactions
}

type Filter struct {
	RemoteName string
	Purpose    string
	ValueType  string
	Month      string
}

func FilterTransactions(transactions []aqbanking.Transaction, filter Filter) []aqbanking.Transaction {
	filteredTransactions := []aqbanking.Transaction{}
	for _, t := range transactions {
		valid := true

		if filter.ValueType != "both" {
			if filter.ValueType == "income" && t.Total <= 0 {
				valid = false
			} else if filter.ValueType == "expenses" && t.Total > 0 {
				valid = false
			}
		}

		if filter.RemoteName != "" && !strings.Contains(t.RemoteName, filter.RemoteName) {
			valid = false
		}

		if filter.Purpose != "" && !strings.Contains(t.Purpose, filter.Purpose) {
			valid = false
		}

		if filter.Month != "" {
			month, err := time.Parse("02-01-2006", filter.Month)
			if err != nil {
				log.Fatalf("Invalid month filter: %v\n", err)
			}

			if !(month.Month() == t.Date.Month() && month.Year() == t.Date.Year()) {
				valid = false
			}
		}

		if valid == true {
			filteredTransactions = append(filteredTransactions, t)
		}
	}
	return filteredTransactions
}
