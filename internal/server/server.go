package server

import (
	"go-prac/internal/router"
	"log/slog"
	"net/http"
	"strings"
)

func Start(port string) {
	masterMux := http.NewServeMux()

	masterMux.HandleFunc("/health", healthHandler)
	masterMux.HandleFunc("/", homeHandler)
	masterMux.Handle("/user/", http.StripPrefix("/user", router.UserRoutes()))

	slog.Info("Server is starting", "port", strings.TrimPrefix(port, ":"))

	if err := http.ListenAndServe(port, masterMux); err != nil {
		slog.Error("Server failed to start", "error", err)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Welcome 👋"}`))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
