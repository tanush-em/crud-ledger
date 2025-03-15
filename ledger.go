package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PlayerEntry struct {
	ID     int     `json:"id"`
	Player string  `json:"player"`
	Role   string  `json:"role"`  // BAT, BALL, AR
	Price  float64 `json:"price"` // Auction price
	Year   int     `json:"year"`  // Auction year
}

type PlayerLedger struct {
	Entries []PlayerEntry `json:"entries"`
}

const ledgerFile = "player_ledger.json"

// Load the ledger from file
func loadLedger() (*PlayerLedger, error) {
	file, err := os.Open(ledgerFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &PlayerLedger{}, nil
		}
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var ledger PlayerLedger
	if err := json.Unmarshal(data, &ledger); err != nil {
		return nil, err
	}

	return &ledger, nil
}

// Save the ledger to file
func saveLedger(ledger *PlayerLedger) error {
	data, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ledgerFile, data, 0644)
}
