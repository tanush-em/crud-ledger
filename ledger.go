package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LedgerEntry struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
	Note   string  `json:"note"`
}

type Ledger struct {
	Entries []LedgerEntry `json:"entries"`
}

const ledgerFile = "ledger.json"

func loadLedger() (*Ledger, error) {
	file, err := os.Open(ledgerFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Ledger{}, nil
		}
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ledger Ledger
	if err := json.Unmarshal(data, &ledger); err != nil {
		return nil, err
	}

	return &ledger, nil
}

func saveLedger(ledger *Ledger) error {
	data, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ledgerFile, data, 0644)
}
