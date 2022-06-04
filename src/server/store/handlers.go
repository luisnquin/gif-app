package store

import (
	"bufio"
	"net/http"
	"os/exec"

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

func (m *database) AutoMockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd := exec.CommandContext(r.Context(), "./tools/automock/main.py", "--stdout", "--length=10")
		pipe, err := cmd.StdoutPipe()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		defer pipe.Close()

		err = cmd.Start()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		stmts := make([]string, 0)

		s := bufio.NewScanner(pipe)
		for s.Scan() {
			stmts = append(stmts, s.Text())
		}

		err = cmd.Wait()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
