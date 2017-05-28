package keys_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step3/complete"

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

	It("should extract maxAge from response header", func() {
		Expect(maxAge).To(Equal(int64(19008)))
	})

	It("should populate serverTokens with four keys", func() {
		Expect(len(serverTokens)).To(Equal(4))
	})

	Context("key server response does not contain max-age", func() {

		BeforeEach(func() {

			/*
			 * given
			 */

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(HeaderCacheControl, "something other than max age")
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

		It("should return an ErrCacheControlHeaderLacksMaxAge error", func() {
			Expect(err).To(Equal(ErrCacheControlHeaderLacksMaxAge))
		})

	})

	Context("no server response", func() {

		BeforeEach(func() {

			/*
			 * given
			 */

			url := "url-to-nowhere.sdlafsdale4.org.uk.net"

			/*
			 * when
			 */

			serverTokens, maxAge, err = GetKeys(url)
		})

		/*
		 * then
		 */

		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
		})

	})

})
