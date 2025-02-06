package middlewares

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	s "app_store_api/services"
)

type (
	Middleware struct{}

	ErrorObject struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

var (
	LimiterExpiration = 30 * time.Second
	RequestID         = middleware.RequestID
	RealIP            = middleware.RealIP
	Logger            = middleware.Logger
	Recoverer         = middleware.Recoverer
	GetHead           = middleware.GetHead
	HttpRate          = httprate.LimitByIP(25, LimiterExpiration)

	HttpOrigins          = []string{"https://*", "http://*"}
	HttpMethods          = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	HttpHeaders          = []string{"Accept", "Authorization", "Content-Type", "X-Requested-With", "X-CSRF-Token"}
	HttpExposedHeaders   = []string{"Link"}
	HttpAllowCredentials = false
	HttpMaxAge           = 300
)

func (m Middleware) Authenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			errObj := &ErrorObject{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized",
			}

			jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(jsonMarshal))

			return
		}

		t := strings.Split(auth, " ")[1]

		if err := s.VerifyJWT(t, w); err != nil {
			errObj := &ErrorObject{
				Code:    http.StatusUnauthorized,
				Message: "Token Expired",
			}

			jsonMarshal, _ := json.MarshalIndent(errObj, "", "  ")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(jsonMarshal))

			return
		}

		h.ServeHTTP(w, r)
	})
}

func (Middleware) UseMiddlewares(r *chi.Mux) {
	r.Use(RequestID)
	r.Use(RealIP)
	r.Use(Logger)
	r.Use(Recoverer)
	r.Use(GetHead)
	r.Use(HttpRate)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   HttpOrigins,
		AllowedMethods:   HttpMethods,
		AllowedHeaders:   HttpHeaders,
		ExposedHeaders:   HttpExposedHeaders,
		AllowCredentials: HttpAllowCredentials,
		MaxAge:           HttpMaxAge,
	}))
}
