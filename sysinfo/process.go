package sysinfo

import (
	"bytes"
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

func GetPROCESSESInfo() string {
	var out bytes.Buffer
	
	out.WriteString("-----PROCESSES-----\n")

	// Process

	process_info, err := process.Processes()
	if err != nil {
		out.WriteString("Error retrieving information about processes: " + err.Error() + "\n")
	    return out.String()
	}

	for _, proc := range process_info {

		user_info, err := proc.Username()
		if err != nil {
			out.WriteString("Error retrieving information about processes: " + err.Error() + "\n")
		return out.String()
		}

		status_info, err := proc.Status()
		if err != nil {
			out.WriteString("Error retrieving information about processes: " + err.Error() + "\n")
		return out.String()
		}

		name_info, err := proc.Name()
		if err != nil {
			out.WriteString("Error retrieving information about processes: " + err.Error() + "\n")
		return out.String()
		}

		pid_info := proc.Pid

		cpu_percent_1_info, err := proc.CPUPercent()
		if err != nil {
			out.WriteString("Error retrieving information about processes: " + err.Error() + "\n")
		return out.String()
		}

		ram_percent_1_info, err := proc.MemoryPercent()
		if err != nil {
			out.WriteString("Error retrieving information about processes: " + err.Error() + "\n")
		return out.String()
		}
			out.WriteString(fmt.Sprintf("User: %s  Status: %s  Name: %s  Pid: %d  CPU Percent: %.2f%%  RAM Percent: %.2f%%\n", user_info, status_info[0], name_info, pid_info, cpu_percent_1_info, ram_percent_1_info))
		}
		
		return out.String()
	}