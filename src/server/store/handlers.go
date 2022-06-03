package store

import (
	"net/http"

	"github.com/luisnquin/meow-app/src/server/log"
)

func (m *database) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")

		if err := m.db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Error(err)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("connection alive"))
	}
}
