package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var gpuCmd = &cobra.Command{
	Use: "gpu",
	Short: "Exibe informações da GPU.",
	Run: func(cmd *cobra.Command, args []string) {
		
		fmt.Println("-----GPU-----")

		// GPU

		gpu_info, err := exec.Command("lspci").Output()
	    if err != nil {
			fmt.Println("Error retrieving information about GPU.", err)
			return
		}

	    text := string(gpu_info)

	    lines := strings.Split(text, "\n")

	    // Iterate through a list to find the gpu

	    for _, line := range lines {
		    if strings.Contains(line, "VGA compatible controller") {
			    parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					model := strings.TrimSpace(parts[1])
					fmt.Println("GPU Model:", model)
					break
				}
			}
		}
		fmt.Println("")
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(gpuCmd)
}