package xrpcuri_internal

// ResolveID resolves the 'id' (which is expected to be an NSID) into an HTTP path.
//
// If `id` is "" then it returns "".
func ResolveID(id string) string {
	if "" == id {
		return ""
	}

	id = NormalizeID(id)

	return  "/xrpc/" + encodeSolidus(encodeQuestionMark(encodeNumberSign(encodePercentSign(id))))
}
