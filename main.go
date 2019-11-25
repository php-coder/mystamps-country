package main

import "log"
import "net/http"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("Running the server on port 8080")

	http.ListenAndServe(":8080", nil)
}
