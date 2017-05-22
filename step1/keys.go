package keys

// GetKeys client tokens must be signed by one of the server keys provided via a url.
// The keys expire after a certain amount of time so we need to track that also.
func GetKeys(keyURL string) (tokens map[string]interface{}, maxAge int64, err error) {
	return nil, 0, nil
}

// Extract the max age from the cache control response header value
// The cache control header should look similar to "..., max-age=19008, ..."
func extractMaxAge(cacheControl string) (int64, error) {
	return 0, nil
}
