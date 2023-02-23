package clusters

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Public clusters", func() {
	It("testing apps", func() {
		fmt.Println("AAPPSS")
		n := rand.Intn(1)
		time.Sleep(time.Duration(n) * time.Second)
	})

	It("testing ingress", func() {
		fmt.Println("INGRESSSSS")
		n := rand.Intn(1)
		time.Sleep(time.Duration(n) * time.Second)
	})

	It("testing autoscaler", func() {
		fmt.Println("AUTOSCALEEEERR")
		n := rand.Intn(1)
		time.Sleep(time.Duration(n) * time.Second)
	})

	It("testing volumes", func() {
		fmt.Println("VOLUMEESSSS")
		n := rand.Intn(1)
		time.Sleep(time.Duration(n) * time.Second)
	})
})
