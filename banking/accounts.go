package banking

import aqbanking "github.com/umsatz/go-aqbanking"

func Accounts(ab *aqbanking.AQBanking) (*aqbanking.AccountCollection, error) {
	accountList, err := ab.Accounts()
	if err != nil {
		return &aqbanking.AccountCollection{}, err
	}

	return accountList, nil
}
