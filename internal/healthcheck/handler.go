package healthcheck

import (
	"net/http"
)

func GetStatus(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.WriteHeader(http.StatusOK)
	statusResponse := "healthy"
	_, _ = responseWriter.Write([]byte(statusResponse))
}
