package banking

import aqbanking "github.com/umsatz/go-aqbanking"

func Users(ab *aqbanking.AQBanking) (*aqbanking.UserCollection, error) {
	userList, err := ab.Users()
	if err != nil {
		return &aqbanking.UserCollection{}, err
	}

	return userList, nil
}
