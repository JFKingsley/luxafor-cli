package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Finding connected Luxafor Flags...")

	devices := Discover()

	if len(devices) == 0 {
		fmt.Println("Could not find any attached flag.")
		os.Exit(0)
		return
	}

	print("Found " + strconv.Itoa(len(devices)) + " Luxafor Flag(s)...")

	device := devices[0]

	err := device.Connect()

	if err != nil {
		panic(err)
	}

	command := &Command{}
	red, _ := strconv.Atoi(os.Args[1])
	green, _ := strconv.Atoi(os.Args[2])
	blue, _ := strconv.Atoi(os.Args[3])
	command = command.Type(Color).Position(Both).Color(uint8(red), uint8(green), uint8(blue)).Fade(50)

	err = device.Command(command)

	if err != nil {
		panic(err)
	}

	err = device.Disconnect()

	if err != nil {
		panic(err)
	}
}