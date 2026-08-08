package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestParser(t *testing.T) {
	app := fiber.New()
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	ctx.Request().Header.Set("Lang", "uz-UZ;q=0.9")

	if FromRequest(ctx).Tag != "uz-UZ" {
		t.Errorf("FromRequest failed")
	}
	if len(Languages(ctx)) == 0 || Languages(ctx)[0].Tag != "uz-UZ" {
		t.Errorf("Languages failed")
	}

	emptyCtx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(emptyCtx)
	if FromRequest(emptyCtx).Tag != "" {
		t.Errorf("FromRequest empty failed")
	}
	if len(Languages(emptyCtx)) != 0 {
		t.Errorf("Languages empty failed")
	}
}
