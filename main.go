package main

import "fmt"
import "github.com/yukkeorg/go-usbio2"

func main() {
	usbio, err := usbio2.NewUsbIO2()
	if err != nil {
		return
	}
	defer usbio.Cleanup()

	fmt.Printf("Device Name: %s\n", usbio.GetDeviceName())

        cmd, _ := usbio.CreateCommand(usbio2.CMD_WRITEREAD)
	r, err := usbio.WriteRead(cmd)
	if err != nil {
		return
	}

	fmt.Printf("ReadData : %v\n", r)
}
