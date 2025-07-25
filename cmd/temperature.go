package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var tempCmd = &cobra.Command{
	Use: "temp",
	Short: "Exibe informações da temperatura.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----TEMPERATURE-----")

		// Temperature

		temperature_info, err := exec.Command("sensors").Output()
		if err != nil {
			fmt.Println("Error retrieving information about temperature.", err)
			return
		}

        cpu_text := string(temperature_info)

		cpu_lines := strings.Split(cpu_text, "\n")

		// // Iterate through a list to find the overall cpu temperature

		for _, cpu_line := range cpu_lines {
			if strings.Contains(cpu_line, "Package id 0") {
				parts := strings.SplitN(cpu_line, ":", 2)
				if len(parts) == 2 {
					cpu := strings.TrimSpace(parts[1])
					fmt.Println("CPU Temperature:", cpu)
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
					fmt.Println("Core 0:", core0)
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
					fmt.Println("Core 1:", core1)
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
					fmt.Println("Motherboard Temperature:", temp1)
					break
				}
			}
		}
		fmt.Println("")
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(tempCmd)
}