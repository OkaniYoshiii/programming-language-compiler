// Bytecode file format
// =================================
// header                           <- metadata about the file
//	  bytecode_size               	<- 4 bytes. Used to determine the start of the "bytecode_section"
//    constant_size                	<- 4 bytes. Used to determine the start of the "constant_section"
//
// bytecode_section                 <- contains all the bytecode instructions
//	  bytecode_instruction          <- 1 byte per OpCode + variable numbers of bytes for operands
//	  ...
//
// constant_section                 <- contains all constants of the program (string litteral, int litteral ...)
//	   constant                     <- 4 bytes for the length of the constant + "length" bytes of the actual data
//	   ...
// =================================
package bytecode

import "encoding/binary"

type Constant struct {
	Value []byte
}

type Instruction struct {
	OpCode byte
	Operands []byte
}

type BytecodeBuilder struct{
	Constants []Constant
	Instructions []Instruction
}

func (builder *BytecodeBuilder) AddInstruction(opcode byte, operands []byte) int {
	instruction := Instruction{OpCode: opcode, Operands: operands}

	builder.Instructions = append(builder.Instructions, instruction)

	return len(builder.Instructions) - 1
}

func (builder *BytecodeBuilder) AddConstant(value []byte) int {
	constant := Constant{Value: value}

	builder.Constants = append(builder.Constants, constant)

	return len(builder.Constants) - 1
}

func (builder *BytecodeBuilder) BytecodeSize() int {
	size := 0
	for _, instruction := range builder.Instructions {
		// opcode -> 1 byte + operands length
		size += 1 + len(instruction.Operands)
	}

	return size
}

func (builder *BytecodeBuilder) ConstantsSize() int {
	size := 0
	for _, constant := range builder.Constants {
		// length -> 4 bytes + data length
		size += 4 + len(constant.Value)
	}

	return size
}

func (builder *BytecodeBuilder) Build() []byte {
	// bytecodeAddress -> 4 bytes + constantAddress -> 4 bytes
	headerSize := 8
	bytecodeSize := builder.BytecodeSize()
	constantsSize := builder.ConstantsSize()

	programSize := headerSize + bytecodeSize + constantsSize

	bytecodeAddress := headerSize
	constantAddress := headerSize + bytecodeSize

	bytecode := make([]byte, programSize)

	// HEADER
	binary.BigEndian.PutUint32(bytecode, uint32(bytecodeSize))
	binary.BigEndian.PutUint32(bytecode[4:8], uint32(constantsSize))

	// BYTECODE
	offset := bytecodeAddress
	for _, instruction := range builder.Instructions {
		bytecode[offset] = instruction.OpCode
		for j, operand := range instruction.Operands {
			bytecode[offset + 1 + j] = operand
		}

		// opcode -> 1 bytes + operands length
		offset += 1 + len(instruction.Operands)
	}

	// CONSTANTS
	offset = constantAddress
	for _, constant := range builder.Constants {
		binary.BigEndian.PutUint32(bytecode[offset:offset+4], uint32(len(constant.Value)))
		for j, b := range constant.Value {
			bytecode[offset + 4 + j] = b
		}

		// length -> 4 bytes + data length
		offset += 4 + len(constant.Value)
	}

	return bytecode
}
