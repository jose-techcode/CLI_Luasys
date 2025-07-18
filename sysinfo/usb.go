package sysinfo

import (
	"bytes"
	"os/exec"

)

func GetUSBInfo() string {
	var out bytes.Buffer

	out.WriteString("-----USB-----\n")

	// USB

	usb_info, err := exec.Command("lsusb").Output()
	if err != nil {
		out.WriteString("Error retrieving information about USB: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(string(usb_info))

	return out.String()
}
