package vm

import (
	"encoding/binary"

	"github.com/okaniyoshiii/programming-language/internal/bytecode"
)

var InstructionPointer uint32 = 0

const WordSize = bytecode.WordSizeInBytes

// Assumes that bytecode uses 4 bytes addresses (32 bits)
func Run(input []byte) []byte {
	bytecodeSize := binary.BigEndian.Uint32(input[0:4])
	constantsSize := binary.BigEndian.Uint32(input[4:8])
	bytecodeSection := input[8:8+bytecodeSize]
	constantsSection := input[8+bytecodeSize:8+bytecodeSize+constantsSize]

	for _, b := range bytecodeSection {
		_ = b
	}

	for _, b := range constantsSection {
		_ = b
	}

	return []byte{}
}
