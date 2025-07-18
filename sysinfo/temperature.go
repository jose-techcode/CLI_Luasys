package sysinfo

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetTEMPERATUREInfo() string {
	var out bytes.Buffer

	out.WriteString("-----TEMPERATURE-----\n")

	// Temperature

	temperature_info, err := exec.Command("sensors").Output()
	if err != nil {
		out.WriteString("Error retrieving information about temperature: " + err.Error() + "\n")
		return out.String()
	}
	
	// Percorer uma lista para encontrar a temperatura geral da CPU

    cpu_text := string(temperature_info)

	cpu_lines := strings.Split(cpu_text, "\n")

	for _, cpu_line := range cpu_lines {
			if strings.Contains(cpu_line, "Package id 0") {
				parts := strings.SplitN(cpu_line, ":", 2)
				if len(parts) == 2 {
					cpu := strings.TrimSpace(parts[1])
					out.WriteString((fmt.Sprintf("CPU Temperature: %s\n", cpu)))
					break
				}
			}
		}

	// Percorer uma lista para encontrar a temperatura do core 0

	core0_text := string(temperature_info)

	core0_lines := strings.Split(core0_text, "\n")

	for _, core0_line := range core0_lines {
			if strings.Contains(core0_line, "Core 0") {
				parts := strings.SplitN(core0_line, ":", 2)
				if len(parts) == 2 {
					core0 := strings.TrimSpace(parts[1])
					out.WriteString(fmt.Sprintf("Core 0: %s\n", core0))
					break
				}
			}
		}

    // Percorer uma lista para encontrar a temperatura do core 1

	core1_text := string(temperature_info)

	core1_lines := strings.Split(core1_text, "\n")

	for _, core1_line := range core1_lines {
			if strings.Contains(core1_line, "Core 0") {
				parts := strings.SplitN(core1_line, ":", 2)
				if len(parts) == 2 {
					core1 := strings.TrimSpace(parts[1])
					out.WriteString(fmt.Sprintf("Core 1: %s\n", core1))
					break
				}
			}
		}

	// Percorer uma lista para encontrar a temperatura da placa-mãe

	temp1_text := string(temperature_info)

	temp1_lines := strings.Split(temp1_text, "\n")

	for _, temp1_line := range temp1_lines {
			if strings.Contains(temp1_line, "temp1") {
				parts := strings.SplitN(temp1_line, ":", 2)
				if len(parts) == 2 {
					temp1 := strings.TrimSpace(parts[1])
					out.WriteString(fmt.Sprintf("Motherboard Temperature: %s\n", temp1))
					break
				}
			}
		}

	return out.String()
}