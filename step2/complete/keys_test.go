package keys_test

import (
	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step2/complete"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	var (
		keyURL       = "blah.org"
		keys         map[string]interface{}
		maxAge       int64
		err          error
		expectedKeys map[string]interface{}
	)

	BeforeEach(func() {

		expectedKeys = make(map[string]interface{})
		expectedKeys["key1"] = "fdjslkdfjfdsalfjds"
		expectedKeys["key2"] = "dsfjasdlfjdsaflkdj"
		expectedKeys["key3"] = "klfjdsalfjdsalkjfd"

		keys, maxAge, err = GetKeys(keyURL)
	})

	It("should not return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return the correct maxAge", func() {
		Expect(maxAge).To(Equal(int64(1000)))
	})

	It("should return the correct key map", func() {
		Expect(keys).To(Equal(expectedKeys))
	})

})
