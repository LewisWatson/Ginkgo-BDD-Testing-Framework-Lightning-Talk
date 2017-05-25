package keys

// GetKeys client tokens must be signed by one of the server keys provided via a url.
// The keys expire after a certain amount of time so we need to track that also.
func GetKeys(keyURL string, mTokens map[string]interface{}, mMaxAge int64, mErr error) (tokens map[string]interface{}, maxAge int64, err error) {
	return mTokens, mMaxAge, mErr
}
