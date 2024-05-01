// Represents processor memory space.
package hardware

const OUT_OF_RANGE = "Address out of range"

// RAM memory as an array of bytes with a sixteen bit size
type Ram struct {
	cells []byte
	size  uint16
}

type MemoryError struct {
	 message string
	}

func (m *MemoryError) Error () string {
	return m.message
}

// Initialize the memory array to the specified size
func (r *Ram) Initialize(size uint16) {
	r.cells = make([]byte, size)
	r.size = size
}

// Set all memory cells to zero. 
func (r *Ram) ClearMemory() {
	for i := 0; i < int(r.size); i++ {
		r.cells[i] = 0
	}
}

// Wrrite a byte to a specfied address
func (r *Ram) WriteAt(value byte, address uint16) error {
	if address >= r.size {
		return &MemoryError{message: OUT_OF_RANGE}
	}
	r.cells[address] = value 
	
	return nil
}

func (r *Ram) Read(address uint16) (byte, error) {
	if address >= r.size {
		return 0, &MemoryError{message: OUT_OF_RANGE}
	}
	return r.cells[address], nil	
}

func (r *Ram) WriteSlice(value []byte, startAddress uint16) error {
	if len(value) + int(startAddress) > int(r.size) {
		return &MemoryError{message: OUT_OF_RANGE}
	}
	copy(r.cells[startAddress:], value)
	return nil
	 
}
