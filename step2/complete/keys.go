package keys

// GetKeys client tokens must be signed by one of the server keys provided via a url.
// The keys expire after a certain amount of time so we need to track that also.
func GetKeys(keyURL string) (tokens map[string]interface{}, maxAge int64, err error) {
	tokens = make(map[string]interface{})
	tokens["key1"] = "fdjslkdfjfdsalfjds"
	tokens["key2"] = "dsfjasdlfjdsaflkdj"
	tokens["key3"] = "klfjdsalfjdsalkjfd"
	return tokens, int64(1000), nil
}
