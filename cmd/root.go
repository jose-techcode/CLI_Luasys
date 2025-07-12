package cmd

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

// Raiz do comando

var rootCmd = &cobra.Command{
	Use: "luatop",
	Short: "Uma CLI tool de monitoramento e diagnóstico de recursos do hardware.",
	Run: func(cmd *cobra.Command, args []string) {

        fmt.Println("-----CPU-----")

		// CPU (info)

		cpu_info, err := cpu.Info()
		if err != nil {
			fmt.Println("Error retrieving information about CPU.", err)
			return
		}

		if len(cpu_info) > 0 {
			fmt.Println("CPU Model:", cpu_info[0].ModelName)
			fmt.Println("CPU Manufacturer:", cpu_info[0].VendorID)
			fmt.Println("CPU Physical Core:", cpu_info[0].Cores)
		} else {
			fmt.Println("Nothing found...")
		}
		
		// CPU (counts)

		cpu_counts, err := cpu.Counts(false)
		if err != nil {
			fmt.Println("Error retrieving information about CPU.", err)
			return
		}

		if len(cpu_info) > 0 {
			fmt.Println("CPU Logical Core:", cpu_counts)
			fmt.Println("-----RAM-----")
		} else {
			fmt.Println("Nothing found...")
		}

		// RAM

		mem_ram_info, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error retrieving information about RAM.", err)
			return
		}
	    
		total_ram_GB := float64(mem_ram_info.Total) / (1024 * 1024 * 1024)
	    available_ram_GB := float64(mem_ram_info.Available) / (1024 * 1024 * 1024)
		used_ram_GB := float64(mem_ram_info.Used) / (1024 * 1024 * 1024)
		
		fmt.Printf("Total RAM: %.2f GB\n", total_ram_GB)
		fmt.Printf("Free RAM: %.2f GB\n", available_ram_GB)
		fmt.Printf("Used RAM: %.2f GB\n", used_ram_GB)
		fmt.Printf("Used Percent RAM: %.2f%%\n", mem_ram_info.UsedPercent)
		fmt.Println("-----Swap-----")
		
        // Swap

		mem_swap_info, err := mem.SwapMemory()
		if err != nil {
			fmt.Println("Error retrieving information about swap.", err)
			return
		}

		total_swap_GB := float64(mem_swap_info.Total) / (1024 * 1024 * 1024)
	    available_swap_GB := float64(mem_swap_info.Free) / (1024 * 1024 * 1024)
		used_swap_GB := float64(mem_swap_info.Used) / (1024 * 1024 * 1024)

		fmt.Printf("Total Swap: %.2f GB\n", total_swap_GB)
		fmt.Printf("Free Swap: %.2f GB\n", available_swap_GB)
		fmt.Printf("Used Swap: %.2f GB\n", used_swap_GB)
		fmt.Printf("Used Percent Swap: %.2f%%\n", mem_swap_info.UsedPercent)
		fmt.Println("-----Disk-----")

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
		fmt.Println("-----Partitions-----")

		// Disk (partitions)

		disk_partitions_info, err := disk.Partitions(false)
		if err != nil {
			fmt.Println("Error retrieving information about partitions.", err)
			return
		}

		for _, dp := range disk_partitions_info {
			fmt.Printf("Device: %s  Mounted on: %s  Fstype: %s\n", dp.Device, dp.Mountpoint, dp.Fstype)
		}
		fmt.Println("-----Kernel-----")

		// Host (Kernel & Operating System)
				
		host_info, err := host.Info()
		if err != nil {
			fmt.Println("Error retrieving information about Kernel and Operating System.", err)
			return
		}

		// Kernel

		fmt.Println("Kernel:", host_info.OS)
	    fmt.Println("Kernel Version:", host_info.KernelVersion)
		fmt.Println("Kernel Architecture:", host_info.KernelArch)
		fmt.Println("-----Operating System-----")
        
		// Operating System

		fmt.Println("Hostname:", host_info.Hostname)
		fmt.Println("Plataform:", host_info.Platform)
		fmt.Println("Uptime (minutes):", host_info.Uptime / 60)
	},
}

// Tratamento de erros

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}