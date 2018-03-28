package diffcalculator

import (
	"encoding/json"
	"log"
	"net/http"
)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	body := struct {
		SourceID string
		Items    []Item
	}{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("[ERROR] Error unmarshaling json: %s", err)
		http.Error(w, "error reading body", http.StatusInternalServerError)
		return
	}

	if err = Calculate(body.SourceID, body.Items); err != nil {
		log.Printf("[ERROR] Calculate error: %s", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}

// ServeHTTP create and run HTTP server at :8000
func ServeHTTP() {
	log.Printf("[INFO] Starting HTTP server at :8000")

	http.HandleFunc("/DiffCalculator/Calculate", calculateHandler)
	http.ListenAndServe(":8000", nil)
}
