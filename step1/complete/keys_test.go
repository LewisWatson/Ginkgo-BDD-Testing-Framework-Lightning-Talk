package keys_test

import (
	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step1/complete"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	var (
		tokens map[string]interface{}
		maxAge int64
		err    error
	)

	BeforeEach(func() {
		tokens, maxAge, err = GetKeys("keys.url")
	})

	It("should not return an error", func() {
		Expect(err).ToNot(HaveOccurred())
	})

})
