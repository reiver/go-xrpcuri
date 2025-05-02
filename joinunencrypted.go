package xrpcuri

func JoinUnencrypted(host string, id string, query string, fragment string) string {

	const scheme string = SchemeUnencrypted

	return join(scheme, host, id, query, fragment)
}
