package bytecode

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("Constant output", func(t *testing.T) {
		builder := BytecodeBuilder{}

		litteral := "Hello, World !"
		index := builder.AddConstant([]byte(litteral))

		operands := [4]byte{}
		binary.BigEndian.PutUint32(operands[:], uint32(index))
		builder.AddInstruction(OpConst, operands[:])

		output := builder.Build()

		bytecodeSize := OpConstSize
		constantSize := 4 + len(litteral)

		expected := []byte{}
		// header
		expected = binary.BigEndian.AppendUint32(expected, uint32(bytecodeSize))
		expected = binary.BigEndian.AppendUint32(expected, uint32(constantSize))
		// bytecode
		expected = append(expected, OpConst)
		expected = append(expected, operands[:]...)
		// constants
		expected = binary.BigEndian.AppendUint32(expected, uint32(len(litteral)))
		expected = append(expected, []byte(litteral)...)

		if !bytes.Equal(output, expected) {
			t.Errorf("expected %04b, got %04b", expected, output)
		}
	})
}
