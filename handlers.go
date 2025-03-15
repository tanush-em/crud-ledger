package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func handleEntries(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAllEntries(w, r)
	case "POST":
		addEntry(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleEntry(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		updateEntry(w, r)
	case "DELETE":
		deleteEntry(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAllEntries(w http.ResponseWriter, r *http.Request) {
	ledger, err := loadLedger()
	if err != nil {
		http.Error(w, "Failed to load ledger", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ledger)
}

func addEntry(w http.ResponseWriter, r *http.Request) {
	var newEntry LedgerEntry
	if err := json.NewDecoder(r.Body).Decode(&newEntry); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ledger, err := loadLedger()
	if err != nil {
		http.Error(w, "Failed to load ledger", http.StatusInternalServerError)
		return
	}

	newEntry.ID = len(ledger.Entries) + 1
	ledger.Entries = append(ledger.Entries, newEntry)

	if err := saveLedger(ledger); err != nil {
		http.Error(w, "Failed to save ledger", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEntry)
}

func updateEntry(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedEntry LedgerEntry
	if err := json.NewDecoder(r.Body).Decode(&updatedEntry); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ledger, err := loadLedger()
	if err != nil {
		http.Error(w, "Failed to load ledger", http.StatusInternalServerError)
		return
	}

	found := false
	for i, entry := range ledger.Entries {
		if entry.ID == id {
			ledger.Entries[i] = updatedEntry
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	if err := saveLedger(ledger); err != nil {
		http.Error(w, "Failed to save ledger", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEntry)
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ledger, err := loadLedger()
	if err != nil {
		http.Error(w, "Failed to load ledger", http.StatusInternalServerError)
		return
	}

	newEntries := []LedgerEntry{}
	found := false
	for _, entry := range ledger.Entries {
		if entry.ID == id {
			found = true
			continue
		}
		newEntries = append(newEntries, entry)
	}

	if !found {
		http.Error(w, "Entry not found", http.StatusNotFound)
		return
	}

	ledger.Entries = newEntries

	if err := saveLedger(ledger); err != nil {
		http.Error(w, "Failed to save ledger", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Entry deleted"))
}
