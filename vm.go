package main

import (
	"io"
	"log"
)

const (
	currentInstruction        = "current instruction: %c"
	currentInstructionPointer = "current ip: %s"
	currentDataPointer        = "current dp: %s"
)

type BFMachine struct {
	instructions []*Instruction
	memory       [30000]int
	ip           int
	dp           int
	out          io.Writer
	in           io.Reader
	buf          []byte
}

func NewMachine(instructions []*Instruction, in io.Reader, out io.Writer) *BFMachine {
	return &BFMachine{
		instructions: instructions,
		in:           in,
		out:          out,
		buf:          make([]byte, 1),
	}
}

func (m *BFMachine) Execute() {
	log.Printf("code length: %d", len(m.instructions))
	for m.ip < len(m.instructions) {
		ins := m.instructions[m.ip]
		// log.Printf(currentInstruction, (instruction))
		switch ins.Type {
		case Plus:
			m.memory[m.dp] += ins.Argument
		case Minus:
			m.memory[m.dp] -= ins.Argument
		case Right:
			m.dp += ins.Argument
		case Left:
			m.dp -= ins.Argument
		case PutChar:
			m.putChar()
		case ReadChar:
			m.readChar()
		case JumpIfZero:
			m.jumpIfZero()
		case JumpIfNotZero:
			m.jumpIfNotZero()
		}

		m.ip++
	}
}

func (m *BFMachine) jumpIfZero() {
	if m.memory[m.dp] == 0 {
		depth := 1
		for depth != 0 {
			m.ip++
			switch m.instructions[m.ip].Type {
			case JumpIfZero:
				depth++
			case JumpIfNotZero:
				depth--
			}
		}
	}
}

func (m *BFMachine) jumpIfNotZero() {
	if m.memory[m.dp] != 0 {
		depth := 1
		for depth != 0 {
			m.ip--
			switch m.instructions[m.ip].Type {
			case JumpIfZero:
				depth--
			case JumpIfNotZero:
				depth++
			}
		}
	}
}

func (m *BFMachine) readChar() {
	n, err := m.in.Read(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong number of bytes written")
	}
	m.memory[m.dp] = int(m.buf[0])
}

func (m *BFMachine) putChar() {
	m.buf[0] = byte(m.memory[m.dp])
	n, err := m.out.Write(m.buf)
	if err != nil {
		panic(err)
	}
	if n != 1 {
		panic("wrong number of bytes written")
	}
}
