// Package alx parses, sorts and matches HTTP Accept-Language headers (RFC 9110 / RFC 4647).
//
// The core package depends only on the standard library. Framework adapters live in
// separate submodules (github.com/bakhod1r/alx/gin, /fiber, /echo, /iris, /beego, /fasthttp)
// so importing alx never pulls a web framework into your build.
//
// Basic matching:
//
//	tag, ok := alx.Best(r.Header.Get("Accept-Language"), []string{"en", "uz", "ru"})
//
// Multi-source resolution (URL segment, query, cookie, header, default):
//
//	res := alx.NewResolver(
//		alx.WithSupported("en", "uz", "ru"),
//		alx.WithURL(),
//		alx.WithQuery("lang"),
//		alx.WithCookie("lang"),
//		alx.WithHeader(),
//		alx.WithDefault("en"),
//	).ResolveRequest(r)
package alx
