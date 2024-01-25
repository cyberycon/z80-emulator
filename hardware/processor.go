// Implementation of Z80 processor logic
package hardware

const NOP byte = 0X00
const HLT byte = 0X01


type Cpu struct {
	pc uint16 
	memory    *Ram
	isRunning bool
}

// type Register8 byte

// type Register16 struct {
// 	value int16
// 	high  Register8
// 	low   Register8
// }

// Reset the program counter to 0 
func (c *Cpu) Reset() {
	c.pc = 0
	c.Run()
}

// Start executing instructions frome the current Program counter location
func (c *Cpu) Run() {
	for c.isRunning {
		c.pc = c.Step()
	}
}

// Execute the instruction at the Program Counter and advance it to the next instruction. 
func (c *Cpu) Step() uint16 {
	instruction := c.memory.cells[c.pc]
	switch instruction {
	case HLT:
		
		c.isRunning = false
		
	}

	return c.pc + 1
}
