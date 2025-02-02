package helpers

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

// ConnectDB initialise la base de données SQLite
func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "collections.db")
	if err != nil {
		return nil, err
	}

	// Création des tables si elles n'existent pas
	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// createTables crée les tables nécessaires
func createTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		start_time DATETIME NOT NULL,
		end_time DATETIME NOT NULL,
		location TEXT NOT NULL,
		resource_id TEXT NOT NULL,
		FOREIGN KEY(resource_id) REFERENCES resources(id)
	);

	CREATE TABLE IF NOT EXISTS resources (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
	    url TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS alerts (
		id TEXT PRIMARY KEY,
		recipient TEXT NOT NULL,
		resource_id TEXT NOT NULL,
		alert_type TEXT NOT NULL,
		FOREIGN KEY(resource_id) REFERENCES resources(id)
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Erreur lors de la création des tables : %v", err)
		return err
	}

	fmt.Println("✅ Base de données et tables créées avec succès !")
	return nil
}
