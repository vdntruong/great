package json

import (
	"encoding/json"
	"net/http"
)

func DecodeRequest(r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(&data)
}
