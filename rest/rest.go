package rest

import "database/sql"
import "fmt"
import "log"
import "net/http"

type rest struct {
	db *sql.DB
}

func New(db *sql.DB) *rest {
	return &rest{
		db: db,
	}
}

func (r *rest) Register(mux *http.ServeMux) {
	mux.HandleFunc("/v0.1/countries/count", r.countHandler)
}

// @todo #1 /countries/count: extract SQL query
func (r *rest) countHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var count int

	// There is no check for ErrNoRows because COUNT(*) always returns a single row
	err := r.db.QueryRow("SELECT COUNT(*) FROM countries").Scan(&count)
	if err != nil {
		log.Printf("Scan() has failed: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, "%d", count)
}
