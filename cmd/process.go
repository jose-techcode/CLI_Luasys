package cmd

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

// Raiz do comando

var processesCmd = &cobra.Command{
	Use: "processes",
	Short: "Exibe informações de processos.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----PROCESSESES-----")

		// Process

		process_info, err := process.Processes()
		if err != nil {
			fmt.Println("Error retrieving information about processes.", err)
			return
		}

		for _, proc := range process_info {

		// Usuário

		user_info, err := proc.Username()
		if err != nil {
			fmt.Println("Error retrieving information about process.", err)
			return
		}

		// Status

		status_info, err := proc.Status()
		if err != nil {
			fmt.Println("Error retrieving information about process.", err)
			return
		}

		// Nome

		name_info, err := proc.Name()
		if err != nil {
			fmt.Println("Error retrieving information about process.", err)
			return
		}

		// PID

		pid_info := proc.Pid

		// CPU Percent

		cpu_percent_1_info, err := proc.CPUPercent()
		if err != nil {
			fmt.Println("Error retrieving information about process.", err)
			return
		}

		// RAM percent

		ram_percent_1_info, err := proc.MemoryPercent()
		if err != nil {
			fmt.Println("Error retrieving information about process.", err)
			return
		}
			fmt.Printf("User: %s  Status: %s  Name: %s  Pid: %d  CPU Percent: %.2f%%  RAM Percent: %.2f%%\n", user_info, status_info[0], name_info, pid_info, cpu_percent_1_info, ram_percent_1_info)
		}
	},
}

// Registro de subcomando

func init() {
	rootCmd.AddCommand(processesCmd)
}