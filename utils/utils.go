package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"time"
)

// GeneralResponse defines a standard response format.
type GeneralResponse struct {
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Status    int    `json:"status"`
	ErrorCode any    `json:"errorCode,omitempty"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewSuccessResponse builds a success response.
func NewSuccessResponse(message string, data any, status int) (GeneralResponse, error) {
	if message == "" {
		return GeneralResponse{}, errors.New("message cannot be empty")
	}
	if status < 100 || status > 599 {
		return GeneralResponse{}, errors.New("invalid HTTP status code")
	}
	return GeneralResponse{Message: message, Data: data, Status: status}, nil
}

// NewErrorResponse builds an error response.
func NewErrorResponse(err error, status int) (GeneralResponse, error) {
	if err == nil {
		return GeneralResponse{}, errors.New("error cannot be nil")
	}
	if status < 400 {
		status = http.StatusInternalServerError
	}
	return GeneralResponse{Message: err.Error(), Status: status}, nil
}

// SendJSON writes the response as JSON to the http.ResponseWriter.
func SendJSON(w http.ResponseWriter, resp GeneralResponse, err error) {
	w.Header().Set("Content-Type", "application/json")

	// If an error is passed, override with an error response
	if err != nil {
		errorResp, _ := NewErrorResponse(err, http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResp)
		return
	}

	// Normal success response
	w.WriteHeader(resp.Status)
	json.NewEncoder(w).Encode(resp)
}

func GenerateCSVBase64(filename string, header []string, arrayRow [][]string) string {
	// Create CSV writer
	var b bytes.Buffer
	writer := csv.NewWriter(&b)

	// Write header row
	writer.Write(header)

	// Write student rows
	for _, s := range arrayRow {
		writer.Write(s)
	}

	// Ensure all buffered data is written to the buffer
	writer.Flush()

	encodedCSV := base64.StdEncoding.EncodeToString(b.Bytes())

	return encodedCSV
}

func RandomIntFromInterval(min int, max int) int {
	// min and max included
	return rand.Intn(max-min+1) + min

}

func GetMap(m map[string]any, key string) map[string]any {
	return m[key].(map[string]any)
}
