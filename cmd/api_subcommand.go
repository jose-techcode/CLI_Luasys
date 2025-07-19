package cmd

import (
	"fmt"

	"github.com/jose-techcode/CLI_Luasys/api"
	"github.com/spf13/cobra"
)

// API REST (GET)

var apiCmd = &cobra.Command{
	Use: "api",
	Short: "Inicia servidor HTTP da API. Localhost.",
	Long: "Rotas GET: /cpu, /gpu, /motherboard, /bios, /mem, /disk, /battery, /temperature, /usb, /host, /net, /processes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Iniciando servidor HTTP na porta 8000...")
		api.Server()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}