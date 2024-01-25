package hardware

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)
var _ = Describe("Memory test suite", func() {
	ram := Ram{} 
	ram.Initialize(256)
	Context("Empty memory", func() {
		It("Should be initialized to a size", func() {

			ram.WriteAt(0x23, 0x10)
			result, _ := ram.Read(0x10)
			Expect(result).Should(Equal(byte(0x23))) 
		})

		It("Should reject reads out of range", func() {
		
			_, err := ram.Read(0x100) 
			Expect(err.Error()).Should(Equal("Address out of range"))
		
		})
	})
})

