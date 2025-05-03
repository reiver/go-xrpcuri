package xrpcuri_internal

const (
	Scheme            = "xrpc"             // example: "xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social"
	SchemeUnencrypted = "xrpc-unencrypted" // example: "xrpc-unencrypted://localhost/app.bsky.actor.getProfile?actor=reiver.bsky.social"
)

const (
	PrefixScheme            = Scheme            + ":"
	PrefixSchemeUnencrypted = SchemeUnencrypted + ":"
)

const (
	LenPrefixScheme            int = len(PrefixScheme)
	LenPrefixSchemeUnencrypted int = len(PrefixSchemeUnencrypted)
)
