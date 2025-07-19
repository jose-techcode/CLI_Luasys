package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var biosCmd = &cobra.Command{
	Use: "bios",
	Short: "Exibe informações da BIOS.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----BIOS-----")

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
					break
				}
			}
		}
		fmt.Println("")
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(biosCmd)
}