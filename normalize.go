package xrpcuri

// Normalize returns the normalized form of an XRPC-URI.
//
// Normalize does NOT validate the XRPC-URI.
// To validate, call [Validate].
//
// An example of a non-normalized XRPC-URI would be:
//
//	xrpc://VIDEO.archive.ORG/COM.Example.fooBar
//
// Normalizing that non-normalized XRPC-URI would result in:
//
//	xrpc://video.archive.org/com.example.fooBar
func Normalize(uri string) string {

	if err := ValidateScheme(uri); nil != err {
		return uri
	}

	authority, id, query, fragment, err := Split(uri)
	if nil != err {
		return uri
	}

	authority = NormalizeAuthority(authority)
	id        = NormalizeID(id)

	return Join(authority, id, query, fragment)
}
