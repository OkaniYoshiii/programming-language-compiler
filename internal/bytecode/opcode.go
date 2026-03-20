package bytecode

import (
	"encoding/binary"
)

const WordSizeInBytes uint32 = 4

const (
	OpConst byte = iota
)

// OpCode = 1 byte; address = 4 bytes
const OpConstSize = 1 + 4

// 1 Operand (2 byte)
var OpConstOperands = [1]uint8{2}

type Metadata struct {
	value []int
}

func NewMetadata(operandsCount int) Metadata {
	return Metadata{
		value: make([]int, operandsCount),
	}
}

func (metadata Metadata) SetOperandWidth(index int, width int) Metadata {
	metadata.value[index] = width

	return metadata
}

func (metadata Metadata) OperandsCount() int {
	return len(metadata.value)
}

func (metadata Metadata) OperandWidth(index int) int {
	return metadata.value[index]
}

func OpConstMeta() Metadata {
	return NewMetadata(1).SetOperandWidth(0, int(WordSizeInBytes))
}

func MakeOpConst(address uint32) Instruction {
	operands := [4]byte{}
	binary.BigEndian.PutUint32(operands[:], address)
	return Instruction{
		OpCode: OpConst,
		Operands: operands[:],
	}
}
