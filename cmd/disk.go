package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cobra"
)

// Raiz do comando

var diskCmd = &cobra.Command{
	Use: "disk",
	Short: "Exibe informações do disco.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----DISK-----")

        // Disk

		disk_info, err := disk.Usage("/")
		if err != nil {
			fmt.Println("Error retrieving information about disk.", err)
			return
		}

		total_disk_GB := float64(disk_info.Total) / (1024 * 1024 * 1024)
	    available_disk_GB := float64(disk_info.Free) / (1024 * 1024 * 1024)
		used_disk_GB := float64(disk_info.Used) / (1024 * 1024 * 1024)

		fmt.Printf("Total Disk: %.2f GB\n", total_disk_GB)
		fmt.Printf("Free Disk: %.2f GB\n", available_disk_GB)
		fmt.Printf("Used Disk: %.2f GB\n", used_disk_GB)
		fmt.Printf("Used Percent Disk: %.2f%%\n", disk_info.UsedPercent)
		fmt.Printf("-----PARTITIONS-----\n")

		// Disk (partitions)

		disk_partitions_info, err := disk.Partitions(false)
		if err != nil {
			fmt.Println("Error retrieving information about partitions.", err)
			return
		}

		for _, dp := range disk_partitions_info {
			fmt.Printf("Device: %s  Mounted on: %s  Fstype: %s\n", dp.Device, dp.Mountpoint, dp.Fstype)
		}
		fmt.Println("")
	},
}

// Registro de subcomando

func init() {
	rootCmd.AddCommand(diskCmd)
}