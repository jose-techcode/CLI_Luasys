package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var motherboardCmd = &cobra.Command{
	Use: "motherboard",
	Short: "Exibe informações da placa-mãe. Admin/Sudo.",
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("-----MOTHERBOARD-----")

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
					break
				}
			}
		}
		fmt.Println("")
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(motherboardCmd)
}