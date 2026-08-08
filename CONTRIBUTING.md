# Contributing to alx

Thanks for taking the time to contribute. This document covers what you need to know
before opening an issue or a pull request.

## Ground rules

- **The core module has zero third-party dependencies, and it stays that way.**
  `go.mod` at the repository root must never gain a `require` block. Anything that needs
  a dependency belongs in an adapter module, or nowhere.
- **Coverage is 100% and CI enforces the tests.** New code needs new tests, including the
  failure paths — malformed headers, `nil` requests, empty supported lists.
- Public API changes need a `CHANGELOG.md` entry under `## [Unreleased]`.

## Getting set up

The repository is a multi-module Go workspace, so the adapter modules build against your
local core without any `replace` directives:

```bash
git clone https://github.com/bakhod1r/alx
cd alx
make all    # fmt + vet + test across every module
```

| Target | What it runs |
|---|---|
| `make test` | `go test -cover ./...` in every module |
| `make race` | the same under `-race` |
| `make cover` | coverage profile and total for the core |
| `make fmt` | `go fmt` everywhere |
| `make fmt-check` | fails if anything is not gofmt'd (this is what CI runs) |
| `make vet` | `go vet` everywhere |
| `make staticcheck` | installs and runs staticcheck |
| `make tidy` | `go mod tidy` per module |

CI runs `fmt-check`, `vet` and `race` on Go 1.24, 1.25 and 1.26. Run `make all` before
pushing and you will not be surprised.

## Pull requests

1. Branch from `main`.
2. Keep the change focused — one concern per PR.
3. Use [Conventional Commits](https://www.conventionalcommits.org/) for the title:
   `feat:`, `fix:`, `docs:`, `test:`, `refactor:`, `chore:`. The type drives the
   changelog section the change lands in.
4. Describe the behaviour change, not the diff. If it fixes a bug, include the header
   string that used to be handled wrongly.
5. Green CI is required before review.

## Adding a framework adapter

Adapters are deliberately thin — two functions, no middleware, no framework-specific
configuration:

```go
func FromRequest(c *Framework.Context) alx.Language
func Languages(c *Framework.Context) []alx.Language
```

To add one:

1. Create a directory named after the framework with its own `go.mod`
   (`module github.com/bakhod1r/alx/<name>`), requiring the current released core
   version. **Do not add a `replace` directive** — local builds resolve through
   `go.work`, and a committed `replace` would break `go get` for everyone else.
2. Add the directory to `go.work`, to `MODULES` in the `Makefile`, and to the
   release loop in `RELEASING.md`.
3. Cover both functions, including the case where the framework hands you a request
   with no `Accept-Language` header at all.

Adapters that would require a framework fork, or that only wrap what `alx` already
exports, will be declined — the point of the split modules is that nobody pays for a
framework they are not using.

## Reporting bugs

Open an issue with the exact `Accept-Language` header, the supported list, what you got,
and what you expected. Header parsing bugs are almost always reproducible in three lines,
and a failing test case is the fastest possible bug report.

## Security

Do not open a public issue for a security problem — see [SECURITY.md](SECURITY.md).

## Code of conduct

This project follows the [Contributor Covenant](CODE_OF_CONDUCT.md).

## License

By contributing, you agree that your contributions are licensed under the
[MIT License](LICENSE).
