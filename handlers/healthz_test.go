package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	responseJSON = `{"alive":true}`
)

func TestIndex(t *testing.T) {
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/v1/healthz", strings.NewReader(responseJSON))

	if assert.NoError(t, err) {
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

		if assert.NoError(t, Index(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, responseJSON, rec.Body.String())
		}
	}
}
