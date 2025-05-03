package xrpcuri

// SplitAndNormalize returns the 'authority', 'id', 'query', and 'fragment' of an XRPC-URI.
//
// An 'id' is an NSID (Namespaced Identifier).
//
// For example:
//
//	var uri string = "at://Host.EXAMPLE/COM.Example.foorBar"
//	
//	authority, id, query, fragment, err := xrpcuri.SplitAndNormalize(uri)
//	if nil != err {
//		return err
//	}
//	
//	// authority == "host.example"
//	// id        == "com.example.foorBar"
//	// query     == ""
//	// fragment  == ""
//
// SplitAndNormalize normalizes the returned values.
// If you are not sure whether to use [Split] or SpliAndNormalize, use [SplitAndNormalize].
func SplitAndNormalize(uri string) (authority string, id string, query string, fragment string, err error) {
	authority, id, query, fragment, err = Split(uri)

	authority  = NormalizeAuthority(authority)
	id         = NormalizeID(id)

	return
}
