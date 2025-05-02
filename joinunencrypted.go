package xrpcuri

func JoinUnencrypted(host string, collection string, query string, fragment string) string {

	const scheme string = SchemeUnencrypted

	return join(scheme, host, collection, query, fragment)
}
