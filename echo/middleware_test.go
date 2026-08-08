package echo

import (
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"testing"
)

func TestParser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Lang", "uz-UZ;q=0.9")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if FromRequest(c).Tag != "uz-UZ" {
		t.Errorf("FromRequest failed")
	}
	if len(Languages(c)) == 0 || Languages(c)[0].Tag != "uz-UZ" {
		t.Errorf("Languages failed")
	}
}
