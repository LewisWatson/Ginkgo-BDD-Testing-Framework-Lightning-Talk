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

	It("should return max age from response header", func() {
		Expect(maxAge).To(Equal(int64(19008)))
	})

	It("should return a max age greater than 10", func() {
		Expect(maxAge).To(BeNumerically(">", 10))
	})

	It("should return 4 keys", func() {
		Expect(len(serverTokens)).To(Equal(4))
	})

	Context("response doesn't contain max age", func() {

		BeforeEach(func() {

			/*
			 * given
			 */

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, jsonKeys)
			}))
			defer ts.Close()

			url := ts.URL

			/*
			 * when
			 */

			serverTokens, maxAge, err = GetKeys(url)
		})

		It("should return ErrCacheControlHeaderLacksMaxAge", func() {
			Expect(err).To(Equal(ErrCacheControlHeaderLacksMaxAge))
		})

	})

	Context("no server response", func() {

		BeforeEach(func() {

			/*
			 * given
			 */

			url := "notgonnaresolv.org.uk.eu"

			/*
			 * when
			 */

			serverTokens, maxAge, err = GetKeys(url)
		})

		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
		})

	})

})
