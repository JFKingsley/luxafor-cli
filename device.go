package main

import "github.com/karalabe/hid"

type LuxaforFlag struct {
	hidDeviceInfo hid.DeviceInfo
	hidDevice *hid.Device
}

func (f *LuxaforFlag) Connect() error {
	device, err := f.hidDeviceInfo.Open()

	if err != nil {
		return err
	}

	f.hidDevice = device

	return nil
}

func (f *LuxaforFlag) Disconnect() error {
	return f.hidDevice.Close()
}

func (f *LuxaforFlag) Command(c *Command) error {
	_, err := f.hidDevice.Write(c.Bytes())

	return err
}

func Discover() []LuxaforFlag {
	devices := hid.Enumerate(1240, 62322)

	flags := make([]LuxaforFlag, 0)

	for _, device := range devices {
		flags = append(flags, LuxaforFlag{
			hidDeviceInfo: device,
			hidDevice:     nil,
		})
	}

	return flags
}
