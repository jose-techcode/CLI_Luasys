package sysinfo

import (
	"bytes"
	"fmt"

	"github.com/shirou/gopsutil/v3/mem"
)

func GetMEMInfo() string {
	var out bytes.Buffer

	out.WriteString("-----MEM-----\n")

	// Memory (RAM)

	ram_info, err := mem.VirtualMemory()
	if err != nil {
		out.WriteString("Error retrieving information about RAM: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(fmt.Sprintf("Total RAM: %.2f GB\n", float64(ram_info.Total)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Free RAM: %.2f GB\n", float64(ram_info.Available)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Used RAM: %.2f GB\n", float64(ram_info.Used)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Used Percent RAM: %.2f%%\n", ram_info.UsedPercent))

	out.WriteString("\n-----SWAP-----\n")

	// Memory (Swap)

	swap_info, err := mem.SwapMemory()
	if err != nil {
		out.WriteString("Error retrieving information about swap: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(fmt.Sprintf("Total Swap: %.2f GB\n", float64(swap_info.Total)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Free Swap: %.2f GB\n", float64(swap_info.Free)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Used Swap: %.2f GB\n", float64(swap_info.Used)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Used Percent Swap: %.2f%%\n", swap_info.UsedPercent))

	return out.String()
}