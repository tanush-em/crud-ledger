# JSON Ledger CRUD API

This is a simple CRUD API built using Go, designed to manage player records stored in a JSON file.

## Features
- Add a player entry
- Retrieve all player entries
- Update a player entry
- Delete a player entry

## Data Format
Each player entry consists of:
- **ID** (int) - Unique identifier
- **Player** (string) - Player name
- **Role** (string) - BAT/BALL/AR (Batsman/Bowler/All-rounder)
- **Price** (float) - Auction price
- **Year** (int) - Year of the auction

## Running the Project
1. Install Go if not already installed.
2. Clone the repository.
3. Run the project:
   ```sh
   go run main.go
   ```
4. Use Postman or cURL to interact with the API.

## API Endpoints
- `GET /entries` - Retrieve all player entries
- `POST /entries` - Add a new player entry
- `PUT /entry?id={id}` - Update an existing player entry
- `DELETE /entry?id={id}` - Delete a player entry

## Dependencies
- Go standard library (no external packages required)

## Notes
- Data is stored in `ledger.json`
- No database is required

---

This project is a simple and lightweight implementation for managing structured JSON data using a RESTful API.

