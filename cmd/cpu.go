package cmd

import(
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use: "cpu",
	Short: "Exibe informações da CPU.",
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Printf("-----CPU-----\n")

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
		} else {
			fmt.Println("Nothing found...")
		}
		fmt.Println("")
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(cpuCmd)
}