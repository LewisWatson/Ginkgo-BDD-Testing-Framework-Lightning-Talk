package keys_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	jsonKeys  string
	jsonKeys2 string
	token     string
	token2    string
)

func TestKeys(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keys suite")
}

var _ = BeforeSuite(func() {
	content, err := ioutil.ReadFile("testdata/keys.json")
	Expect(err).NotTo(HaveOccurred())
	jsonKeys = string(content)
})
