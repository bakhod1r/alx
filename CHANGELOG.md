# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Adapter modules (`gin`, `fiber`, `echo`, `iris`, `beego`, `fasthttp`) are tagged
separately as `<dir>/vX.Y.Z` but are released together with the core; see
[RELEASING.md](RELEASING.md).

## [Unreleased]

## [0.1.0] - 2026-08-08

Initial release.

### Added

- **Parsing** — `Parse`, `ParseRaw`, `ParseSorted`, `ParseStrict` and `Validate` for
  `Accept-Language` headers, with quality-value (`q=`) handling, wildcard (`*`) support
  and a `*ParseError` carrying `Position`, `Token` and `Reason`.
- **`Language` type** — decomposed `Tag`, `Language`, `Script`, `Region`, `Quality`,
  with a `String` method that round-trips the `tag;q=value` form.
- **Matching** — `Best`, `Match` and `MatchWith` with the `Exact`, `Basic`, `Lookup`,
  `Fallback` and `PreferRegion` strategies, plus `Fallbacks` for fallback chains
  (`zh-Hant-TW` → `zh-Hant` → `zh`).
- **Normalization** — `Normalize` and `Canonical` for canonical tag casing.
- **Sorting** — `Sort`, stable ordering by descending quality.
- **Resolver** — `NewResolver` with `WithSupported`, `WithDefault`, `WithFallback`,
  `WithURL`, `WithQuery`, `WithCookie` and `WithHeader`; `Resolve` (header string) and
  `ResolveRequest` (`*http.Request`) return a `Resolution` reporting which `Source`
  the language came from.
- **`net/http` helpers** — `FromRequest`, `Languages`, `PrimaryLanguage` (allocation-free
  fast path), `SetContentLanguage`, `SetVary`.
- **Framework adapters** — `gin`, `fiber`, `echo`, `iris`, `beego` and `fasthttp`
  submodules, each exposing `FromRequest` and `Languages` for that framework's context.
  Each is its own Go module, so the core package stays dependency-free.
- MIT license, package documentation, contributing and security policies, CI across
  Go 1.24–1.26 with race detection and golangci-lint, and 100% statement coverage
  across every module.

[Unreleased]: https://github.com/bakhod1r/alx/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/bakhod1r/alx/releases/tag/v0.1.0
