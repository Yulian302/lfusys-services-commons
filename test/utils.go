package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func PerformRequest(r *gin.Engine, t *testing.T, method, url string, body io.Reader, headers []string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	for _, h := range headers {
		splittedStrings := strings.Split(h, ":")
		if len(splittedStrings) != 2 {
			break
		}
		req.Header.Set(strings.Trim(splittedStrings[0], " "), strings.Trim(splittedStrings[1], " "))
	}

	r.ServeHTTP(w, req)
	return w
}
