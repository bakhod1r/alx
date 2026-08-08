package iris

import (
	"github.com/kataras/iris/v12"
	"net/http/httptest"
	"testing"
)

func TestParser(t *testing.T) {
	app := iris.New()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Lang", "uz-UZ;q=0.9")
	rec := httptest.NewRecorder()

	ctx := app.ContextPool.Acquire(rec, req)
	defer app.ContextPool.Release(ctx)

	if FromRequest(ctx).Tag != "uz-UZ" {
		t.Errorf("FromRequest failed")
	}
	if len(Languages(ctx)) == 0 || Languages(ctx)[0].Tag != "uz-UZ" {
		t.Errorf("Languages failed")
	}
}
