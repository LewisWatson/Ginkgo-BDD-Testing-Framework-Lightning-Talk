package keys_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	var (
		maxAge       int64
		err          error
		serverTokens map[string]interface{}
	)

	BeforeEach(func() {

		/*
		 * given
		 */

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(HeaderCacheControl, "..., max-age=19008, ...")
			fmt.Fprintln(w, jsonKeys)
		}))
		defer ts.Close()

		url := ts.URL

		/*
		 * when
		 */

		serverTokens, maxAge, err = GetKeys(url)
	})

	/*
	 * then
	 */

	It("should not throw an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

})
