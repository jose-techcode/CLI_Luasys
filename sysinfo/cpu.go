package sysinfo

import (
	"bytes"
	"fmt"
	
	"github.com/shirou/gopsutil/v3/cpu"

	
)

func GetCPUInfo() string {
	var out bytes.Buffer

	out.WriteString("-----CPU-----\n")

	// CPU (info)

	cpu_info, err := cpu.Info()
	if err != nil {
		out.WriteString("Error retrieving information about CPU: " + err.Error() + "\n")
		return out.String()
	}

	if len(cpu_info) > 0 {
		out.WriteString(fmt.Sprintf("CPU Model: %v\n", cpu_info[0].ModelName))
		out.WriteString(fmt.Sprintf("CPU Manufacturer: %v\n", cpu_info[0].VendorID))
		out.WriteString(fmt.Sprintf("CPU Physical Core: %v\n", cpu_info[0].Cores))
	} else {
		out.WriteString("Nothing found...\n")
	}

	// CPU (counts)

	cpu_counts, err := cpu.Counts(false)
	if err != nil {
		out.WriteString("Error retrieving information about CPU.\n")
		return out.String()
	}

	if len(cpu_info) > 0 {
		out.WriteString(fmt.Sprintf("CPU Logical Core: %v\n", cpu_counts))
	} else {
		out.WriteString("Nothing found...\n")
	}
	
	return out.String()
}