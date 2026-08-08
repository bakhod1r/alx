# Releasing

This repository is a multi-module Go repo. The core module is tagged `vX.Y.Z`; each
adapter module is tagged `<dir>/vX.Y.Z` (Go's required path-prefix tag scheme).

Adapter modules `require github.com/bakhod1r/alx vX.Y.Z` with **no `replace` directive**,
so consumers resolve the published core. Local development uses `go.work` instead —
`go.work` is committed, but it is ignored by the module proxy and never affects consumers.

## Order matters

The core must be tagged and visible on the proxy *before* the adapters are tagged,
otherwise `go get` of an adapter fails with `unknown revision`.

1. Make sure `make all` passes.
2. Move the `Unreleased` entries in [CHANGELOG.md](CHANGELOG.md) under a new
   `## [X.Y.Z] - YYYY-MM-DD` heading and update the link definitions at the bottom.
   The release workflow extracts that exact section as the GitHub release notes and
   fails if it is missing.
3. Tag and push the core:

   ```bash
   git tag v0.1.0
   git push origin v0.1.0
   ```

4. Confirm the proxy has it:

   ```bash
   GOPROXY=proxy.golang.org GOFLAGS=-mod=mod go list -m github.com/bakhod1r/alx@v0.1.0
   ```

5. Bump each adapter's `require github.com/bakhod1r/alx` line to the new core version,
   then tag the adapters:

   ```bash
   for d in gin fiber echo iris beego fasthttp; do
     git tag "$d/v0.1.0"
   done
   git push origin --tags
   ```

6. Verify from a clean directory outside the workspace:

   ```bash
   cd $(mktemp -d) && go mod init tmp && go get github.com/bakhod1r/alx/gin@v0.1.0
   ```

## Version policy

- Pre-`v1.0.0`: breaking changes are allowed on minor bumps.
- From `v1.0.0`: the core module follows semver; breaking changes require a `/v2` path.
- Adapter modules track the core's major version and may bump minor/patch independently.
