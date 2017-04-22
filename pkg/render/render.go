package render

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, v interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if v == nil || code == http.StatusNoContent {
		return
	}

	err = encode(w, v)

	return
}

func encode(w http.ResponseWriter, v interface{}) (err error) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err = enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return
}
