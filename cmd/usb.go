package cmd

import(
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var usbCmd = &cobra.Command{
	Use: "usb",
	Short: "Exibe informações da dispositivos USB.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("-----USB-----")

		// Usb

		usb_info, err := exec.Command("lsusb").Output()
		if err != nil {
			fmt.Println("Error retrieving information about usb devices.", err)
			return
		}

		usb_text := strings.TrimSpace(string(usb_info))

		fmt.Println(usb_text)
	},
}

// Registro do subcomando

func init() {
	rootCmd.AddCommand(usbCmd)
}