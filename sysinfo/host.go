package sysinfo

import (
	"bytes"
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

func GetHOSTInfo() string {
	var out bytes.Buffer
	
	out.WriteString("\n-----KERNEL-----\n")

	// Host (Kernel & Operating System)

	host_info, err := host.Info()
	if err != nil {
		out.WriteString("Error retrieving information about kernel and OS: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(fmt.Sprintf("Kernel: %s\n", host_info.OS))
	out.WriteString(fmt.Sprintf("Kernel Version: %s\n", host_info.KernelVersion))
	out.WriteString(fmt.Sprintf("Kernel Architecture: %s\n", host_info.KernelArch))

    out.WriteString("\n-----OPERATING SYSTEM-----\n")

	out.WriteString(fmt.Sprintf("Hostname: %s\n", host_info.Hostname))
	out.WriteString(fmt.Sprintf("Platform: %s\n", host_info.Platform))
	out.WriteString(fmt.Sprintf("Uptime (minutes): %d", host_info.Uptime/60))

	return out.String()
}