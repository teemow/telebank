package export

import (
	"fmt"
	"strconv"

	"github.com/ryanuber/columnize"
	aqbanking "github.com/umsatz/go-aqbanking"
)

const (
	transactionsHeader        = "Date | Name | Purpose | Type | Total"
	transactionsScheme        = "%s | %s | %s | %s | %s %s"
	transactionsMonthlyHeader = "Month | Total"
	transactionsMonthlyScheme = "%s | %s %s"
	defaultEllipsizeLength    = 25
)

func stringifyTransactions(transactions []aqbanking.Transaction) [][]string {
	records := [][]string{}

	for _, t := range transactions {
		records = append(records, []string{
			t.Purpose,
			t.Text,
			t.Status,
			t.CustomerReference,
			t.LocalBankCode,
			t.LocalAccountNumber,
			t.LocalIBAN,
			t.LocalBIC,
			t.LocalName,
			t.RemoteBankCode,
			t.RemoteAccountNumber,
			t.RemoteIBAN,
			t.RemoteBIC,
			t.RemoteName,
			t.Date.Format("02-01-2006"),
			t.ValutaDate.Format("02-01-2006"),
			strconv.FormatFloat(float64(t.Total), 'f', 2, 32),
			t.TotalCurrency,
			strconv.FormatFloat(float64(t.Fee), 'f', 2, 32),
			t.FeeCurrency,
		})
	}

	return records
}

func ellipsize(str string, length int) string {
	if len(str) > length {
		str = fmt.Sprintf("%s…", str[:length-1])
	}
	return str
}

func Out(transactions []aqbanking.Transaction, full bool) {
	lines := []string{transactionsHeader}
	var totals float32
	totals = 0
	for _, t := range transactions {
		date := t.Date.Format("02-01-2006")
		remote := t.RemoteName
		purpose := t.Purpose
		total := strconv.FormatFloat(float64(t.Total), 'f', 2, 32)

		if !full {
			remote = ellipsize(remote, defaultEllipsizeLength)
			purpose = ellipsize(purpose, defaultEllipsizeLength)
		}

		lines = append(lines, fmt.Sprintf(transactionsScheme, date, remote, purpose, t.Text, total, t.TotalCurrency))
		totals += t.Total
	}
	lines = append(lines, fmt.Sprintf(transactionsScheme, "Total", "", "", "", strconv.FormatFloat(float64(totals), 'f', 2, 32), "EUR"))
	fmt.Println(columnize.SimpleFormat(lines))
}

func Monthly(transactions []aqbanking.Transaction) {
	lines := []string{transactionsMonthlyHeader}
	var (
		totals       float32
		monthlyTotal float32
		date         string
	)
	totals = 0
	monthlyTotal = 0
	for _, t := range transactions {
		if date != "" && date != t.Date.Format("01-2006") {
			lines = append(lines, fmt.Sprintf(transactionsMonthlyScheme, t.Date.Format("01-2006"), strconv.FormatFloat(float64(monthlyTotal), 'f', 2, 32), "EUR"))
			monthlyTotal = 0
		}
		date = t.Date.Format("01-2006")
		monthlyTotal += t.Total
		totals += t.Total
	}
	lines = append(lines, fmt.Sprintf(transactionsMonthlyScheme, "Total", strconv.FormatFloat(float64(totals), 'f', 2, 32), "EUR"))
	fmt.Println(columnize.SimpleFormat(lines))
}
