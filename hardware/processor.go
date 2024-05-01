// Implementation of Z80 processor logic
package hardware

const NOP byte = 0X00
const LD_A_I = 0b00111110
const ADD_A_n = 0XC6 
const HLT byte = 0X76


type Cpu struct {
	pc uint16 
	memory    *Ram
	isRunning bool
	A byte 
	F byte 
}

// type Register8 byte

// type Register16 struct {
// 	value int16
// 	high  Register8
// 	low   Register8
// }

func makeCpu(memory *Ram) *Cpu {
	return &Cpu{0, memory, false, 0, 0}
}

// Reset the program counter to 0 
func (c *Cpu) Reset() {
	c.pc = 0
	c.Run()
}

// Start executing instructions frome the current Program counter location
func (c *Cpu) Run() {
	c.isRunning = true
	for c.isRunning {
		println("pc", c.pc)
		c.pc = c.Step()
	}
}

// Execute the instruction at the Program Counter and advance it to the next instruction. 
func (c *Cpu) Step() uint16 {
	instruction, err := c.memory.Read(c.pc)
	if err != nil {
		c.isRunning = false; 
		return c.pc
	}
	switch instruction {
	case HLT:
		
		c.isRunning = false

	case LD_A_I: 
		return c.loadRegisterImmediate(instruction)
		
	}

	return c.pc + 1
}

// Return the contents of register A
func (c *Cpu) GetAccumulator() (byte)  {
	return c.A 
}
func (c *Cpu) loadRegisterImmediate(instruction byte) (uint16) {
	c.pc++
	value, err := c.memory.Read(c.pc)
	if (err != nil) {
		c.isRunning = false; 
	}
	c.A = value 
	return c.pc + 1 // advance 1 byte for instruction 1 byte for data
}
