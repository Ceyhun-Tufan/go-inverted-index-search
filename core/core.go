package core

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var punctuationRegex = regexp.MustCompile(`[^\p{L}\p{N}\s]+`) // I hate you

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./search.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS inverted_index (
		token TEXT NOT NULL,
		doc_id INTEGER NOT NULL,
		PRIMARY KEY (token, doc_id)
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error creating table: %q: %s\n", err, sqlStmt)
	}
}

func InsertIndex(stringsTokenized *[]string, id int) error {

	if stringsTokenized == nil {
		return fmt.Errorf("tokenized strings are nil")
	}

	var seen = make(map[string]bool)

	var values []string

	for _, token := range *stringsTokenized {

		if seen[token] {
			continue
		}
		seen[token] = true

		values = append(values, fmt.Sprintf("('%s',%d)", token, id))

	}

	sqlStmt := fmt.Sprintf("INSERT OR IGNORE INTO inverted_index (token, doc_id) VALUES\n %s;", strings.Join(values, ",\n "))
	_, err := DB.Exec(sqlStmt)
	return err

}

func Tokenize(val *string) []string {

	if val == nil || *val == "" {
		return []string{}
	}

	// lower casing and deleting unnecessary parts of the string
	lowered := strings.ToLower(*val)
	cleaned := punctuationRegex.ReplaceAllString(lowered, "")

	return strings.Fields(cleaned)
}

func Search(val string) []string {

	rows, err := DB.Query(fmt.Sprintf("SELECT doc_id FROM inverted_index WHERE token == '%s'", val))
	if err != nil {
		log.Println("Query error:", err)
		return nil
	}
	defer rows.Close()

	var docIDs []string
	for rows.Next() {
		var docID int
		if err := rows.Scan(&docID); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		docIDs = append(docIDs, fmt.Sprintf("%d", docID))
	}

	if len(docIDs) > 0 {
		return docIDs
	}

	return nil
}
