package keys

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrCacheControlHeaderLacksMaxAge indicates that the key server response didnt contain a max age
	// as specified by the firebase docs
	ErrCacheControlHeaderLacksMaxAge = errors.New("cache control header doesn't contain a max age")
)

// HeaderCacheControl Cache-Control field in http response header
const HeaderCacheControl = "Cache-Control"

var myClient = &http.Client{Timeout: 30 * time.Second}

// GetKeys client tokens must be signed by one of the server keys provided via a url.
// The keys expire after a certain amount of time so we need to track that also.
func GetKeys(keyURL string) (tokens map[string]interface{}, maxAge int64, err error) {

	// r, err := myClient.Get(keyURL)
	// if err != nil {
	// 	return tokens, maxAge, err
	// }
	// defer r.Body.Close()

	// maxAge, err = extractMaxAge(r.Header.Get(HeaderCacheControl))
	// if err != nil {
	// 	return tokens, maxAge, err
	// }

	// err = json.NewDecoder(r.Body).Decode(&tokens)

	return tokens, maxAge, err
}

// Extract the max age from the cache control response header value
// The cache control header should look similar to "..., max-age=19008, ..."
func extractMaxAge(cacheControl string) (int64, error) {
	// "..., max-age=19008, ..."" to ["..., max-age="]["19008, ..."]
	tokens := strings.Split(cacheControl, "max-age=")
	if len(tokens) == 1 {
		return 0, ErrCacheControlHeaderLacksMaxAge
	}
	// "19008, ..." to ["19008"][" ..."]
	tokens2 := strings.Split(tokens[1], ",")
	// convert "19008" to int64
	return strconv.ParseInt(tokens2[0], 10, 64)
}
