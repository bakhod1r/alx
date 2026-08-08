# Security Policy

## Supported versions

| Version | Supported |
|---|---|
| 0.1.x | ✅ |

While `alx` is pre-1.0, only the latest minor release receives security fixes.

## Reporting a vulnerability

**Do not open a public issue.**

Report privately through GitHub's
[private vulnerability reporting](https://github.com/bakhod1r/alx/security/advisories/new),
or by email to the address on the [maintainer's GitHub profile](https://github.com/bakhod1r).

Please include:

- the `Accept-Language` header or request that triggers the problem,
- the `alx` version and Go version,
- the impact you observed — panic, hang, excessive allocation, wrong match.

You will get an acknowledgement within 72 hours and an assessment within 7 days. If the
report is confirmed, a fix and an advisory are published together, and you are credited
unless you ask otherwise.

## Scope

`alx` parses attacker-controlled input: the `Accept-Language` header, URL path segments,
query parameters and cookies all arrive from the client. The following are in scope:

- a panic or out-of-range access reachable from any exported function,
- non-linear time or memory growth on a crafted header (algorithmic complexity attacks),
- a match returning a tag that is **not** in the caller's supported list,
- `ParseStrict` or `Validate` accepting a header that violates RFC 9110.

Out of scope:

- lenient parsing dropping malformed segments — that is `Parse`'s documented behaviour;
  use `ParseStrict` if you need rejection,
- vulnerabilities in a web framework used by an adapter module; report those upstream,
- resource exhaustion from a header your own server let through unbounded. `alx` does not
  impose a length limit — cap the header size at your HTTP layer.

## Hardening notes

- `Best` and `MatchWith` return a tag from **your** supported slice, never a
  client-supplied string, so their output is safe to interpolate into file paths or
  template names. Values from `Parse`, `FromRequest` and `Resolution.Tag` are
  client-controlled and must be treated as untrusted.
- Call `SetVary` on any response whose content varies by language, or a shared cache will
  serve one user's language to another.
