package helpers

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "collections.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS resources (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Insérer des données de test
	insertTestData(db)

	return db
}

func insertTestData(db *sql.DB) {
	testData := []struct {
		id   string
		name string
		url  string
	}{
		{"1", "Cours Informatique", "https://example.com/informatique"},
		{"2", "Mathématiques Avancées", "https://example.com/maths"},
		{"3", "Physique Quantique", "https://example.com/physique"},
	}

	for _, data := range testData {
		_, err := db.Exec("INSERT OR IGNORE INTO resources (id, name, url) VALUES (?, ?, ?)", data.id, data.name, data.url)
		if err != nil {
			log.Printf("Erreur lors de l'insertion : %v\n", err)
		}
	}
	log.Println("✅ Données de test insérées avec succès !")
}
