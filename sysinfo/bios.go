package sysinfo

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetbiosInfo() string {
	var out bytes.Buffer

	out.WriteString("-----BIOS-----\n")

	// BIOS

	bios_info, err := exec.Command("dmidecode", "-t", "bios").Output()
	if err != nil {
		out.WriteString("Error retrieving information about BIOS: " + err.Error() + "\n")
		return out.String()
	}
	for _, line := range strings.Split(string(bios_info), "\n") {
		if strings.Contains(line, "Vendor") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				vendor := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("BIOS Vendor: %s\n", vendor))
				break
			}
		}
	}
	for _, line := range strings.Split(string(bios_info), "\n") {
		if strings.Contains(line, "Version") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				version := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("BIOS Version: %s\n", version))
				break
			}
		}
	}
	
	return out.String()
}