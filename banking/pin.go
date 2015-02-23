package banking

import (
	"encoding/json"
	"log"
	"os"

	aqbanking "github.com/umsatz/go-aqbanking"
)

type pin struct {
	Blz string `json:"blz"`
	UID string `json:"uid"`
	PIN string `json:"pin"`
}

func (p *pin) BankCode() string {
	return p.Blz
}

func (p *pin) UserID() string {
	return p.UID
}

func (p *pin) Pin() string {
	return p.PIN
}

func LoadPins(filename string) []aqbanking.Pin {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("%v", err)
		return nil
	}

	var _pins []pin
	if err = json.NewDecoder(f).Decode(&_pins); err != nil {
		log.Fatal("%v", err)
		return nil
	}

	var pins = make([]aqbanking.Pin, len(_pins))
	for i, pin := range _pins {
		pins[i] = aqbanking.Pin(&pin)
	}

	return pins
}
