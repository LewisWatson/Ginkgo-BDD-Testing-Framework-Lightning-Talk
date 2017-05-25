package keys_test

import (
	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step2/complete"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	var (
		tokens      map[string]interface{}
		maxAge      int64
		err         error
		givenTokens map[string]interface{}
	)

	BeforeEach(func() {

		/*
		 * given
		 */

		givenTokens = make(map[string]interface{})
		givenTokens["key1"] = "fdjslkdfjfdsalfjds"
		givenTokens["key2"] = "dsfjasdlfjdsaflkdj"
		givenTokens["key3"] = "klfjdsalfjdsalkjfd"
		// givenTokens["bad token"] = "bad token"

		givenMaxAge := int64(1001)

		/*
		 * when
		 */

		tokens, maxAge, err = GetKeys("keys.url", givenTokens, givenMaxAge, nil)
	})

	It("should not return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return a maxAge greater than 10", func() {
		Expect(maxAge).To(BeNumerically(">", 10))
	})

	It("should return the correct key map", func() {
		Expect(tokens).To(Equal(givenTokens))
	})

	It("should not return the bad token", func() {
		Expect(tokens).ToNot(ContainElement("bad token"))
	})

})
