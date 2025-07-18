package sysinfo

import (
	"bytes"
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func GetDISKInfo() string {
	var out bytes.Buffer

	out.WriteString("-----DISK-----\n")

    // Disk

	disk_info, err := disk.Usage("/")
	if err != nil {
		out.WriteString("Error retrieving information about Disk: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(fmt.Sprintf("Total Disk: %.2f GB\n", float64(disk_info.Total)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Free Disk: %.2f GB\n", float64(disk_info.Free)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Used Disk: %.2f GB\n", float64(disk_info.Used)/(1024*1024*1024)))
	out.WriteString(fmt.Sprintf("Used Percent Disk: %.2f%%\n", disk_info.UsedPercent))

	out.WriteString("\n-----PARTITIONS-----\n")

	// Disk (Partitions)

	partitions, err := disk.Partitions(false)
	if err != nil {
		out.WriteString("Error retrieving information about partitions: " + err.Error() + "\n")
		return out.String()
	}
	for _, part := range partitions {
		out.WriteString(fmt.Sprintf("Device: %s  Mounted on: %s\n", part.Device, part.Mountpoint))
	}


	return out.String()
}