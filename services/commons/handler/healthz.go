package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func HealthCheck(startedAt time.Time, serviceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var uptime = time.Since(startedAt)
		var data = map[string]interface{}{
			"uptime":       uptime.String(),
			"started_at":   startedAt.String(),
			"service_name": serviceName,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}
