package cmd

import (
	"bytes"
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

// Exportable function

func Runsys() string {
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

	out.WriteString("\n-----GPU-----\n")

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

	out.WriteString("\n-----MOTHERBOARD-----\n")

	// Motherboard

	motherboard_info, err := exec.Command("dmidecode", "-t", "2").Output()
	if err != nil {
		out.WriteString("Error retrieving information about Motherboard: " + err.Error() + "\n")
		return out.String()
	}

	info_lines := strings.Split(string(motherboard_info), "\n")

	// Iterate through a list to find the motherboard manufacturer

	for _, line := range info_lines {
		if strings.Contains(line, "Manufacturer") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				manufacturer := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("Motherboard Manufacturer: %s\n", manufacturer))
				break
			}
		}
	}

	// Iterate through a list to find the motherboard model

	for _, line := range info_lines {
		if strings.Contains(line, "Product Name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				model := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("Motherboard model: %s\n", model))
				break
			}
		}
	}

	out.WriteString("\n-----BIOS-----\n")

	// BIOS

	bios_info, err := exec.Command("dmidecode", "-t", "bios").Output()
	if err != nil {
		out.WriteString("Error retrieving information about BIOS: " + err.Error() + "\n")
		return out.String()
	}
 
	// Iterate through a list to find the bios vendor

	for _, line := range strings.Split(string(bios_info), "\n") {
		if strings.Contains(line, "Vendor") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				vendor := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("BIOS Vendor: %s\n", vendor))
				break
			}
		}
	}

    // Iterate through a list to find the bios version

	for _, line := range strings.Split(string(bios_info), "\n") {
		if strings.Contains(line, "Version") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				version := strings.TrimSpace(parts[1])
				out.WriteString(fmt.Sprintf("BIOS Version: %s\n", version))
				break
			}
		}
	}

	out.WriteString("\n-----RAM-----\n")

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

	out.WriteString("\n-----DISK-----\n")

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

	out.WriteString("\n-----BATTERY-----\n")

	// Battery

	battery_info, err := exec.Command("acpi").Output()
	if err != nil {
		out.WriteString("Error retrieving information about battery: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(string(battery_info))

	out.WriteString("\n-----TEMPERATURE-----\n")

    // Temperature

	temperature_info, err := exec.Command("sensors").Output()
	if err != nil {
		out.WriteString("Error retrieving information about temperature: " + err.Error() + "\n")
		return out.String()
	}

    cpu_text := string(temperature_info)

	cpu_lines := strings.Split(cpu_text, "\n")

	// Iterate through a list to find the overall CPU temperature

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

	core0_text := string(temperature_info)

	core0_lines := strings.Split(core0_text, "\n")

	// Iterate through a list to find the core 0 temperature

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

	core1_text := string(temperature_info)

	core1_lines := strings.Split(core1_text, "\n")

	// Iterate through a list to find the core 1 temperature

	for _, core1_line := range core1_lines {
			if strings.Contains(core1_line, "Core 1") {
				parts := strings.SplitN(core1_line, ":", 2)
				if len(parts) == 2 {
					core1 := strings.TrimSpace(parts[1])
					out.WriteString(fmt.Sprintf("Core 1: %s\n", core1))
					break
				}
			}
		}

	temp1_text := string(temperature_info)

	temp1_lines := strings.Split(temp1_text, "\n")

	// Iterate through a list to find the motherboard temperature

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

	out.WriteString("\n-----USB-----\n")

	// USB

	usb_info, err := exec.Command("lsusb").Output()
	if err != nil {
		out.WriteString("Error retrieving information about USB: " + err.Error() + "\n")
		return out.String()
	}
	out.WriteString(string(usb_info))

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
	out.WriteString(fmt.Sprintf("Uptime (minutes): %d\n", host_info.Uptime/60))

	return out.String()
}

// Raiz do comando

var rootCmd = &cobra.Command{
	Use: "luasys",
	Short: "Uma CLI tool multifuncional de auditoria, diagnóstico e monitoramento de recursos do hardware. Admin/Sudo.",
	Run: func(cmd *cobra.Command, args []string) {
		info := Runsys()
		fmt.Println(info)
	},
}

// Tratamento de erros

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}