package response

import (
	"encoding/json"
	"net/http"
)

func CreateJSONResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
