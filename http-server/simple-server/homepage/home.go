package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "hello gopher"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %v \n", time.Now().Sub(startTime))
		next(writer, request)
	}
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home))
}
