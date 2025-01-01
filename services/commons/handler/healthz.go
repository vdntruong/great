package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthCheck(startedAt time.Time, serviceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uptime := time.Since(startedAt)
		resp := map[string]interface{}{
			"uptime":      uptime.String(),
			"started_at":  startedAt.String(),
			"serviceName": serviceName,
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}
