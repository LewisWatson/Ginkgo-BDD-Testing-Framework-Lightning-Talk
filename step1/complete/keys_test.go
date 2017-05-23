package keys_test

import (
	. "github.com/LewisWatson/Ginkgo-BDD-Testing-Framework-Lightning-Talk/step1/complete"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	var (
		keyURL = "blah.org"
		// keys map[string]interface{}
		// maxAge int64
		err error
	)

	BeforeEach(func() {
		_, _, err = GetKeys(keyURL)
	})

	It("should not return an error", func() {
		Expect(err).NotTo(HaveOccurred())
	})

})
