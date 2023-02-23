package clusters

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestClusters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Public clusters Suite")
}

var _ = BeforeSuite(func() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Before suite PARENT SUITE")
	time.Sleep(5 * time.Second)
})

var _ = AfterSuite(func() {
	fmt.Println("After suite PARENT SUITE")
})
