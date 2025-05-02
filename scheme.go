package xrpcuri

const (
	Scheme            = "xrpc"             // example: "xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social"
	SchemeUnencrypted = "xrpc-unencrypted" // example: "xrpc-unencrypted://localhost/app.bsky.actor.getProfile?actor=reiver.bsky.social"
)

const (
	prefixScheme            = Scheme            + ":"
	prefixSchemeUnencrypted = SchemeUnencrypted + ":"
)

const (
	lenPrefixScheme            int = len(prefixScheme)
	lenPrefixSchemeUnencrypted int = len(prefixSchemeUnencrypted)
)
