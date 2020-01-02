package rest

import "fmt"
import "log"
import "net/http"
import "github.com/php-coder/mystamps-country/db"

type rest struct {
	db db.CountryDB
}

func New(db db.CountryDB) *rest {
	return &rest{
		db: db,
	}
}

func (r *rest) Register(mux *http.ServeMux) {
	mux.HandleFunc("/v0.1/countries/count", r.countHandler)
}

func (r *rest) countHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	count, err := r.db.CountAll()
	if err != nil {
		log.Printf("CountAll() has failed: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintf(w, "%d", count)
}
