# go-xrpcuri

Package **xrpcuri** provides an implementation of the **XRPC-URI** for **BlueSky** and its **AT-Protocol**, for the Go programming language.

XRPC is defined here:
https://atproto.com/specs/xrpc

The **XRPC-URI** is not part of the https://atproto.com/specs/xrpc specification, but is instead something invented by this package.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-xrpcuri

[![GoDoc](https://godoc.org/github.com/reiver/go-xrpcuri?status.svg)](https://godoc.org/github.com/reiver/go-xrpcuri)

## XRPC URI Resolution

This package introduces two URI schemes for XRPC:

* `xrpc`, and
* `xrpc-unencrypted`.

For example:

* `xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social`
* `xrpc-unencrypted://localhost/app.bsky.actor.getProfile?actor=reiver.bsky.social`

Behind-the-scenes the scenes, these XRPC URI schemes are resolved to HTTPS, HTTP, WSS, and WS URLs.

MOST DEVELOPERS WHO ARE JUST MAKING XRPC CLIENT REQUESTS DO NOT HAVE TO WORRY ABOUT THE DETAILS OF THE RESOLUTION.
IN THE SAME WAY THAT MOST DEVELOPERS DO NOT HAVE TO WORRY ABOUT HOW HTTP URIS ARE RESOLVED TO TCP.

How an XRPC URI gets resolved to an HTTPS, HTTP, WSS, or WS URL, depends on the XRPC request type.

Here are some examples:

| XRPC URI                                          | XRPC Request Type | Resolved URL                                 |
|---------------------------------------------------|-------------------|---------------------------------------------:|
| `xrpc://example.com/app.cherry.fooBar`            | `execute`         | `https://example.com/xrpc/app.cherry.fooBar` |
| `xrpc://example.com/app.cherry.fooBar`            | `query`           | `https://example.com/xrpc/app.cherry.fooBar` |
| `xrpc://example.com/app.cherry.fooBar`            | `subscribe`       |   `wss://example.com/xrpc/app.cherry.fooBar` |
| `xrpc-unencrypted://localhost/link.banana.bazQux` | `execute`         |   `http://localhost/xrpc/link.banana.bazQux` |
| `xrpc-unencrypted://localhost/link.banana.bazQux` | `query`           |   `http://localhost/xrpc/link.banana.bazQux` |
| `xrpc-unencrypted://localhost/link.banana.bazQux` | `subscribe`       |     `ws://localhost/xrpc/link.banana.bazQux` |

## Import

To import package **xrpcuri** use `import` code like the follownig:
```
import "github.com/reiver/go-xrpcuri"
```

## Installation

To install package **xrpcuri** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-xrpcuri
```

## Author

Package **xrpcuri** was written by [Charles Iliya Krempeaux](http://reiver.link)

## See Also

* https://github.com/reiver/go-athandle
* https://github.com/reiver/go-atproto
* https://github.com/reiver/go-aturi
* https://github.com/reiver/go-bsky
* https://github.com/reiver/go-did
* https://github.com/reiver/go-didplc
* https://github.com/reiver/go-nsid
* https://github.com/reiver/go-xrpc
* https://github.com/reiver/go-xrpcuri
