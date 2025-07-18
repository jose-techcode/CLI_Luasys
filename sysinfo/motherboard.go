package sysinfo

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetMOTHERBOARDInfo() string {
	var out bytes.Buffer

	out.WriteString("-----MOTHERBOARD-----\n")

	// Motherboard

	motherboard_info, err := exec.Command("dmidecode", "-t", "2").Output()
	if err != nil {
		out.WriteString("Error retrieving information about Motherboard: " + err.Error() + "\n")
		return out.String()
	}
	info_lines := strings.Split(string(motherboard_info), "\n")
	for _, line := range info_lines {
		if strings.Contains(line, "Manufacturer") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				manufacturer := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("Motherboard Manufacturer: %s\n", manufacturer))
				break
			}
		}
	}
	for _, line := range info_lines {
		if strings.Contains(line, "Product Name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				model := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("Motherboard model: %s\n", model))
				break
			}
		}
	}

	return out.String()
}