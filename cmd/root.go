package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

// Raiz do comando

var rootCmd = &cobra.Command{
	Use: "luasys",
	Short: "Uma CLI tool de auditoria, diagnóstico e monitoramento de recursos do hardware.",
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
			fmt.Println("-----GPU-----")
		} else {
			fmt.Println("Nothing found...")
		}

		// GPU

		gpu_info, err := exec.Command("lspci").Output()
		if err != nil {
			fmt.Println("Error retrieving information about GPU.", err)
			return
		}

		text := string(gpu_info)

		lines := strings.Split(text, "\n")

		// Percorrer uma lista para encontrar a linha da GPU

		for _, line := range lines {
			if strings.Contains(line, "VGA compatible controller") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					model := strings.TrimSpace(parts[1])
					fmt.Println("GPU Model:", model)
					fmt.Println("-----MOTHERBOARD-----")
					break
				}
			}
		}

		// Motherboard

		motherboard_info, err := exec.Command("dmidecode", "-t", "2").Output()
		if err != nil {
			fmt.Println("Error retrieving information about motherboard. Need sudo.", err)
			return
		}

	    info_text := string(motherboard_info)

		info_lines := strings.Split(info_text, "\n")

		// Percorrer uma lista para encontrar o manufacturer

        for _, info_line := range info_lines {
			if strings.Contains(info_line, "Manufacturer") {
				parts := strings.SplitN(info_line, ":", 2)
				if len(parts) == 2 {
					manufacturer := strings.TrimSpace(parts[1])
					fmt.Println("Motherboard Manufacturer:", manufacturer)
					break
				}
			}
		}

		// Percorer uma lista para encontrar a linha da placa-mãe e tirar a redundãncia "Product Name"

		for _, info_line := range info_lines {
			if strings.Contains(info_line, "Product Name") {
				parts := strings.SplitN(info_line, ":", 2)
				if len(parts) == 2 {
					model := strings.TrimSpace(parts[1])
					fmt.Println("Motherboard model:", model)
					fmt.Println("-----BIOS-----")
					break
				}
			}
		}

		// BIOS

		bios_info, err := exec.Command("dmidecode", "-t", "bios").Output()
		if err != nil {
			fmt.Println("Error retrieving information about GPU.", err)
			return
		}

		info_list := string(bios_info)

		info_strings := strings.Split(info_list, "\n")

		// Percorer uma lista para encontrar o vendor

		for _, info_string := range info_strings {
			if strings.Contains(info_string, "Vendor") {
				parts := strings.SplitN(info_string, ":", 2)
				if len(parts) == 2 {
					vendor := strings.TrimSpace(parts[1])
					fmt.Println("BIOS Vendor:", vendor)
					break
				}
			}
		}

        // Percorer uma lista para encontrar a versão

		for _, info_string := range info_strings {
			if strings.Contains(info_string, "Version") {
				parts := strings.SplitN(info_string, ":", 2)
				if len(parts) == 2 {
					version := strings.TrimSpace(parts[1])
					fmt.Println("BIOS Version:", version)
					fmt.Println("-----RAM-----")
					break
				}
			}
		}

		// Memory (RAM)

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
		fmt.Println("-----SWAP-----")
		
        // Memory (Swap)

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
		fmt.Println("-----PARTITIONS-----")

		// Disk (Partitions)

		disk_partitions_info, err := disk.Partitions(false)
		if err != nil {
			fmt.Println("Error retrieving information about partitions.", err)
			return
		}

		for _, dp := range disk_partitions_info {
			fmt.Printf("Device: %s  Mounted on: %s\n", dp.Device, dp.Mountpoint)
		}
		fmt.Println("-----KERNEL-----")

		// Battery

		// Temperature

		// Fans

		// Usb

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
		fmt.Println("-----OPERATING SYSTEM-----")
        
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