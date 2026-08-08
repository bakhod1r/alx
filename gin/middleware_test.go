package gin

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestParser(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/", nil)
	c.Request.Header.Set("Language", "uz-UZ;q=0.9")

	if FromRequest(c).Tag != "uz-UZ" {
		t.Errorf("FromRequest failed")
	}
	if len(Languages(c)) == 0 || Languages(c)[0].Tag != "uz-UZ" {
		t.Errorf("Languages failed")
	}
}
