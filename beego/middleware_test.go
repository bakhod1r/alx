package beego

import (
	"github.com/beego/beego/v2/server/web/context"
	"net/http/httptest"
	"testing"
)

func TestParser(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Lang", "uz-UZ;q=0.9")
	rec := httptest.NewRecorder()

	c := context.NewContext()
	c.Reset(rec, req)

	if FromRequest(c).Tag != "uz-UZ" {
		t.Errorf("FromRequest failed")
	}
	if len(Languages(c)) == 0 || Languages(c)[0].Tag != "uz-UZ" {
		t.Errorf("Languages failed")
	}
}
