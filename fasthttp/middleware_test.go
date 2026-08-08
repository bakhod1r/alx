package fasthttp

import (
	"github.com/valyala/fasthttp"
	"testing"
)

func TestParser(t *testing.T) {
	ctx := &fasthttp.RequestCtx{}
	ctx.Request.Header.Set("Lang", "uz-UZ;q=0.9")

	if FromRequest(ctx).Tag != "uz-UZ" {
		t.Errorf("FromRequest failed")
	}
	if len(Languages(ctx)) == 0 || Languages(ctx)[0].Tag != "uz-UZ" {
		t.Errorf("Languages failed")
	}

	emptyCtx := &fasthttp.RequestCtx{}
	if FromRequest(emptyCtx).Tag != "" {
		t.Errorf("FromRequest empty failed")
	}
	if len(Languages(emptyCtx)) != 0 {
		t.Errorf("Languages empty failed")
	}
}
