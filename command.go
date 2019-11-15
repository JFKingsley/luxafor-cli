package main

import "errors"

type Command struct {
	commandType CommandType
	position Position

	patternByte byte

	redByte byte
	greenByte byte
	blueByte byte

	timingByteOne byte
	timingByteTwo byte
	timingByteThree byte
}

func (c *Command) Type (commandType CommandType) *Command {
	c.commandType = commandType

	return c
}

func (c *Command) Position(position Position) *Command {
	c.position = position

	return c
}

func (c *Command) Pattern(patternType PatternType) *Command {
	if c.commandType != Pattern {
		panic(errors.New("You cannot set a pattern on a command that is not a Pattern command"))
	}

	c.patternByte = byte(patternType)

	return c
}

func (c *Command) Color(red uint8, green uint8, blue uint8) *Command {
	c.redByte = red
	c.greenByte = green
	c.blueByte = blue

	return c
}

func (c *Command) Fade(speed byte) *Command {
	if c.commandType != Color && c.commandType != Fade {
		panic(errors.New("You cannot set a fade on a command that is not a Color or Fade command."))
	}

	if c.commandType == Color {
		c.commandType = Fade
	}

	c.timingByteOne = speed

	return c
}

func (c *Command) Speed(speed byte) *Command {
	if c.commandType == Color || c.commandType == Pattern {
		panic(errors.New("You cannot set speed on a command that is a Color or Pattern command."))
	}

	if c.commandType == Fade {
		panic(errors.New("You must set fade timings via the Fade command."))
	}

	if c.commandType == Strobe {
		c.timingByteOne = speed
	}

	if c.commandType == Wave {
		c.timingByteThree = speed
	}

	return c
}

func (c *Command) Repeat(repeat byte) *Command {
	if c.commandType != Strobe && c.commandType != Wave {
		panic(errors.New("You cannot set repeat on a command that is not a Strobe or Wave command."))
	}

	if c.commandType == Strobe {
		c.timingByteThree = repeat
	}

	if c.commandType == Wave {
		c.timingByteTwo = repeat
	}

	return c
}

func (c *Command) Bytes() []byte {
	patternOrPositionByte := byte(c.position)

	if c.commandType == Pattern {
		patternOrPositionByte = c.patternByte
	}

	return []byte{
		byte(c.commandType),
		patternOrPositionByte,
		c.redByte,
		c.greenByte,
		c.blueByte,
		c.timingByteOne,
		c.timingByteTwo,
		c.timingByteThree,
	}
}