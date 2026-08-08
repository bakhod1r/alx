# alx

**Parse the header once. `alx` tells you which language to serve.**

[![Go Reference](https://pkg.go.dev/badge/github.com/bakhod1r/alx.svg)](https://pkg.go.dev/github.com/bakhod1r/alx)
[![Tests](https://github.com/bakhod1r/alx/actions/workflows/test.yml/badge.svg)](https://github.com/bakhod1r/alx/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bakhod1r/alx)](https://goreportcard.com/report/github.com/bakhod1r/alx)
[![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](#development)
[![Zero dependencies](https://img.shields.io/badge/dependencies-zero-blue)](go.mod)
[![Go version](https://img.shields.io/badge/go-%3E%3D1.21-00ADD8)](go.mod)
[![MIT License](https://img.shields.io/badge/license-MIT-yellow.svg)](LICENSE)

`alx` is an `Accept-Language` toolkit for Go (RFC 9110 / RFC 4647), with zero third-party
dependencies. It gives you two layers of usage:

- **Direct functions** — parse, sort, normalize and match a header in one call.
- **Resolver** — decide the request's language across URL, query, cookie, header and
  default, and tell you which one won.

On top of either layer you get quality-value ordering, wildcards, fallback chains,
canonical tag casing, and framework adapters that stay out of your dependency graph.

## Installation

```bash
go get github.com/bakhod1r/alx
```

## 🚀 The killer feature: one line at the edge

Content negotiation is usually a hand-rolled `strings.Split` that quietly ignores
`q=` values and serves the wrong language. It is one call:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    tag, ok := alx.Best(r.Header.Get("Accept-Language"), []string{"en", "uz", "ru"})
    if !ok {
        tag = "en"
    }
    alx.SetContentLanguage(w, tag)
    alx.SetVary(w) // so caches key on the header you just read
}
```

`Best` sorts by quality, walks the fallback chain (`zh-Hant-TW` → `zh-Hant` → `zh`),
honours `*`, and returns your supported tag — spelled exactly as you wrote it, not as
the client sent it.

## 🌍 Resolver: the language is not only in the header

Real apps read the language from a URL prefix, a `?lang=` link, a cookie the user set
last week, and only then the browser. The `Resolver` does that in one place, in the
order you declare:

```go
resolver := alx.NewResolver(
    alx.WithSupported("en", "uz", "ru"),
    alx.WithURL(),            // /uz/products
    alx.WithQuery("lang"),    // ?lang=uz
    alx.WithCookie("lang"),
    alx.WithHeader(),         // Accept-Language, or the headers you name
    alx.WithFallback("uz-Cyrl", "ru"),
    alx.WithDefault("en"),
)

res := resolver.ResolveRequest(r)
```

A `Resolution` reports the decision *and* its provenance:

```go
res.Tag      // "uz"
res.Language // "uz"
res.Script   // "" — e.g. "Hant"
res.Region   // "" — e.g. "UZ"
res.Quality  // 1.0
res.Matched  // the supported tag that matched
res.Source   // SourceURL | SourceQuery | SourceCookie | SourceHeader | SourceDefault
```

`res.Source` is the field that makes this debuggable: when a user insists the site is in
the wrong language, you can log whether it came from their cookie or their browser.

For a bare header string, use `resolver.Resolve(header)`.

## 🔤 Parsing

```go
langs := alx.ParseSorted("uz-UZ,uz;q=0.9,en;q=0.8,*;q=0.1")

langs[0].Tag      // "uz-UZ"
langs[0].Language // "uz"
langs[0].Region   // "UZ"
langs[0].Script   // ""
langs[0].Quality  // 1.0
langs[0].String() // "uz-UZ"  — round-trips as "uz;q=0.9" when q != 1
```

| Function | Behaviour |
|---|---|
| `Parse` / `ParseRaw` | Lenient. Header order preserved, malformed segments dropped. |
| `ParseSorted` | Lenient, then sorted by descending quality. |
| `ParseStrict` | Returns `error` on any malformed segment. |
| `Validate` | `ParseStrict` with the result discarded — error only. |
| `PrimaryLanguage` | Allocation-free fast path: first tag, no `q=` handling. |
| `Sort` | Stable sort of `[]Language` by descending quality. |

Strict mode returns a `*ParseError` that says where it broke:

```go
_, err := alx.ParseStrict("en;q=banana")
var pe *alx.ParseError
if errors.As(err, &pe) {
    pe.Position // byte offset in the header
    pe.Token    // the offending token
    pe.Reason   // why it was rejected
}
```

Equal quality values keep header order, so a client's own preference ordering survives —
that is what makes the sort stable rather than merely correct.

## 🎯 Matching strategies

```go
alx.Best(header, supported)                    // Lookup — the sensible default
alx.Match(header, supported)                   // alias for Best
alx.MatchWith(header, supported, alx.Exact)    // tag must match exactly
alx.MatchWith(header, supported, alx.Basic)    // base language only: en-US matches en
alx.MatchWith(header, supported, alx.Lookup)   // fallback chain: zh-Hant-TW → zh-Hant → zh
```

| Strategy | `en-GB` against `["en"]` | `zh-Hant-TW` against `["zh"]` |
|---|---|---|
| `Exact` | no match | no match |
| `Basic` | `en` | `zh` |
| `Lookup` | `en` | `zh` |

`Fallback` is an alias for `Lookup`; `PreferRegion` currently behaves as `Lookup`.

The chain itself is exported, so you can drive your own lookup — translation files,
a CDN path, whatever:

```go
alx.Fallbacks("zh-Hant-TW") // ["zh-Hant-TW", "zh-Hant", "zh"]
```

## 🧹 Normalization

Clients send `EN-us`, config files hold `en_US`, your file names use `en-US`. One
function each way:

```go
alx.Normalize("ZH-hant-tw") // "zh-Hant-TW"  — language lower, script Title, region UPPER
alx.Canonical("en_US")      // "en-US"       — underscores to dashes, then Normalize
```

## 🌐 Framework adapters

Each adapter is **its own Go module**, so importing `alx` never pulls a web framework
into your build. Install only the one you use:

| Framework | Import |
|---|---|
| Gin | `go get github.com/bakhod1r/alx/gin` |
| Fiber | `go get github.com/bakhod1r/alx/fiber` |
| Echo | `go get github.com/bakhod1r/alx/echo` |
| Iris | `go get github.com/bakhod1r/alx/iris` |
| Beego | `go get github.com/bakhod1r/alx/beego` |
| fasthttp | `go get github.com/bakhod1r/alx/fasthttp` |

Every adapter exposes the same two functions, typed for that framework's context:

```go
import alxgin "github.com/bakhod1r/alx/gin"

func handler(c *gin.Context) {
    lang := alxgin.FromRequest(c) // best Language
    all := alxgin.Languages(c)    // all languages, sorted by quality
}
```

```go
import alxfiber "github.com/bakhod1r/alx/fiber"

func handler(c *fiber.Ctx) error {
    lang := alxfiber.FromRequest(c)
    return c.SendString(lang.Tag)
}
```

## 🔌 net/http helpers

```go
alx.FromRequest(r)      // highest-quality Language
alx.Languages(r)        // all languages, sorted
alx.SetContentLanguage(w, "uz")
alx.SetVary(w)
```

`FromRequest` and `Languages` check `Accept-Language`, then `Language`, then `Lang` —
some clients and proxies use the shorter spellings. Both return zero values on a `nil`
request rather than panicking, so they are safe in middleware that runs before routing.

## Development

The repository is a Go workspace, so the adapter modules build against the local core:

```bash
make test   # go test -cover across every module
make race   # go test -race
make cover  # coverage report for the core
make all    # fmt + vet + test
```

Contributions welcome — see [CONTRIBUTING.md](CONTRIBUTING.md). Releases are documented
in [RELEASING.md](RELEASING.md), changes in [CHANGELOG.md](CHANGELOG.md).

## License

MIT © bakhod1r
