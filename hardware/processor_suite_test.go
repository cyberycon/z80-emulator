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
 

var _ = Describe("Processor test suite", func() {
	Context("PC step", func() {
		var testMemory = Ram{} 
		testMemory.Initialize(256)
		var cpu = makeCpu(&testMemory)
		
		It("Should stop on HLT", func() {
			testMemory.WriteAt(HLT, 0x10)
			cpu.Reset()
			Expect(cpu.pc).Should(Equal(uint16(0x11)))
		})
		It("Should load accumulator", func() {
			testMemory.ClearMemory()
			testMemory.WriteAt(LD_A_I, 0)
			testMemory.WriteAt(5,1)
			testMemory.WriteAt(HLT, 2)
			cpu.Reset() 
			Expect(cpu.GetAccumulator()).Should(Equal(byte(5))) 
		})
	})
})