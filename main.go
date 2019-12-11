package main

import "database/sql"
import "fmt"
import "log"
import "net/http"
import "os"
import _ "github.com/go-sql-driver/mysql"

// @todo #1 /countries/count: add tests
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// @todo #1 Load configuration from a file
	user := os.Getenv("MYSQL_USER")
	if user == "" {
		log.Printf("MYSQL_USER env variable is not set or empty. Defaults to 'mystamps'")
		user = "mystamps"
	}

	pass := os.Getenv("MYSQL_PASSWORD")
	if pass == "" {
		log.Fatalf("MYSQL_PASSWORD env variable is not set or empty")
	}

	dbName := os.Getenv("MYSQL_DB")
	if dbName == "" {
		log.Printf("MYSQL_DB env variable is not set or empty. Defaults to 'mystamps'")
		dbName = "mystamps"
	}

	// @todo #1 Consider passing params to db driver
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", user, pass, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Open() has failed: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ping() has failed: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	// @todo #1 /countries/count: extract handler
	// @todo #1 /countries/count: extract SQL query
	http.HandleFunc("/v0.1/countries/count", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		var count int

		// There is no check for ErrNoRows because COUNT(*) always returns a single row
		err := db.QueryRow("SELECT COUNT(*) FROM countries").Scan(&count)
		if err != nil {
			log.Printf("Scan() has failed: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")

		fmt.Fprintf(w, "%d", count)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Running the server on port %v", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
