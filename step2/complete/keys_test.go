package keys_test

import (
	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step2"

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

		// given

		givenTokens = make(map[string]interface{})
		givenTokens["key1"] = "fdjslkdfjfdsalfjds"
		givenTokens["key2"] = "dsfjasdlfjdsaflkdj"
		givenTokens["key3"] = "klfjdsalfjdsalkjfd"
		// givenTokens["bad token"] = "bad token"

		givenMaxAge := int64(1001)

		// when

		tokens, maxAge, err = GetKeys("keys.url", givenTokens, givenMaxAge, nil)
	})

	//then

	It("should not return an error", func() {
		Expect(err).ToNot(HaveOccurred())
	})

	It("should return a max age greater than 1000", func() {
		Expect(maxAge).To(BeNumerically(">", 1000))
	})

	It("should return the expected tokens", func() {
		Expect(tokens).To(Equal(givenTokens))
	})

	It("should not return the bad token", func() {
		Expect(tokens).NotTo(ContainElement("bad token"))
	})
})
