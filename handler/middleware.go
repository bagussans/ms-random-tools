package handler

import (
	"errors"
	"net/http"
	"os"

	"github.com/bagussans/ms-support-golang/utils"
)

type AuthResponse struct {
	Valid   bool   `json:"valid"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type ErrorResp struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func AuthCheckApiKey() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// api_key := os.Getenv("MS_AUTH_API_KEY")
			api_key := r.Header.Get("X-API-KEY")

			if api_key == "" {
				resp, err := utils.NewErrorResponse(errors.New("Unauthorized E-01"), http.StatusForbidden)
				utils.SendJSON(w, resp, err)
				return
			} else {
				if api_key != os.Getenv("API_KEY") {
					resp, err := utils.NewErrorResponse(errors.New("Unauthorized E-02"), http.StatusForbidden)
					utils.SendJSON(w, resp, err)
					return
				} else {
					next.ServeHTTP(w, r.WithContext(r.Context()))
				}
			}
		})
	}
}
