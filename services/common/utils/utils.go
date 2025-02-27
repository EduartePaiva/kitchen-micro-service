package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, value any) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("Content-Type header is not json")
	}
	decode := json.NewDecoder(r.Body)
	return decode.Decode(value)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.Write([]byte("{\"error\":\"" + err.Error() + "\"}"))
	w.WriteHeader(status)
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Add("Content-Type", "application/json")
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
	w.WriteHeader(status)
}
