package sysinfo

import (
	"bytes"
	"os/exec"
)

func GetBATTERYInfo() string {
	var out bytes.Buffer

	out.WriteString("-----BATTERY-----\n")

	// Battery

	battery_info, err := exec.Command("acpi").Output()
	if err != nil {
		out.WriteString("Error retrieving information about battery: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(string(battery_info))
	
	return out.String()
}