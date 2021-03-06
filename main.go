package main

import "database/sql"
import "fmt"
import "log"
import "net/http"
import "os"
import _ "github.com/go-sql-driver/mysql"
import "github.com/php-coder/mystamps-country/db"
import "github.com/php-coder/mystamps-country/rest"

// @todo #1 /countries/count: add integration tests
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

	mysql, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Open() has failed: %v", err)
	}

	err = mysql.Ping()
	if err != nil {
		log.Fatalf("Ping() has failed: %v", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("ok"))
	})

	dao := db.New(mysql)
	rest.New(dao).Register(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Running the server on port %v", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
