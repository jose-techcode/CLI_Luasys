package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var batteryCmd = &cobra.Command{
	Use: "battery",
	Short: "Exibe informações da bateria.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----BATTERY-----")

		// Battery

		battery_info, err := exec.Command("acpi").Output()
		if err != nil {
			fmt.Println("Error retrieving information about battery.", err)
			return
		}

		battery_text := strings.TrimSpace(string(battery_info))

		fmt.Println(battery_text)
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(batteryCmd)
}