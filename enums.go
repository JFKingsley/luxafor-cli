package main

type CommandType byte

const (
	Color CommandType = 1
	Fade = 2
	Strobe = 3
	Wave = 4
	Pattern = 6
)

func StringToCommandType(commandtype string) CommandType {
	if commandtype == "color" {
		return Color
	}

	if commandtype == "fade" {
		return Fade
	}

	if commandtype == "strobe" {
		return Strobe
	}

	if commandtype == "wave" {
		return Wave
	}

	if commandtype == "pattern" {
		return Pattern
	}

	return Color
}

type Position byte

const (
	One Position = 1
	Two = 2
	Three = 3
	Four = 4
	Five = 5
	Six = 6
	Both = 0xff
	Back = 0x42
	Front = 0x41
)

type PatternType byte

const (
	Police PatternType = 5
)