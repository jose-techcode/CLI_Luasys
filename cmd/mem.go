package cmd


import(
	"fmt"

	"github.com/spf13/cobra"
    "github.com/shirou/gopsutil/v3/mem"
)

// Raiz do comando

var memCmd = &cobra.Command{
	Use: "mem",
	Short: "Exibe informações da memória.",
	Run: func(cmd *cobra.Command, args []string) {

        fmt.Println("-----RAM-----")

		// RAM

		ram_info, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error retrieving information about RAM.", err)
			return
		}
	    
		total_ram_GB := float64(ram_info.Total) / (1024 * 1024 * 1024)
	    available_ram_GB := float64(ram_info.Available) / (1024 * 1024 * 1024)
		used_ram_GB := float64(ram_info.Used) / (1024 * 1024 * 1024)
		
		fmt.Printf("Total RAM: %.2f GB\n", total_ram_GB)
		fmt.Printf("Free RAM: %.2f GB\n", available_ram_GB)
		fmt.Printf("Used RAM: %.2f GB\n", used_ram_GB)
		fmt.Printf("Used Percent RAM: %.2f%%\n", ram_info.UsedPercent)
		fmt.Println("-----Swap-----")

		// Swap

		swap_info, err := mem.SwapMemory()
		if err != nil {
			fmt.Println("Error retrieving information about swap.", err)
			return
		}

		total_swap_GB := float64(swap_info.Total) / (1024 * 1024 * 1024)
	    available_swap_GB := float64(swap_info.Free) / (1024 * 1024 * 1024)
		used_swap_GB := float64(swap_info.Used) / (1024 * 1024 * 1024)

		fmt.Printf("Total Swap: %.2f GB\n", total_swap_GB)
		fmt.Printf("Free Swap: %.2f GB\n", available_swap_GB)
		fmt.Printf("Used Swap: %.2f GB\n", used_swap_GB)
		fmt.Printf("Used Percent Swap: %.2f%%\n", swap_info.UsedPercent)
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(memCmd)
}