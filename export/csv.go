package export

import (
	"encoding/csv"
	"fmt"
	"os"

	aqbanking "github.com/umsatz/go-aqbanking"
)

func WriteCSV(filename string, transactions []aqbanking.Transaction) {
	csvfile, err := os.Create(fmt.Sprintf("/log/%s", filename))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvfile.Close()

	records := stringifyTransactions(transactions)

	writer := csv.NewWriter(csvfile)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}
	writer.Flush()
}
