package sysinfo

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func GetGPUInfo() string {
	var out bytes.Buffer

	out.WriteString("-----GPU-----\n")

    // GPU

	gpu_info, err := exec.Command("lspci").Output()
	if err != nil {
		out.WriteString("Error retrieving information about GPU: " + err.Error() + "\n")
		return out.String()
	}

    // Iterate through a list to find the gpu

	for _, line := range strings.Split(string(gpu_info), "\n") {
		if strings.Contains(line, "VGA compatible controller") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				model := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("GPU Model: %s\n", model))
				break
			}
		}
	}
	
	return out.String()
}