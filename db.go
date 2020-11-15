package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./crawler.db"
const schemaFile = "./schema.sql"

func insertDocumentFrequency(db *sql.DB, DF DocumentFrequency) {
	sql := `INSERT INTO document_frequency(term, occurencies) VALUES (?, ?)`

	for term, occurencies := range DF {
		occJSON, err := json.Marshal(occurencies)
		statement, err := db.Prepare(sql)
		checkErr(err)

		_, err = statement.Exec(term, string(occJSON))
		checkErr(err)
	}

	fmt.Println("Worked")
}

func loadState() State {
	db := openDatabase()
	defer closeDatabase(db)

	return State{
		documents: getDocuments(db),
		DF:        getDocumentFrequency(db),
	}
}

//TODO Update tables instead of flushing all of them and re-inserting everything
//TODO Implement bulk insertions
func saveState(state State) {
	db := openDatabase()
	defer closeDatabase(db)

	emptyAllTables(db)
	insertDocuments(db, state.documents)
	insertDocumentFrequency(db, state.DF)
}

func emptyAllTables(db *sql.DB) {
	tablesToEmpty := []string{"document_frequency", "crawled_document"}

	for _, table := range tablesToEmpty {
		statement, err := db.Prepare("DELETE FROM " + table)
		if err != nil {
			log.Fatal(err.Error())
		}
		statement.Exec()
	}
}

func getDocumentFrequency(db *sql.DB) (DF DocumentFrequency) {
	row, err := db.Query("SELECT * FROM document_frequency")
	checkErr(err)
	defer row.Close()

	DF = make(DocumentFrequency)

	for row.Next() {
		var term string
		var occurencies string
		var occurenciesParsed []string

		row.Scan(&term, &occurencies)
		json.Unmarshal([]byte(occurencies), &occurenciesParsed)

		DF[term] = occurenciesParsed
	}

	return
}

func insertDocuments(db *sql.DB, documents []Document) {
	sql := `INSERT INTO crawled_document(id, url, title, neighbors, termfreq) VALUES (?, ?, ?, ?, ?)`

	for _, document := range documents {
		neighborsJSON, err := json.Marshal(document.neighbors)
		checkErr(err)
		termfreqJSON, err := json.Marshal(document.tf)
		checkErr(err)
		statement, err := db.Prepare(sql)
		checkErr(err)

		_, err = statement.Exec(document.id, document.url, document.title, string(neighborsJSON), string(termfreqJSON))
		checkErr(err)
	}

}

func getDocuments(db *sql.DB) (documents []Document) {
	row, err := db.Query("SELECT * FROM crawled_document")
	checkErr(err)
	defer row.Close()

	for row.Next() {
		var id, url, title string
		var neighbors, termfreq string
		var neighborsParsed []string
		var termFreqParsed map[string]float64

		row.Scan(&id, &url, &title, &neighbors, &termfreq)
		json.Unmarshal([]byte(neighbors), &neighborsParsed)
		json.Unmarshal([]byte(termfreq), &termFreqParsed)

		documents = append(documents, Document{
			id:        id,
			title:     title,
			url:       url,
			neighbors: neighborsParsed,
			tf:        termFreqParsed,
		})
	}

	return
}

func createNewDatabase() *sql.DB {
	deleteDatabase()

	file, err := os.Create(dbName)
	checkErr(err)
	file.Close()

	db := openDatabase()
	migrateDatabase(db)

	return db
}

func deleteDatabase() {
	os.Remove(dbName)
}

func openDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func closeDatabase(db *sql.DB) {
	db.Close()
}

func migrateDatabase(db *sql.DB) {
	schema, err := ioutil.ReadFile(schemaFile)
	commands := strings.Split(string(schema), ";")

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, command := range commands {
		statement, err := db.Prepare(command)
		if err != nil {
			log.Fatal(err.Error())
		}
		statement.Exec()
	}
	log.Println("[DB] Database was migrated")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
