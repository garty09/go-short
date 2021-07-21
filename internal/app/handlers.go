package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/url"
)

//UrlToggle is the response body
type UrlToggle struct {
	URL string `json:"url"`
}

//Shorten url POST method
func (a *App) Shorten(w http.ResponseWriter, r *http.Request) {
	var id uint64
	var body UrlToggle

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}
	defer r.Body.Close()
	url := body.URL

	if !isValidURL(url) {
		respondWithError(w, http.StatusBadRequest, "Invalid url")
		return
	}

	err := a.db.QueryRow("SELECT id FROM urls WHERE url = $1", url).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		respondWithError(w, http.StatusInternalServerError, "Internal error")
	}

	if id > 0 {
		hash, err := a.h.Encode([]int{int(id)})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal error")
			return
		}
		body.URL = r.Host + "/" + hash
		sendResponse(w, http.StatusOK, body)
	}
	res, err := a.db.Exec("INSERT INTO urls (url, created) VALUES($1, now) RETURNING id;", url)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	rid, _ := res.LastInsertId()
	hash, err := a.h.Encode([]int{int(rid)})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal error")
		return
	}
	body.URL = r.Host + "/" + hash
	sendResponse(w, http.StatusOK, body)
}

// Longer url POST method
func (a *App) Longer(w http.ResponseWriter, r *http.Request) {
	var body UrlToggle

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	u := body.URL

	if !isValidURL(u) {
		respondWithError(w, http.StatusBadRequest, "Invalid url")
		return
	}

	p := url.PathEscape(u)
	ids, err := a.h.DecodeWithError(p)
	if err != nil || len(ids) == 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid url")
		return
	}

	var res UrlToggle
	a.db.QueryRow("SELECT url FROM urls WHERE id = $1", ids[0]).Scan(&res.URL)
	sendResponse(w, http.StatusOK, res)
}

func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	sendResponse(w, code, map[string]string{"error": message})
}

func sendResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
